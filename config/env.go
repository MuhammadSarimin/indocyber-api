package config

import (
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
	"github.com/muhammadsarimin/indocyber-api/models"
	"github.com/pkg/errors"
)

var Env models.Config

func Init() {

	if err := load(); err != nil {
		panic(err)
	}

}

func load() error {

	if err := godotenv.Load(); err != nil {
		return errors.Wrap(err, "config/env: load .env file")
	}

	if err := env.Parse(&Env); err != nil {
		return errors.Wrap(err, "config/env: parse config .env file")
	}

	if err := env.Parse(&Env.DB); err != nil {
		return errors.Wrap(err, "config/env: parse redis .env file")
	}

	return nil
}
