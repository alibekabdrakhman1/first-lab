package configs

import (
	"errors"

	"github.com/caarlos0/env/v7"
)

type Config struct {
	PORT string `env:"PORT" envDefault:":8080"`
}

func New() (*Config, error) {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		return nil, errors.New("cfg not created")
	}
	return &cfg, nil
}
