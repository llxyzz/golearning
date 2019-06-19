package part2

import (
	"fmt"
	"strconv"
	"testing"
)

func TestMapt(t *testing.T) {
	mapt()
}

func mapt() {

	m := make(map[string]int)

	ages := map[string]int{
		"a": 20,
		"b": 19,
	}

	m["a"] = 20

	for k, v := range ages {
		fmt.Println(k + "------->" + strconv.Itoa(v))
	}

}
