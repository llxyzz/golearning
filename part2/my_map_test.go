package part2

import (
	"fmt"
	"sort"
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

func SortMap(m map[string]string) {

	var keys []string
	for k := range m {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	for _, k := range keys {
		fmt.Println("key:" + k + "-------->map:" + m[k])
	}

}

func TestSortMap(t *testing.T) {
	m := map[string]string{
		"d": "13",
		"b": "312",
		"w": "12",
		"f": "12",
	}
	//一般用布尔变量ok来接收map中是否存在key,适用与马上用来判断的场景
	if key, ok := m["d"]; ok {
		fmt.Println(key)
	}
	SortMap(m)
}

func charCount() {
	s := []string{"1", "2", "a", "b"}
	fmt.Println(fmt.Sprintf("%q", s))
}

func TestCharCount(t *testing.T) {

	charCount()
}
