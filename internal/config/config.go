package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type APIConfig struct {
	HostPort string
}

func Load() (*APIConfig, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("error loading .env file")
	}

	hostPort := os.Getenv("HOST_PORT")

	return &APIConfig{
		HostPort: hostPort,
	}, nil

}
