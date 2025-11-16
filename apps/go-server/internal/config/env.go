package config

type Env struct {
	HOST                   string
	PORT                   string
	JWTSecret              string
	JWTExpirationInSeconds int64
}

func getEnv() Env {
	return Env{
		HOST:                   getEnvStr("API_PUBLIC_HOST", "localhost", false),
		PORT:                   getEnvStr("API_SERVER_PORT", "", true),
		JWTSecret:              getEnvStr("API_JWT_SECRET", "supersecret", false), // TODO: must be required
		JWTExpirationInSeconds: getEnvInt64("API_JWT_EXPIRATION_IN_SECONDS", 3600*24*7, false),
	}
}
