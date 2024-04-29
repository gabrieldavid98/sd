package transport

import (
	"log"
	"master/data"
	"net"

	pb "idl/gen/master/v1"

	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
)

func RunGrpcServer(registry *data.Registry) {
	lis, err := net.Listen("tcp", ":7000")

	if err != nil {
		log.Fatalf("failed connection: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterNodeServiceServer(s, newNodeServiceServer(registry))

	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}

func RunHttpServer(
	app *fiber.App,
	registry *data.Registry,
) {
	httpServer := newHttpServer(app, registry)
	httpServer.registerRoutes()

	log.Fatal(app.Listen(":7070"))
}
