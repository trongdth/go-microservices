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
	entryStoreEndpoint := os.Getenv("entry_store_endpoint")
	redis := os.Getenv("redis")
	redisPwd := os.Getenv("redis_pwd")

	log.Printf("env: %s, port: %d, entry_store: %s", env, port, entryStoreEndpoint)

	config = &Config{
		Environment:        env,
		Port:               port,
		EntryStoreEndpoint: entryStoreEndpoint,
		Redis:              redis,
		RedisPwd:           redisPwd,
	}
}

func GetConfig() *Config {
	return config
}

// Config : struct
type Config struct {
	Environment        string `json:"env"`
	Port               int    `json:"port"`
	EntryStoreEndpoint string `json:"entry_store_endpoint"`
	Redis              string `json:"redis"`
	RedisPwd           string `json:"redis_pwd"`
}
