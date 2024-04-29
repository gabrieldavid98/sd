package data

import (
	"fmt"
	pb "idl/gen/node/v1"
)

type Nodes []*Node

type Node struct {
	Host   string
	Port   int32
	Client pb.LeibnizPiServiceClient
}

func (n *Node) Addr() string {
	return fmt.Sprintf("%s:%d", n.Host, n.Port)
}
