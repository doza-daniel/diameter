package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/doza-daniel/diameter/convexhull"
	"github.com/doza-daniel/diameter/point"
)

func main() {
	var pointSet []point.Point
	pointSet = []point.Point{
		point.Point{X: 1, Y: 1},
		point.Point{X: 2, Y: 3},
		point.Point{X: 2, Y: -1},
		point.Point{X: 2, Y: 3},
		point.Point{X: 2, Y: 3},
		point.Point{X: 2, Y: -1},
		point.Point{X: 1, Y: -1},
		point.Point{X: 2, Y: 3},
	}
	if err := json.NewEncoder(os.Stdout).Encode(pointSet); err != nil {
		log.Fatal(err)
	}
	p1, p2, err := diameter(pointSet)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Two most distant points: %s and %s\n", p1, p2)
}

func diameter(pointSet []point.Point) (point.Point, point.Point, error) {
	fmt.Println("Before sort:")
	fmt.Printf("%+v\n", pointSet)
	hull := convexhull.ConvexHull(point.Points{Pts: pointSet})
	fmt.Println("After sort:")
	fmt.Printf("%+v\n", hull.Pts)
	return point.Point{}, point.Point{}, nil
}
