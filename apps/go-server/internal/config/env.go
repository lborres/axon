package config

import "log"

type Env struct {
	HOST string `env:"API_PUBLIC_HOST" default:"localhost"`
	PORT string `env:"API_SERVER_PORT" required:"true"`

	// Server Environment
	// "dev" | "test" | "prod"
	APP_ENV string `env:"APP_ENV" default:"prod" required:"true"`

	DATABASE_URL string `env:"DATABASE_URL" required:"true" expand:"true"`
}

func getEnv() Env {
	var e Env
	if err := Load(&e); err != nil {
		log.Fatal(err)
	}
	return e
}

func (e *Env) IsProd() bool {
	return e.APP_ENV == "prod" || e.APP_ENV == "production"
}
