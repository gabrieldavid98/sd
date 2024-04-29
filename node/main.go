package main

import (
	"log"
	"node/gateway"
	"node/transport"
	"time"
)

func main() {
	time.Sleep(time.Second * 10)
	err := gateway.RegisterNode()
	if err != nil {
		log.Fatalln(err)
	}

	transport.RunGrpcServer()
}
