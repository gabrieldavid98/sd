package server

import (
	"context"
	pb "idl/gen/node/v1"
	"math"
	"time"
)

type Server = pb.LeibnizPiServiceServer

type server struct {
	pb.UnimplementedLeibnizPiServiceServer
}

func New() Server {
	return &server{}
}

// Compute ...
func (s *server) Compute(
	ctx context.Context,
	r *pb.ComputeRequest,
) (*pb.ComputeResponse, error) {
	var (
		now = time.Now()
		sum float64
	)

	for n := r.Start; n <= r.End; n++ {
		sum += math.Pow(-1, n) / (2*n + 1)
	}

	end := time.Since(now)

	return &pb.ComputeResponse{
		Sum:           sum,
		ExecutionTime: end.Milliseconds(),
	}, nil
}
