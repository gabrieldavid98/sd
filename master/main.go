package main

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	limit := 9e7
	pp := partitions(limit)
	ch := make(chan float64, len(pp))

	var wg sync.WaitGroup
	wg.Add(len(pp))

	for _, p := range pp {
		go func(p partition, wg *sync.WaitGroup) {
			defer wg.Done()

			var sum float64
			for n := p[0]; n <= p[1]; n++ {
				sum += math.Pow(-1, n) / (2*n + 1)
			}

			ch <- sum
		}(p, &wg)
	}

	wg.Wait()
	close(ch)

	var total float64
	for sum := range ch {
		total += sum
	}

	fmt.Println(total * 4)
}

type partition [2]float64

func partitions(limit float64) []partition {
	var pp []partition
	per := limit * 0.1 // 5%

	for n := 0.0; n < limit; n += per + 1 {
		var p partition
		p[0] = n
		p[1] = n + per

		if p[1] > limit {
			p[1] = limit
		}

		pp = append(pp, p)
	}

	return pp
}
