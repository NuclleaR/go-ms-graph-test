package server

import (
	context "context"
	"fmt"
	"log"
	"net"
	"r-booker/auth/config"
	"r-booker/auth/server/service"

	"google.golang.org/grpc"
)

type AuthServer struct {
	service.UnimplementedAuthServer
}

// Login implements service.AuthServer
func (*AuthServer) Login(context.Context, *service.LoginRequest) (*service.LoginResponse, error) {
	return &service.LoginResponse{
		Token: "token",
	}, nil
}

func (as *AuthServer) Start(conf config.Config) {
	// flag.Parse()

	log.Println("Starting server...")

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", conf.Host, conf.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	} else {
		log.Printf("listening on tcp %s:%s", conf.Host, conf.Port)
	}

	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	service.RegisterAuthServer(grpcServer, as)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
