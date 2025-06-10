package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

type DatabaseConfig struct {
	DBHost     string `env:"DB_HOST" required:"true"`
	DBUser     string `env:"DB_USER" required:"true"`
	DBPassword string `env:"DB_PASSWORD" required:"true"`
	DBName     string `env:"DB_NAME" required:"true"`
	DBPort     string `env:"DB_PORT" required:"true"`
}

func loadDbConfig() (*DatabaseConfig, error) {
	var cfg *DatabaseConfig

	err := env.Parse(cfg)
	if err != nil {
		return nil, fmt.Errorf("load database config error: %w", err)
	}

	return cfg, nil
}

func (c DatabaseConfig) GetDSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		c.DBHost, c.DBUser, c.DBPassword, c.DBName, c.DBPort,
	)
}
