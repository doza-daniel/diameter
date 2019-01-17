package convexhull

import (
	"github.com/doza-daniel/diameter/point"
)

// ConvexHull ...
func ConvexHull(pts point.Points) point.Points {
	if pts.Len() < 3 {
		return point.Points{}
	}

	pts.Sort(sortByY())

	// all points are colinear
	if pts.Find(diffY(pts.Pts[0].Y)) == -1 ||
		pts.Find(diffX(pts.Pts[0].X)) == -1 {
		return point.Points{}
	}

	angleSorted := point.Points{
		Pts: pts.Pts[pts.Find(diffY(pts.Pts[0].Y)):],
	}

	angleSorted.Sort(sortByAngle(pts.Pts[0]))
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
