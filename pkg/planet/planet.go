package planet

type Planet interface {
	X() int
	Y() int
}

type Mars struct {
	xSize int
	ySize int
}

func NewMars(x, y int) Mars {
	return Mars{
		xSize: x,
		ySize: y,
	}
}

func (m *Mars) X() int {
	return m.xSize
}

func (m *Mars) Y() int {
	return m.ySize
}
