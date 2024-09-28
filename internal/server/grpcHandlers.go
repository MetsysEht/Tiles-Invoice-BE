package server

import (
	"github.com/MetsysEht/Tiles-Invoice-BE/internal/health"
	v1 "github.com/MetsysEht/Tiles-Invoice-BE/rpc/health/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func registerGRPCHandlers(server *grpc.Server) {
	reflection.Register(server)
	v1.RegisterHealthServer(server, health.NewServer(health.NewService()))
}
