package gateway

import (
	"context"
	pb "idl/gen/node/v1"
	"master/data"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Node interface {
	Compute(ctx context.Context, start, end float64) (*ComputeResponse, error)
}

type nodeGateway struct {
	client pb.LeibnizPiServiceClient
	node   *data.Node
}

func NewNode(registry *data.Registry) (Node, error) {
	node, err := createNodeGrpcClient(registry)
	if err != nil {
		return nil, err
	}

	return &nodeGateway{
		client: node.Client,
		node:   node,
	}, nil
}

func (n *nodeGateway) Compute(
	ctx context.Context,
	start, end float64,
) (*ComputeResponse, error) {
	req := &pb.ComputeRequest{
		Start: start,
		End:   end,
	}

	r, err := n.client.Compute(ctx, req)
	if err != nil {
		return nil, err
	}

	return &ComputeResponse{
		DownstreamService: n.node.Addr(),
		Sum:               r.GetSum(),
		ExecutionTimeMs:   r.GetExecutionTime(),
	}, nil
}

func createNodeGrpcClient(
	registry *data.Registry,
) (*data.Node, error) {
	node := registry.Get()
	if node.Client != nil {
		return node, nil
	}

	conn, err := grpc.Dial(
		node.Addr(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		return nil, err
	}

	client := pb.NewLeibnizPiServiceClient(conn)
	node.Client = client

	return node, nil
}
