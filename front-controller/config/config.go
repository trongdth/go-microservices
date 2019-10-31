package config

import (
	"log"
	"os"
	"strconv"
)

var config *Config

func init() {
	env := os.Getenv("env")
	port, _ := strconv.Atoi(os.Getenv("port"))
	secretKey := os.Getenv("secret_key")
	entryCacheEndpoint := os.Getenv("entry_cache_endpoint")

	log.Printf("env: %s, port: %d, secret_key: %s, entry_cache: %s",
		env, port, secretKey, entryCacheEndpoint)

	config = &Config{
		Environment:        env,
		Port:               port,
		TokenSecretKey:     secretKey,
		EntryCacheEndpoint: entryCacheEndpoint,
	}
}

func GetConfig() *Config {
	return config
}

// Config : struct
type Config struct {
	Environment        string `json:"env"`
	Port               int    `json:"port"`
	TokenSecretKey     string `json:"secret_key"`
	EntryCacheEndpoint string `json:"entry_cache_endpoint"`
}
