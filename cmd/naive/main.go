package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

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
	var pt1, pt2 point.Point
	for _, a := range points {
		for _, b := range points {
			if point.Distance(a, b) > point.Distance(pt1, pt2) {
				pt1, pt2 = a, b
			}
		}
	}

	return pt1, pt2
}
