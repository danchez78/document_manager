package domain

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"regexp"
	"strings"

	"golang.org/x/crypto/argon2"
)

type params struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	saltLength  uint32
	keyLength   uint32
}

type User struct {
	ID             string
	Login          string
	HashedPassword string
	Token          string
}

type UserInput struct {
	Login    string
	Password string
}

func NewUserInput(login, password string) (*UserInput, error) {
	u := &UserInput{
		Login:    login,
		Password: password,
	}

	if err := u.validate(); err != nil {
		return nil, err
	}

	return u, nil
}

func (u *UserInput) HashPassword() (string, error) {
	p := &params{
		memory:      64 * 1024,
		iterations:  3,
		parallelism: 2,
		saltLength:  16,
		keyLength:   32,
	}

	salt := make([]byte, p.saltLength)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(u.Password), salt, p.iterations, p.memory, p.parallelism, p.keyLength)

	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	encodedHash := fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, p.memory, p.iterations, p.parallelism, b64Salt, b64Hash)

	return encodedHash, nil
}

func (u *UserInput) СomparePasswordAndHash(encodedHash string) (bool, error) {
	p, salt, hash, err := decodeHash(encodedHash)
	if err != nil {
		return false, err
	}

	otherHash := argon2.IDKey([]byte(u.Password), salt, p.iterations, p.memory, p.parallelism, p.keyLength)

	if subtle.ConstantTimeCompare(hash, otherHash) == 1 {
		return true, nil
	}
	return false, nil
}

func (u *UserInput) validate() error {
	if err := u.validateLogin(); err != nil {
		return err
	}
	if err := u.validatePassword(); err != nil {
		return err
	}

	return nil
}

func (u *UserInput) validateLogin() error {
	if len(u.Login) < 8 {
		return fmt.Errorf("user login length less than 8")
	}

	re := regexp.MustCompile(`^[a-zA-Z0-9]+$`)
	if !re.MatchString(u.Login) {
		return fmt.Errorf("user login not contains any latin or numeric symbols")
	}

	return nil
}

func (u *UserInput) validatePassword() error {
	if len(u.Password) < 8 {
		return fmt.Errorf("user password length less than 8")
	}

	re := regexp.MustCompile(`[a-z]`)
	if !re.MatchString(u.Password) {
		return fmt.Errorf("user password not contains any lower symbols")
	}

	re = regexp.MustCompile(`[A-Z]`)
	if !re.MatchString(u.Password) {
		return fmt.Errorf("user password not contains any upper symbols")
	}

	re = regexp.MustCompile(`[0-9]`)
	if !re.MatchString(u.Password) {
		return fmt.Errorf("user password not contains any numbers")
	}

	re = regexp.MustCompile(`[!"#$%&'()*+,\-./:;<=>?@[\\\]^_{|}~]`)
	if !re.MatchString(u.Password) {
		return fmt.Errorf("user password not contains any special symbols")
	}

	return nil
}

var (
	ErrInvalidHash         = errors.New("the encoded hash is not in the correct format")
	ErrIncompatibleVersion = errors.New("incompatible version of argon2")
)

func decodeHash(encodedHash string) (p *params, salt, hash []byte, err error) {
	vals := strings.Split(encodedHash, "$")
	if len(vals) != 6 {
		return nil, nil, nil, ErrInvalidHash
	}

	var version int
	_, err = fmt.Sscanf(vals[2], "v=%d", &version)
	if err != nil {
		return nil, nil, nil, err
	}
	if version != argon2.Version {
		return nil, nil, nil, ErrIncompatibleVersion
	}

	p = &params{}
	_, err = fmt.Sscanf(vals[3], "m=%d,t=%d,p=%d", &p.memory, &p.iterations, &p.parallelism)
	if err != nil {
		return nil, nil, nil, err
	}

	salt, err = base64.RawStdEncoding.Strict().DecodeString(vals[4])
	if err != nil {
		return nil, nil, nil, err
	}
	p.saltLength = uint32(len(salt)) //nolint:gosec

	hash, err = base64.RawStdEncoding.Strict().DecodeString(vals[5])
	if err != nil {
		return nil, nil, nil, err
	}
	p.keyLength = uint32(len(hash)) //nolint:gosec

	return p, salt, hash, nil
}
