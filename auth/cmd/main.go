package main

import (
	"r-booker/auth/config"
	"r-booker/auth/server"
)

func main() {
	s := server.AuthServer{}
	conf := config.Config{
		Host: "localhost",
		Port: "8081",
	}

	s.Start(conf)
}
