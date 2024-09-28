package server

import (
	"github.com/MetsysEht/Tiles-Invoice-BE/internal/config"
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
	return &Server{}
}

func (s *Server) Start() {

}

func (s *Server) StartGrpcServer() {
	listener, err := net.Listen("tcp", s.config.GrpcServerAddress)
	if err != nil {
		panic(err)
	}

	err = s.grpcServer.Serve(listener)
	if err != nil {
		panic(err)
	}
}
