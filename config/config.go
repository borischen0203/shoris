package config

import (
	"log"

	"github.com/caarlos0/env"
)

type envConfig struct {
	Api_token string `env:"Api_token,required"`
}

var (
	// Env is the config
	Env = envConfig{}
)

// Setup setup config function
func Setup() {
	if err := env.Parse(&Env); err != nil {
		log.Fatalf("%+v\n", err)
	}

	// fmt.Printf("%+v\n", Env)
}
