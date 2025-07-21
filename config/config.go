package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

type (
	Config struct {
		HTTP   HTTP
		Log    Log
		PG     PG
		JWT    JWT
		OAuth2 OAuth2
	}

	HTTP struct {
		Port           int    `env:"HTTP_PORT,required"`
		IP             string `env:"HTTP_IP"`
		Schema         string `env:"HTTP_SCHEMA"`
		UsePreforkMode bool   `env:"HTTP_USE_PREFORK_MODE,required"`
	}

	PG struct {
		PoolMax int    `env:"PG_POOL_MAX,required"`
		URL     string `env:"PG_URL,required"`
	}

	Log struct {
		Level int `env:"LOG_LEVEL,required"`
	}

	JWT struct {
		AccessTokenExpiryHour  int    `env:"ACCESS_TOKEN_EXPIRY_HOUR"`
		RefreshTokenExpiryHour int    `env:"REFRESH_TOKEN_EXPIRY_HOUR"`
		AccessTokenSecret      string `env:"ACCESS_TOKEN_SECRET"`
		RefreshTokenSecret     string `env:"REFRESH_TOKEN_SECRET"`
	}
	OAuth2 struct {
		GorillaSession struct {
			Key string `env:"GORILLA_SESSION_KEY,required"`
		}
		Google struct {
			GoogleClientId     string `env:"GOOGLE_CLIENT_ID,required"`
			GoogleClientSecret string `env:"GOOGLE_CLIENT_SECRET,required"`
		}
		OAuthStateString string `env:"OAUTH_STATE_STRING"`
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	return cfg, nil
}
