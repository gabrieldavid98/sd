package main

import (
	"master/data"
	"master/transport"

	"github.com/gofiber/fiber/v2"
)

func main() {

	var (
		registry = data.NewRegistry()
		app      = fiber.New()
	)

	go transport.RunGrpcServer(registry)
	transport.RunHttpServer(app, registry)
}
