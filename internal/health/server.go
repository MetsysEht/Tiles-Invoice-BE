package health

import (
	"context"
	v1 "github.com/MetsysEht/Tiles-Invoice-BE/rpc/health/v1"
)

// Server has methods implementing of server rpc.
type Server struct {
	v1.UnimplementedHealthServer
	check *Service
}

// NewServer returns a server.
func NewServer(check *Service) *Server {
	return &Server{check: check}
}

func (s *Server) Ping(_ context.Context, _ *v1.PingRequest) (*v1.PingResponse, error) {
	response := &v1.PingResponse{
		Success: true,
		Message: "Pong",
	}

	return response, nil
}
