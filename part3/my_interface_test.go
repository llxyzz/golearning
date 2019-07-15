package part3

import (
	"fmt"
	"testing"
	"time"
)

func TestTest(t *testing.T) {

	fmt.Println(getStr())

	var s = []string{"1", "2", "3"}
	d := DataArray{"1", "2", "3"}

	fmt.Println(s)
	fmt.Println(d)

	now := time.Now()
	var dt = DateTime{}
	fmt.Println(dt.DateFormat(now))
}
