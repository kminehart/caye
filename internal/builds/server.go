package builds

import "github.com/kminehart/caye"

type Server struct {
	builds map[string]caye.Build
}

func NewServer() *Server {
	return &Server{
		builds: make(map[string]caye.Build),
	}
}
