package part0

import (
	"fmt"
	"strconv"
)

const final = 13213

func Test() {
	//var c1 = final
	//var c2 = 1111

	//fmt.Printf("%d====%d\n",c1,c2)

	//x := 2
	//p := &x
	//fmt.Println(p)			//0xc000054190
	//fmt.Println(*p)			//2
	//fmt.Println(x)			//2
	//*p = 100
	//fmt.Println(p)			//0xc000054190
	//fmt.Println(*p)			//100
	//fmt.Println(x)			//100
	//fmt.Println(x==*p)		//true

	//var p *int					//初始化指针变量，零值为nil
	//fmt.Println(p) 				//<nil>
	//fmt.Println(p == nil)		//true
	//
	//var x, y string
	//fmt.Println(&x,&y)			//0xc000032bc0 0xc000032bd0 	x,y的零值为"",但是分配了;不同的内存地址
	//fmt.Println(&x == &y,&x == nil, &y == nil)		//false false false
	//
	//fmt.Println(f())			//0xc000054198
	//fmt.Println(f())			//0xc0000541c0
	//fmt.Println(f() == f())		//false

	p := new(int)
	fmt.Println(p) //0xc000054190
	*p = 1000
	fmt.Println(*p) //1000

	fmt.Printf("最大公约数为%v\n", gys(16, 24))

	fmt.Printf("斐波那契数列的第%v个数是%v\n", 15, fib(15))

	s := "13213"
	fmt.Println(strconv.Atoi(s))

}

/*
计算斐波那契数列的第N个数
*/
func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}

/*计算最大公约数*/
func gys(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func f() *int {
	y := 1
	return &y
}

var global *int

func a() {
	var x int
	x = 1
	global = &x
}

func b() {
	x := new(int)
	*x = 1
}
