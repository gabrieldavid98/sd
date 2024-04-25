package client

import (
	pb "idl/gen/master/v1"

	"google.golang.org/grpc"
)

func RegisterNode() {
	conn, err := grpc.Dial(":7000")

	master := pb.NewNodeServiceClient(conn)

	master.AddNode()
}
