package main

import (
	"fmt"
	"golearning/part2"
	"strconv"
)

func main() {

	//var s = "D:/a/b/c/d.go"
	//
	//fmt.Println(part2.BaseName(s))

	var s = 1234567
	//fmt.Println(part2.SplitInt(strconv.Itoa(s)))

	//fmt.Println(part2.IntsToString([]int{1,2,3}))

	fmt.Println(part2.SplitIntByBuffer(strconv.Itoa(s)))

	type Count int

	const (
		FIRST Count = 1 + iota
		SECOND
		THIRD
		FOURTH
		FIFTH
	)

	fmt.Println(FIRST, SECOND, THIRD, FOURTH, FIFTH) //1 2 3 4 5

}
