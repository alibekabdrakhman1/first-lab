package configs

import (
	"errors"

	"github.com/caarlos0/env/v7"
)

type Config struct {
	HTTP     HTTP
	Database Database
}
type HTTP struct {
	PORT string `env:"PORT" envDefault:":8080"`
	URL  string `env:"PORT" envDefault:"localhost"`
}
type Database struct {
	HOST     string `env:"DB_HOST" envDefault:"localhost"`
	PORT     string `env:"DB_PORT" envDefault:"5432"`
	USER     string `env:"DB_USER" envDefault:"postgres"`
	PASSWORD string `env:"DB_PASSWORD" envDefault:"253165alibek"`
	DB       string `env:"DB_NAME" envDefault:"onelab"`
	PgUrl    string `env:"PG_URL" envDefault:"host=localhost port=5432 user=postgres password=253165alibek dbname=onelab sslmode=disable"`
}

func New() (*Config, error) {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		return nil, errors.New("cfg not created")
	}
	return &cfg, nil
}
