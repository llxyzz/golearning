package part3

import (
	"testing"
)

//Go函数的返回值变量能被提前声明，并且作用于整个函数的区块内
func f(x, y int) (z int, a string) {
	z = x + y
	a = "OK"
	return
}

func TestOne(t *testing.T) {

	//fmt.Println(f(1,1))

	/*if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}*/
}
