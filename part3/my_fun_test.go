package part3

import (
	"fmt"
	"golang.org/x/net/html"
	"golearning/part1"
	"strings"
	"testing"
)

func TestOne(t *testing.T) {

	fmt.Println(f(1, 1))

	f1 := square
	fmt.Printf("%T\n", f1)
	fmt.Println(f1(2))

	f2 := show
	fmt.Printf("%T\n", f2)
	fmt.Println(f2(3, 12))

}

func TestTwo(t *testing.T) {
	f3 := square2()

	fmt.Println(f3()) //1
	fmt.Println(f3()) //4
	fmt.Println(f3()) //9
	fmt.Println(f3()) //16
	fmt.Println(f3()) //25

}

func square(n int) int {
	return n * n
}

func show(m, n int) int {
	return m + n
}

func square2() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func TestRecursive(t *testing.T) {
	var url = []string{"https://www.sina.cn"}
	s := part1.FetchUrl(url)

	doc, err := html.Parse(strings.NewReader(s))
	if err != nil {
		fmt.Printf("findlinks1: %v\n", err)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

func TestOutline(t *testing.T) {
	var url = []string{"https://www.sina.cn"}
	s := part1.FetchUrl(url)
	doc, err := html.Parse(strings.NewReader(s))
	if err != nil {
		fmt.Printf("findlinks1: %v\n", err)
	}
	outline(nil, doc)
}

func TestOutline2(t *testing.T) {
	var url = []string{"https://www.baidu.cn"}
	s := part1.FetchUrl(url)
	doc, err := html.Parse(strings.NewReader(s))
	if err != nil {
		fmt.Printf("findlinks1: %v\n", err)
	}

	forEachElement(doc, startElement, endElement)

}

func TestDeferred1(t *testing.T) {
	a, b := 1, 2

	defer func() {
		fmt.Println(a, b)

		a = a * 3
		b = 0

		fmt.Println(a, b)

	}()

	a, b = 10, 20

	fmt.Printf("----->%d\t%d\n", a, b)

}

func TestDeferred2(t *testing.T) {
	a, b := 1, 2

	defer func(a, b int) {
		fmt.Println(a, b)

		a = a * 3
		b = 0

		fmt.Println(a, b)

	}(a, b)

	a, b = 10, 20

	fmt.Printf("----->%d\t%d\n", a, b)

}
func TestDeferred3(t *testing.T) {
	a, b := 1, 2

	defer sh(a, b)

	a, b = 10, 20

	fmt.Printf("----->%d\t%d\n", a, b)

}

func sh(a, b int) {
	fmt.Println(a, b)

	a = a * 3
	b = 0

	fmt.Println(a, b)

}

func TestClosure(t *testing.T) {
	f2 := f1
	fmt.Println(f2(1))
	fmt.Println("---------")
	fmt.Println(f2(1))

}

func f1(i int) (b int) {
	b = i + 1
	p := func(c int) (l int) {
		l = 9
		l += c
		return
	}(b)
	b++
	b = p + 1
	fmt.Println(b)
	fmt.Println(p)
	return
}

func TestFetchUrl(t *testing.T) {
	url := "https://www.qq.com"
	n := fetchUrl(url)
	fmt.Printf("网页的大小：%d\n", n)
}
