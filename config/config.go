package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

type (
	Config struct {
		HTTP HTTP
		Log  Log
		PG   PG
	}

	HTTP struct {
		Port           int  `env:"HTTP_PORT,required"`
		UsePreforkMode bool `env:"HTTP_USE_PREFORK_MODE,required"`
		OAuth2Auth struct {
			GorillaSession struct {
				Key string `env:"GORILLA_SESSION_KEY,required"`
			}
			Google struct {
				GoogleClientId     string `env:"GOOGLE_CLIENT_ID,required"`
				GoogleClientSecret string `env:"GOOGLE_CLIENT_SECRET,required"`
			}
		}
	}

	PG struct {
		PoolMax int    `env:"PG_POOL_MAX,required"`
		URL     string `env:"PG_URL,required"`
	}

	Log struct {
		Level int `env:"LOG_LEVEL,required"`
	}
)

func NewCOnfig() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	return cfg, nil
}