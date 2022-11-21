package config

import (
	"log"
	"os"
	"r-booker/common"
)

type ApiConfig struct {
	common.Config

	AuthService string
}

func New() ApiConfig {
	as := os.Getenv("AUTH_SERVICE")

	if as == "" {
		log.Fatal("env variable AUTH_SERVICE is not set")
	}

	return ApiConfig{
		Config:      common.NewConfigFromFlags(),
		AuthService: as,
	}
}
