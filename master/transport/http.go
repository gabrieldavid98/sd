package transport

import (
	"log"
	"master/data"
	"master/gateway"
	"math"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/sync/semaphore"
)

type httpServer struct {
	app      *fiber.App
	registry *data.Registry
}

func newHttpServer(
	app *fiber.App,
	registry *data.Registry,
) *httpServer {
	return &httpServer{
		app:      app,
		registry: registry,
	}
}

func (h *httpServer) registerRoutes() {
	h.app.Get("/compute-pi-single/:limit", func(c *fiber.Ctx) error {
		l, err := c.ParamsInt("limit", 90_000_000)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).
				JSON(&fiber.Map{
					"msg":   "Invalid Param",
					"error": err.Error(),
				})
		}

		var (
			limit = float64(l)
			start = time.Now()
		)

		var (
			sum  float64
			sign = 1.0
		)
		for n := 0.0; n <= limit; n++ {
			sum += sign / (2*n + 1)
			sign = -sign
		}

		end := time.Since(start)

		return c.JSON(&fiber.Map{
			"result":     sum * 4,
			"durationMs": end.Milliseconds(),
		})
	})

	h.app.Get("/compute-pi-multicore/:limit", func(c *fiber.Ctx) error {
		l, err := c.ParamsInt("limit", 90_000_000)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).
				JSON(&fiber.Map{
					"msg":   "Invalid Param",
					"error": err.Error(),
				})
		}

		var (
			limit = float64(l)
			start = time.Now()
			pp    = partitions(0.0, limit)
			ch    = make(chan float64, len(pp))

			wg sync.WaitGroup
		)

		for _, p := range pp {
			wg.Add(1)

			var (
				start = p[0]
				end   = p[1]
			)

			go func(start, end float64, wg *sync.WaitGroup) {
				defer wg.Done()

				var (
					sum  float64
					sign = math.Pow(-1, start)
				)
				for n := start; n <= end; n++ {
					tempSum := sum
					sum += sign / (2*n + 1)

					if math.IsNaN(sum) {
						sum = tempSum
						break
					}

					sign = -sign
				}

				ch <- sum
			}(start, end, &wg)
		}

		wg.Wait()
		close(ch)

		var total float64
		for sum := range ch {
			total += sum
		}

		end := time.Since(start)

		return c.JSON(&fiber.Map{
			"result":     total * 4,
			"durationMs": end.Milliseconds(),
		})
	})

	h.app.Get("/compute-pi-distributed/:limit", func(c *fiber.Ctx) error {
		l, err := c.ParamsInt("limit", 90_000_000)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).
				JSON(&fiber.Map{
					"msg":   "Invalid Param",
					"error": err.Error(),
				})
		}

		var (
			limit = float64(l)
			start = time.Now()
			pp    = partitions(0.0, limit)
			ch    = make(chan *gateway.ComputeResponse, len(pp))
			sem   = semaphore.NewWeighted(40)
			ctx   = c.Context()

			wg sync.WaitGroup
		)

		for _, p := range pp {
			if err := sem.Acquire(ctx, 1); err != nil {
				log.Printf("Failed to acquire semaphore: %v", err)
				break
			}

			wg.Add(1)

			var (
				start = p[0]
				end   = p[1]
			)

			go func(start, end float64, wg *sync.WaitGroup) {
				defer sem.Release(1)
				defer wg.Done()

				nodeGateway, err := gateway.NewNode(h.registry)
				if err != nil {
					log.Println("Stating NodeGateway error:", err.Error())
					return
				}

				r, err := nodeGateway.Compute(ctx, start, end)
				if err != nil {
					log.Println("NodeGateway::Compute error:", err.Error())
					return
				}

				ch <- r
			}(start, end, &wg)
		}

		wg.Wait()
		close(ch)

		var (
			total       float64
			nodesResult []*gateway.ComputeResponse
		)
		for r := range ch {
			total += r.Sum
			nodesResult = append(nodesResult, r)
		}

		end := time.Since(start)

		return c.JSON(&fiber.Map{
			"result":     total * 4,
			"durationMs": end.Milliseconds(),
			"nodes":      nodesResult,
		})
	})

	h.app.Get("/nodes", func(c *fiber.Ctx) error {
		return c.JSON(&fiber.Map{
			"nodes": h.registry.GetAll(),
		})
	})
}

type partition [2]float64

func partitions(start, end float64) []partition {
	var pp []partition
	per := end * 0.1 // 10%

	for n := start; n < end; n += per + 1 {
		var p partition
		p[0] = n
		p[1] = n + per

		if p[1] > end {
			p[1] = end
		}

		pp = append(pp, p)
	}

	return pp
}
