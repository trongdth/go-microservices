package config

import (
	"os"
	"strconv"
)

var config *Config

func init() {
	env := os.Getenv("env")
	db := os.Getenv("db")
	port, _ := strconv.Atoi(os.Getenv("port"))

	config = &Config{
		Environment: env,
		Db:          db,
		Port:        port,
	}
}

func GetConfig() *Config {
	return config
}

// Config : struct
type Config struct {
	Environment string `json:"env"`
	Port        int    `json:"port"`
	Db          string `json:"db"`
}
