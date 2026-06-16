package config

import (
	"fmt"
	"strconv"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	DBHost        string `env:"DB_HOST" env-default:"localhost"`
	DBPort        string `env:"DB_PORT" env-default:"5432"`
	DBUser        string `env:"DB_USER" env-default:"postgres"`
	DBPassword    string `env:"DB_PASSWORD" env-required:"true"`
	DBName        string `env:"DB_NAME_AUTH" env-default:"auth"`
	DBSSLMode     string `env:"DB_SSL_MODE" env-default:"disable"`
	GrpcPort      string `env:"GRPC_SERVER_PORT" env-default:"50051"`
	JWTSecret     string `env:"JWT_SECRET"`
	JWTExpiration string `env:"JWT_EXPIRATION_HOURS" env-default:"1"`
	BcryptCost    int    `env:"BCRYPT_COST" env-default:"12"`
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

func LoadConfig() (*Config, error) {
	var cfg Config
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		return nil, fmt.Errorf("failed to read env: %w", err)
	}

	return &cfg, nil
}

func (c *Config) JWTExpirationToTime() time.Duration {
	hours, _ := strconv.Atoi(c.JWTExpiration)

	return time.Duration(hours) * time.Hour
}
