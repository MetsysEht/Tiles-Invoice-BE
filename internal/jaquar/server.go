package jaquar

// Server has methods implementing of server rpc.
type Server struct {
	manager IManager
}

// NewServer returns a server.
func NewServer(manager IManager) *Server {
	return &Server{manager: manager}
}
