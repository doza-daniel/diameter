package point

import (
	"fmt"
	"math"
	"sort"
)

// Point ...
type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type sortable struct {
	pts  []Point
	less func(Point, Point) bool
}

// Sort ...
func Sort(pts []Point, less func(Point, Point) bool) {
	s := sortable{pts, less}
	sort.Sort(s)
}

// Find ...
func Find(pts []Point, f func(Point) bool) int {
	if f == nil {
		return -1
	}

	for i, pt := range pts {
		if f(pt) {
			return i
		}
	}
	return -1
}

// Count ...
func Count(pts []Point, f func(Point) bool) int {
	if f == nil {
		return 0
	}

	cnt := 0
	for _, pt := range pts {
		if f(pt) {
			cnt++
		}
	}
	return cnt
}

// Distance ...
func Distance(a, b Point) float64 {
	return math.Sqrt((a.X-b.X)*(a.X-b.X) + (a.Y-b.Y)*(a.Y-b.Y))
}

func (s sortable) Len() int {
	return len(s.pts)
}

func (s sortable) Swap(i, j int) {
	s.pts[i], s.pts[j] = s.pts[j], s.pts[i]
}

func (s sortable) Less(i, j int) bool {
	if s.less == nil {
		return false
	}
	return s.less(s.pts[i], s.pts[j])
}

func (p Point) String() string {
	return fmt.Sprintf("(%.2f, %.2f)", p.X, p.Y)
}
