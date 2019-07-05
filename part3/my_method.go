package part3

type Sub struct {
	X, Y float64
}

func (s Sub) getSubtraction(t Sub) float64 {
	return s.X - t.Y
}
