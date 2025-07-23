package config

import (
	"fmt"
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/joho/godotenv"
)

type (
	Config struct {
		App    App    `toml:"app"`
		API    API    `toml:"api"`
		Log    Log    `toml:"log"`
		DB     DB     `toml:"db"`
		JWT    JWT    `toml:"jwt"`
		OAuth2 OAuth2 `toml:"oauth2"`
	}

	App struct {
		Mode string `toml:"mode"`
	}
	API struct {
		Port   int    `toml:"port"`
		Host   string `toml:"host"`
		Schema string `toml:"schema"`
	}

	DB struct {
		PoolMax  int    `toml:"pool_max"`
		Host     string `toml:"host"`
		DataBase string `toml:"database"`
		User     string `toml:"user"`
		Password string `toml:"password"`
		Port     int    `toml:"port"`
	}

	Log struct {
		Level string `toml:"level"`
	}

	JWT struct {
		AccessTokenExpiryHour  int    `toml:"access_token_expiry_hour"`
		RefreshTokenExpiryHour int    `toml:"refresh_token_expiry_hour"`
		AccessTokenSecret      string `toml:"access_token_secret"`
		RefreshTokenSecret     string `toml:"refresh_token_secret"`
	}
	OAuth2 struct {
		Google           Google `toml:"google"`
		OAuthStateString string `toml:"state_string"`
	}
	Google struct {
		GoogleClientId     string `toml:"client_id"`
		GoogleClientSecret string `toml:"client_secret"`
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := godotenv.Load("config/.env")
	if err != nil {
		log.Fatal("failed to load .env file")
	}

	tomlCfg, err := os.ReadFile("config/dev.toml")
	if err != nil {
		log.Fatal("failed to load toml config")
	}
	tomlCfgStr := os.Expand(string(tomlCfg), func(key string) string {
		return os.Getenv(key)
	})

	if _, err := toml.Decode(tomlCfgStr, cfg); err != nil {
		log.Fatal(fmt.Errorf("failed to decode TOML: %w", err))
	}

	return cfg, nil
}
