package main

import (
	"encoding/json"
	"flag"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/doza-daniel/diameter/point"
)

func main() {
	upper := flag.Int("u", 100, "upper bound")
	lower := flag.Int("l", 0, "lower bound")
	num := flag.Int("n", 10, "number of points")
	flag.Parse()

	rand.Seed(time.Now().Unix())

	gen := func(low, high int) []int {
		var numbers []int
		for i := low; i <= high; i++ {
			numbers = append(numbers, i)
		}
		rand.Shuffle(len(numbers), func(i, j int) {
			numbers[i], numbers[j] = numbers[j], numbers[i]
		})
		return numbers
	}

	xs := gen(*lower, *upper)
	ys := gen(*lower, *upper)

	points := make([]point.Point, 0)
	for i := 0; i < *num; i++ {
		points = append(
			points,
			point.Point{X: float64(xs[i%len(xs)]), Y: float64(ys[i%len(ys)])},
		)
	}

	if err := json.NewEncoder(os.Stdout).Encode(points); err != nil {
		log.Fatal(err)
	}
}
