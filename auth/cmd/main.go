package main

import (
	"r-booker/auth/server"
	"r-booker/common"
)

func main() {
	conf := common.NewConfigFromFlags()

	s := server.AuthService{}

	s.Start(conf)
}
