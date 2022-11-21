package server

import (
	"context"
	"log"
	"r-booker/api/config"
	"r-booker/proto/service/auth"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ApiServer struct {
}

func (s *ApiServer) Start(conf config.ApiConfig) {
	conn, err := grpc.Dial(conf.AuthService, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Can't connect: %v", err)
	}

	c := auth.NewAuthClient(conn)

	defer conn.Close()

	res, errl := c.Login(context.TODO(), &auth.LoginRequest{Username: " ", Password: " "})

	if errl != nil {
		// TODO add logger
		log.Fatalf("Can't login: %v", errl)
	}

	log.Println(res)
}
