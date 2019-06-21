package part2

import (
	"fmt"
	"strconv"
	"testing"
)

func Test(t *testing.T) {
	//array()

	//slice()
	//arr := [5]int{1,2,3,4,5}
	slice := []int{1, 2, 3, 4, 5}
	//fmt.Println(slice)		//[1 2 3 4 5]
	//reverse(slice)
	//fmt.Println(slice)		//[5 4 3 2 1]

	//modifyArr(arr)
	//fmt.Println(arr)
	//modifySlice(slice)
	//fmt.Println(slice)
	fmt.Printf("slice cap is %d\n", cap(slice)) //slice cap is 5
	slice1 := append(slice, 9)
	fmt.Printf("slice1 cap is %d\t%p\t%p\n", cap(slice1), slice1, slice) //slice1 cap is 10	0xc0000520a0	0xc000078030
	slice2 := append(slice, 10)
	fmt.Printf("slice2 cap is %d\t%p\t%p\n", cap(slice2), slice2, slice) //slice2 cap is 10	0xc0000520f0	0xc000078030
	slice3 := append(slice2, 10)
	fmt.Printf("slice3 cap is %d\t%p\t%p\n", cap(slice3), slice3, slice2) //slice3 cap is 10	0xc0000520f0	0xc0000520f0
	slice4 := append(slice2, 6, 7, 7, 5, 0)
	fmt.Printf("slice4 cap is %d\t%p\t%p\n", cap(slice4), slice4, slice2) //slice4 cap is 20	0xc0000b2000	0xc0000520f0
}

func array() {

	o := [3]int{1, 2, 3}

	p := [3]int{1, 2, 3}

	q := [4]string{1: "a", 2: "2", 3: "21321"}

	r := [4]string{3: "123"}

	for _, i := range o {
		fmt.Println(i)
	}
	for _, i := range p {
		fmt.Println(i)
	}
	for i, v := range q {
		fmt.Println(strconv.Itoa(i) + "-----" + v)
	}
	for i, v := range r {
		fmt.Println(strconv.Itoa(i) + "-----" + v)
	}

	//数组的比较，如果长度不想同，IDE会直接报错，无法比较，在长度相同的情况下则会比较数组元素的值是否相等
	//并且顺序要一致 ，如 [...]int{1,2,3}与[...]int{1,2,3}是相等的，但是与[...]itn{2,1,3}则是不相等的
	fmt.Println(o == p, r == q)
}

func slice() {
	month := [...]string{1: "Jan", 2: "Feb", 3: "Mar", 4: "Apr", 5: "May", 6: "Jun",
		7: "Jul", 8: "Aug", 9: "Sep", 10: "Oct", 11: "Nov", 12: "Dec"}
	for i, v := range month {
		fmt.Println(strconv.Itoa(i) + "-----" + v)
	}

	fmt.Println(cap(month), len(month))

	spring := month[1:4]

	Q2 := month[4:7]

	fmt.Println(cap(spring), len(spring))
	fmt.Println(cap(Q2), len(Q2))
}

func reverse(n []int) {
	for i, j := 0, len(n)-1; i < j; i, j = i+1, j-1 {
		n[i], n[j] = n[j], n[i]
	}
}

func modifyArr(n [5]int) {
	new := n
	new[0] = 100
}

func modifySlice(n []int) {
	new := n
	new[0] = 100
}

func nonempty(n []string) []string {
	i := 0
	for _, v := range n {
		if v != "" {
			n[i] = v
			i++
		}
	}
	return n[:i]
}
func TestNonempty(t *testing.T) {
	s := []string{"a", "", "c"}
	var s1 = nonempty(s)
	fmt.Println(s)
	fmt.Println(s1)
}
