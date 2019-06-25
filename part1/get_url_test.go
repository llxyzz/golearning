package part1

import (
	"fmt"
	"testing"
)

func TestFetchUrl(t *testing.T) {
	var urlArr = []string{"https://www.baidu.com", "https://www.qq.com", "https://www.sina.cn"}
	//var urlArr = []string{"https://www.qq.com"}
	//FetchUrlAll(urlArr)
	fmt.Println(FetchUrl(urlArr))
}
