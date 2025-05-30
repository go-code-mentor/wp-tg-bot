package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
)

func New() *Config {
	return &Config{}
}

type Config struct {
	GRPC
	Telegram
}

type GRPC struct {
	Host string `env:"GRPC_HOST" env-default:"localhost"`
	Port string `env:"GRPC_HOST_PORT" env-default:"8080"`
}

type Telegram struct {
	Token  string `env:"TELEGRAM_TOKEN" env-default:""`
	ChatID string `env:"TELEGRAM_CHAT_ID" env-default:""`
}

func (c *Config) ParseConfig() error {
	if err := cleanenv.ReadEnv(c); err != nil {
		return fmt.Errorf("failed to read env's :%w", err)
	}
	return nil
}

func (c *Config) GrpcConnString() string {
	return fmt.Sprintf("%s:%s", c.Host, c.Port)
}
