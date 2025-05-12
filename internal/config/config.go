package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

func ParseConfig() (*Config, error) {
	cfg := Config{}

	if err := cleanenv.ReadEnv(&cfg); err != nil {
		return &cfg, err
	}

	return &cfg, nil
}

type Config struct {
	GRPC
}

type GRPC struct {
	Host string `yaml:"host" env:"GRPC_HOST" env-default:"localhost"`
	Port string `yaml:"port" env:"GRPC_HOST_PORT" env-default:"8080"`
}

func (c *Config) GrpcConnString() string {
	return fmt.Sprintf("%s:%s", c.GRPC.Host, c.GRPC.Port)
}
