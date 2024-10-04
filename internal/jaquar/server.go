package jaquar

import (
	"context"

	v1 "github.com/MetsysEht/Tiles-Invoice-BE/rpc/jaquar/v1"
)

// Server has methods implementing of server rpc.
type Server struct {
	v1.UnimplementedJaquarServer
	manager IManager
}

// NewServer returns a server.
func NewServer(manager IManager) *Server {
	return &Server{manager: manager}
}

func (s *Server) GetProductDetails(_ context.Context, req *v1.GetProductRequest) (*v1.JaquarProduct, error) {
	prod, err := s.manager.GetProductDetails(req.GetSeries(), req.GetColorCode(), req.GetCodeNumber())
	if err != nil {
		return nil, err
	}
	return &v1.JaquarProduct{
		Description: prod.Description,
		Nrp:         prod.NRP,
		Mrp:         prod.MRP,
	}, nil
}
