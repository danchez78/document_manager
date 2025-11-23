package config

import (
	"errors"
	"fmt"
	"os"

	"gopkg.in/yaml.v2"

	"document_manager/internal/application/infrastructure/postgres"
	"document_manager/internal/application/infrastructure/redis"
	"document_manager/internal/application/infrastructure/token"
)

type Config struct {
	AdminToken     string          `yaml:"admin_token"`
	Postgres       postgres.Config `yaml:"postgres"`
	RedisClient    redis.Config    `yaml:"redis"`
	TokenGenerator token.Config    `yaml:"token_generator"`
}

func NewConfig(configPath string) (*Config, error) {
	cfg, err := loadConfig(configPath)
	if err != nil {
		return nil, err
	}

	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	return cfg, nil
}

func loadConfig(configPath string) (*Config, error) {
	content, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to find configuration file")
	}

	cfg := &Config{}

	if err := yaml.Unmarshal([]byte(string(content)), cfg); err != nil {
		return nil, fmt.Errorf("parsing YAML file %s failed: %v", configPath, err)
	}

	return cfg, nil
}

func (c *Config) Validate() error {
	if c.AdminToken == "" {
		return errors.New("admin token not found")
	}
	if c.Postgres == (postgres.Config{}) {
		return errors.New("postgres config not found")
	}

	if err := c.TokenGenerator.Validate(); err != nil {
		return fmt.Errorf("token generator config: %v", err)
	}

	return nil
}
