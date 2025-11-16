package config

type Config struct {
	Env Env
}

func Init() Config {
	return Config{
		getEnv(),
	}
}
