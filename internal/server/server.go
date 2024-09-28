package server

import (
	"github.com/MetsysEht/Tiles-Invoice-BE/internal/config"
	"github.com/MetsysEht/Tiles-Invoice-BE/pkg/logger"
	"google.golang.org/grpc"
	"net"
)

type Server struct {
	config     config.NetworkInterfaces
	grpcServer *grpc.Server

	//internalServer *http.Server
	//httpServer     *http.Server
}

func NewServer(config config.NetworkInterfaces) *Server {
	grpcServer := NewGrpcServer()
	return &Server{
		config:     config,
		grpcServer: grpcServer,
	}
}

func (s *Server) Start() {
	s.startGrpcServer()
}

func (s *Server) startGrpcServer() {
	listener, err := net.Listen("tcp", s.config.GrpcServerAddress)
	if err != nil {
		panic(err)
	}

	err = s.grpcServer.Serve(listener)
	if err != nil {
		panic(err)
	}

	logger.Sl.Infow("server started", "address", s.config.GrpcServerAddress)
}

func NewGrpcServer() *grpc.Server {
	grpcServer := grpc.NewServer(grpc.ChainUnaryInterceptor(getInterceptors()...))
	registerGRPCHandlers(grpcServer)
	return grpcServer
}
