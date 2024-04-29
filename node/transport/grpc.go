package transport

import (
	"context"
	pb "idl/gen/node/v1"
	"log"
	"math"
	"sync"
	"time"
)

type leibnizPiServiceServer struct {
	pb.UnimplementedLeibnizPiServiceServer
}

func newLeibnizPiServiceServer() pb.LeibnizPiServiceServer {
	return &leibnizPiServiceServer{}
}

// Compute ...
func (s *leibnizPiServiceServer) Compute(
	ctx context.Context,
	r *pb.ComputeRequest,
) (*pb.ComputeResponse, error) {
	var (
		start = time.Now()
		pp    = partitions(r.Start, r.End)
		ch    = make(chan float64, len(pp))

		wg sync.WaitGroup
	)

	log.Printf("LeibnizPiServiceServer::Compute range starting on %f and ending on %f\n", r.Start, r.End)

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

	log.Println("LeibnizPiServiceServer::Compute ended")

	res := &pb.ComputeResponse{
		Sum:           total,
		ExecutionTime: end.Milliseconds(),
	}

	return res, nil
}

type partition [2]float64

func partitions(start, end float64) []partition {
	var pp []partition
	per := (end - start) * 0.1 // 10%

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
