package server

import (
	context "context"
	"fmt"
	"log"
	"net"
	"r-booker/common"
	"r-booker/proto/service/auth"
	"r-booker/proto/service/health"

	"google.golang.org/grpc"
)

type AuthService struct {
	auth.UnimplementedAuthServer
	health.UnimplementedHealthServer
}

// Check implements health.HealthServer
func (*AuthService) Check(context.Context, *health.HealthCheckRequest) (*health.HealthCheckResponse, error) {
	panic("unimplemented")
}

// Watch implements health.HealthServer
func (*AuthService) Watch(*health.HealthCheckRequest, health.Health_WatchServer) error {
	panic("unimplemented")
}

// Login implements service.AuthServer
func (*AuthService) Login(context.Context, *auth.LoginRequest) (*auth.LoginResponse, error) {
	return &auth.LoginResponse{
		AccessToken: "hello-token",
	}, nil
}

func (as *AuthService) Start(conf common.Config) {

	log.Println("Starting server...")

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", conf.Host, conf.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	} else {
		log.Printf("gRPC server is running on %s:%s", conf.Host, conf.Port)
	}

	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	auth.RegisterAuthServer(grpcServer, as)
	health.RegisterHealthServer(grpcServer, as)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
