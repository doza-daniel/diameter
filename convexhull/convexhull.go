package convexhull

import (
	"github.com/doza-daniel/diameter/point"
)

// ConvexHull ...
func ConvexHull(pts []point.Point) []point.Point {
	if len(pts) < 3 {
		return []point.Point{}
	}

	point.Sort(pts, sortByY())

	dx := point.Find(pts, diffX(pts[0].X))
	dy := point.Find(pts, diffY(pts[0].Y))

	// all points are colinear
	if dx == -1 || dy == -1 {
		return []point.Point{}
	}

	point.Sort(pts[dy:], sortByAngle(pts[0]))
	return pts
}

func diffY(y float64) func(point.Point) bool {
	return func(p point.Point) bool {
		return p.Y != y
	}
}

func diffX(x float64) func(point.Point) bool {
	return func(p point.Point) bool {
		return p.X != x
	}
}

func sortByY() func(point.Point, point.Point) bool {
	return func(a, b point.Point) bool {
		return a.Y < b.Y || a.Y == b.Y && a.X < b.X
	}
}

func sortByAngle(origin point.Point) func(point.Point, point.Point) bool {
	return func(a, b point.Point) bool {
		first := -(a.X - origin.X) / (a.Y - origin.Y)
		second := -(b.X - origin.X) / (b.Y - origin.Y)
		return first < second
	}
}

func ccw(a, b, c point.Point) float64 {
	return ((b.X-a.X)*(c.Y-a.Y) - (b.Y-a.Y)*(c.X-a.X))
}
