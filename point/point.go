package point

import (
	"fmt"
	"sort"
)

// Point ...
type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

// Points ...
type Points struct {
	Pts  []Point
	less func(Point, Point) bool
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

// Sort ...
func (p Points) Sort(f func(Point, Point) bool) {
	p.less = f
	sort.Sort(p)
	p.less = nil
}

// Find ...
func (p Points) Find(f func(Point) bool) int {
	if f == nil {
		return -1
	}

	for i, pt := range p.Pts {
		if f(pt) {
			return i
		}
	}
	return -1
}

// Count ...
func (p Points) Count(f func(Point) bool) int {
	if f == nil {
		return 0
	}

	cnt := 0
	for _, pt := range p.Pts {
		if f(pt) {
			cnt++
		}
	}
	return cnt
}

func (p Points) Len() int {
	return len(p.Pts)
}
func (p Points) Swap(i, j int) {
	p.Pts[i], p.Pts[j] = p.Pts[j], p.Pts[i]
}
func (p Points) Less(i, j int) bool {
	if p.less == nil {
		return false
	}
	return p.less(p.Pts[i], p.Pts[j])
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
