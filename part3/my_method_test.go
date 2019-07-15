package part3

import (
	"fmt"
	"testing"
	"time"
)

func TestMyMethod(t *testing.T) {

	const hour = time.Hour

	fmt.Println(hour.Seconds())

	n1 := Num{1, 2}

	fmt.Println(n1.getSubtraction(n1))

	n2 := &Num{10, 20}
	fmt.Println(n2.getPlus(n1))

	var cn ColorNum
	cn.X = 1
	fmt.Println(cn)
	cn.Num.X = 2
	fmt.Println(cn)

}

func TestCache(t *testing.T) {
	v := Lookup("123")
	fmt.Println(v)
}
