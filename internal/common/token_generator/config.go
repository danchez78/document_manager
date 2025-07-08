package token_generator

import "fmt"

type Config struct {
	SecretKey       string `yaml:"secret_key"`
	TokenTTLMinutes int    `yaml:"token_ttl_minutes"`
}

func (c *Config) Validate() error {
	if c.SecretKey == "" {
		return fmt.Errorf("secret_key not provided")
	}

	if c.TokenTTLMinutes == 0 {
		return fmt.Errorf("token_ttl_minutes not provided")
	}

	return nil
}
