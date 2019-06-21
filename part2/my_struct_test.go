package part2

import (
	"fmt"
	"testing"
	"time"
)

type Employee struct {
	ID        string
	Name      string
	Address   string
	Salary    int
	DoB       time.Time
	Position  string
	ManagerId string
}

var emp Employee

func TestMyStruct(t *testing.T) {

	emp.Address = "广州"

	position := &emp.Position
	*position = "高级工程师"

	var empOfNow *Employee = &emp

	empOfNow.Name = "tom"

	v := []int{1, 2, 23, 4, 45, 5, 31}
	fmt.Println(v[:0])

}

type tree struct {
	value       int
	left, right *tree
}

//通过二叉树实现的插入排序
func Sort(values []int) {
	var root *tree
	for _, value := range values {
		root = add(root, value)
	}
	appendValues(values[:0], root)
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {

	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}

	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func TestSort(t *testing.T) {
	values := []int{88, 221, 11, 2, 3, 1, 99, 12}
	Sort(values)
	fmt.Println(values)
}

type Pointer struct {
	X, Y int
}

type Circle struct {
	Pointer
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func TestS(t *testing.T) {
	var w Wheel
	w.X = 5
	w.Y = 5
	w.Radius = 10
	w.Spokes = 4
	fmt.Printf("%v\n", w)

}
