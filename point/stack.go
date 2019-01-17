package point

// Stack ...
type Stack struct {
	pts []Point
}

// Push ...
func (s *Stack) Push(p Point) {
	s.pts = append(s.pts, p)
}

// Top ...
func (s Stack) Top() Point {
	p := s.pts[len(s.pts)-1]
	return p
}

// NextToTop ...
func (s Stack) NextToTop() Point {
	p := s.pts[len(s.pts)-2]
	return p
}

// Pop ...
func (s *Stack) Pop() Point {
	p := s.Top()
	s.pts = s.pts[:len(s.pts)-1]
	return p
}

// Count ...
func (s Stack) Count() int {
	return len(s.pts)
}

// Points ...
func (s Stack) Points() []Point {
	return s.pts
}
