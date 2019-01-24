package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/doza-daniel/diameter/convexhull"
	"github.com/doza-daniel/diameter/line"
	"github.com/doza-daniel/diameter/point"
)

func main() {
	flag.String("json", "", "JSON array of points")
	flag.Parse()

	file := flag.Arg(0)

	var in *os.File
	if file == "" || file == "-" {
		in = os.Stdin
	} else {
		var err error
		if in, err = os.Open(file); err != nil {
			log.Fatal(err)
		}
	}
	defer in.Close()

	var points []point.Point
	if err := json.NewDecoder(in).Decode(&points); err != nil {
		log.Fatal(err)
	}

	p1, p2 := diameter(points)
	fmt.Printf("Two most distant points: %s and %s\n", p1, p2)
}

func diameter(points []point.Point) (point.Point, point.Point) {
	type candidate struct {
		a, b point.Point
	}
	hull := convexhull.ConvexHull(points)
	candidates := make([]candidate, 0)
	last := 2
	for i, vertex := range hull {
		edge := line.ThroughPoints(vertex, hull[(i+1)%len(hull)])
		for {
			next := (last + 1) % len(hull)
			dLast := edge.DistanceFromPoint(hull[last])
			dNext := edge.DistanceFromPoint(hull[next])
			if dNext < dLast {
				candidates = append(candidates, candidate{vertex, hull[last]})
				break
			}
			last = next
		}
	}

	max := candidates[0]
	for _, c := range candidates[1:] {
		if point.Distance(c.a, c.b) > point.Distance(max.a, max.b) {
			max = c
		}
	}
	return max.a, max.b
}
