package main

import (
	"r-booker/api/config"
	"r-booker/api/server"
	"r-booker/common"
)

func main() {
	s := server.ApiServer{}
	conf := config.ApiConfig{
		Config: common.Config{
			Host: "localhost",
			Port: "8080",
		},
		AuthService: "localhost:8081",
	}

	s.Start(conf)
}
