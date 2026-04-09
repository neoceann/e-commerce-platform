package config

import (
	"fmt"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort string `env:"SERVER_PORT" env-default:"8080"`
	ServerHost string `env:"SERVER_HOST" env-default:"localhost"`

	DBHost     string `env:"DB_HOST" env-default:"localhost"`
	DBPort     string `env:"DB_PORT" env-default:"5432"`
	DBUser     string `env:"DB_USER" env-default:"postgres"`
	DBPassword string `env:"DB_PASSWORD" env-required:"true"`
	DBName     string `env:"DB_NAME" env-default:"store"`
	DBSSLMode  string `env:"DB_SSL_MODE" env-default:"disable"`

	ReadTimeout  time.Duration `env:"READ_TIMEOUT" env-default:"10s"`
	WriteTimeout time.Duration `env:"WRITE_TIMEOUT" env-default:"10s"`
	IdleTimeout  time.Duration `env:"IDLE_TIMEOUT" env-default:"60s"`
}

func (c *Config) GetDSN() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		c.DBUser,
		c.DBPassword,
		c.DBHost,
		c.DBPort,
		c.DBName,
		c.DBSSLMode,
	)
}

func NewDSN(c *Config) string {
	return c.GetDSN()
}

func LoadConfig() (*Config, error) {
	_ = godotenv.Load()

	var cfg Config
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		return nil, fmt.Errorf("failed to read env: %w", err)
	}

	return &cfg, nil
}

func (c *Config) ServerAddress() string {
	return fmt.Sprintf("%s:%s", c.ServerHost, c.ServerPort)
}