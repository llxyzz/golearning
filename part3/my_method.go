package part3

import (
	"image/color"
	"sync"
)

type Num struct {
	X, Y float64
}

type ColorNum struct {
	Num
	Color color.RGBA
}

type Person struct {
	name string
	age  int
}

func (s Num) getSubtraction(t Num) float64 {
	return s.X - t.Y
}

func (s *Num) getPlus(t Num) float64 {
	return s.X + t.Y
}

var cache = struct {
	sync.Mutex
	mapping map[string]string
}{}

func Lookup(key string) string {
	cache.Lock()
	v := cache.mapping[key]
	cache.Unlock()
	return v
}
