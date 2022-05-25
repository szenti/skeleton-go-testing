package rectangle

type Rectangle struct {
	sideA float64
	sideB float64
}

func (r Rectangle) Area() float64 {
	return r.sideA * r.sideB
}

func NewRectangle(sideA, sideB float64) *Rectangle {
	return &Rectangle{
		sideA: sideA,
		sideB: sideB,
	}
}
