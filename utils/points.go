package utils

type Point struct {
	X, Y int
}

func (p Point) Up() Point {
	return Point{X: p.X, Y: p.Y - 1}
}

func (p Point) Down() Point {
	return Point{X: p.X, Y: p.Y + 1}
}

func (p Point) Right() Point {
	return Point{X: p.X + 1, Y: p.Y}
}

func (p Point) Left() Point {
	return Point{X: p.X - 1, Y: p.Y}
}
