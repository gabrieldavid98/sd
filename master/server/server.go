package server

import (
	"context"
	pb "idl/gen/master/v1"
	"master/registry"
)

type Server = pb.NodeServiceServer

type server struct {
	pb.NodeServiceServer

	r *registry.Registry
}

func New(r *registry.Registry) Server {
	return &server{
		r: r,
	}
}

// AddNode ...
func (s *server) AddNode(
	ctx context.Context,
	r *pb.AddNodeRequest,
) (*pb.AddNodeResponse, error) {
	s.r.Add(r.GetHost())
	panic("")
}
