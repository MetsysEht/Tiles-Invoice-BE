package server

import (
	"github.com/MetsysEht/Tiles-Invoice-BE/internal/health"
	"github.com/MetsysEht/Tiles-Invoice-BE/internal/jaquar"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	healthv1 "github.com/MetsysEht/Tiles-Invoice-BE/rpc/health/v1"
	jaquarv1 "github.com/MetsysEht/Tiles-Invoice-BE/rpc/jaquar/v1"
)

func registerGRPCHandlers(server *grpc.Server) {
	jaquarRepo := jaquar.NewRepo()
	jaquarManager := jaquar.NewManager(jaquarRepo)
	jaquarServer := jaquar.NewServer(jaquarManager)

	reflection.Register(server)
	healthv1.RegisterHealthServer(server, health.NewServer(health.NewService()))
	jaquarv1.RegisterJaquarServer(server, jaquarServer)
}
