package config

import (
	"os"
)

var config *Config

func init() {
	env := os.Getenv("env")
	db := os.Getenv("db")
	secretKey := os.Getenv("token_secret_key")

	config = &Config{
		Environment:    env,
		Db:             db,
		TokenSecretKey: secretKey,
	}
}

func GetConfig() *Config {
	return config
}

// Config : struct
type Config struct {
	Environment    string `json:"env"`
	Db             string `json:"db"`
	TokenSecretKey string `json:"token_secret_key"`
}
