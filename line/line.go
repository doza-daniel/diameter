package line

import (
	"fmt"
	"math"

	"github.com/doza-daniel/diameter/point"
)

// Line ...
type Line struct {
	A float64
	B float64
	C float64
}

// ThroughPoints derive the general line equation ax + by + c = 0, using this:
// (y - y1) / (y2 - y1) = (x - x1) / (x2 - x1)
func ThroughPoints(p, q point.Point) Line {
	a := -(q.Y - p.Y)
	b := q.X - p.X
	c := p.X*(q.Y-p.Y) - p.Y*(q.X-p.X)
	return Line{A: a, B: b, C: c}
}

// DistanceFromPoint ...
func (l Line) DistanceFromPoint(p point.Point) float64 {
	return math.Abs(l.A*p.X+l.B*p.Y+l.C) / math.Sqrt(l.A*l.A+l.B*l.B)
}

func (l Line) String() string {
	return fmt.Sprintf("%.2f * x + %.2f * y = %.2f", l.A, l.B, -l.C)
}
