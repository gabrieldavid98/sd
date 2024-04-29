package gateway

import (
	"context"
	pb "idl/gen/master/v1"
	"log"
	"os"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func RegisterNode() error {
	conn, err := grpc.Dial(
		"master:7000",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		return err
	}

	defer conn.Close()

	var (
		master  = pb.NewNodeServiceClient(conn)
		host    = os.Getenv("HOST")
		port, _ = strconv.Atoi(os.Getenv("PORT"))
		req     = &pb.AddNodeRequest{
			Host: host,
			Port: int32(port),
		}
	)

	r, err := master.AddNode(context.Background(), req)
	if err != nil {
		return err
	}

	log.Printf("master::NodeService::AddNode response: %v", r)
	return nil
}
