package transport

import (
	"context"
	pb "idl/gen/master/v1"
	"master/data"
)

type nodeServiceServer struct {
	pb.UnimplementedNodeServiceServer

	r *data.Registry
}

func newNodeServiceServer(r *data.Registry) pb.NodeServiceServer {
	return &nodeServiceServer{
		r: r,
	}
}

// AddNode ...
func (s *nodeServiceServer) AddNode(
	ctx context.Context,
	r *pb.AddNodeRequest,
) (*pb.AddNodeResponse, error) {
	s.r.Add(r.GetHost(), r.GetPort())

	return &pb.AddNodeResponse{
		Ok: true,
	}, nil
}
