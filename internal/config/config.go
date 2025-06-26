package config

import (
	"os"
)

type APIConfig struct {
	Host   string
	Port   string
	MLHost string
}

func Load() (*APIConfig, error) {
	host := os.Getenv("BE_HOST")
	port := os.Getenv("BE_PORT")
	mlHost := os.Getenv("ML_HOST")
	return &APIConfig{
		Host:   host,
		Port:   port,
		MLHost: mlHost,
	}, nil

}
