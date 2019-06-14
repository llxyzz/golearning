package part0

import "fmt"

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
	fmt.Println(p)				//0xc000054190
	*p = 1000
	fmt.Println(*p)				//1000


}


func f() *int {
	y := 1
	return &y
}

