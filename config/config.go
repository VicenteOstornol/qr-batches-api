package config

import "github.com/Netflix/go-env"

type (
	Config struct {
		HTTP
		Postgres
	}

	HTTP struct {
		HTTPPort string `env:"SERVICE_PORT,required=true"`
	}

	Postgres struct {
		Username string `env:"POSTGRES_USERNAME,required=true"`
		Password string `env:"POSTGRES_PASSWORD,required=true"`
		Host     string `env:"POSTGRES_HOST,required=true"`
		Port     string `env:"POSTGRES_PORT,required=true"`
		DBName   string `env:"POSTGRES_DB,required=true"`
	}
)

func New() (*Config, error) {
	var conf Config
	_, err := env.UnmarshalFromEnviron(&conf)
	if err != nil {
		return nil, err
	}

	return &conf, nil
}
