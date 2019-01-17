package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/doza-daniel/diameter/convexhull"
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
	p1, p2, err := diameter(points)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Two most distant points: %s and %s\n", p1, p2)
}

func diameter(points []point.Point) (point.Point, point.Point, error) {
	fmt.Println("Before sort:")
	fmt.Printf("%+v\n", points)
	hull := convexhull.ConvexHull(points)
	fmt.Println("After sort:")
	fmt.Printf("%+v\n", hull)
	return point.Point{}, point.Point{}, nil
}
