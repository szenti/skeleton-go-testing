package shapes

type Square struct {
	side float64
}

func (s Square) Area() float64 {
	return s.side * s.side
}

func NewSquare(side float64) *Square {
	return &Square{side: side}
}
