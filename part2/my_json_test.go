package part2

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

type Book struct {
	IndexNum  int `json:"num"`
	Name      string
	Price     string
	Hardcover bool `json:"favorite,omitempty"`
	Language  []string
}

func TestJson(t *testing.T) {
	//结构体中`json:"num"`是结构体成员Tag,这些tag是在编译阶段关联到该成员的元信息字符串，通常有key:"value"
	//键值对格式组成，在Book结构体中，IndexNum的输出值已经由num代替了，而在HardCover后面的tag中，还多了一个
	//omitempty选项，表示结构体成员为空或者为零值的时候不生成JSON对象，在这里bool类型false为零值

	var books = []Book{
		{IndexNum: 1, Name: "Thinking in Java", Price: "78", Hardcover: false,
			Language: []string{"English", "Chinese", "Japanese"}},
		{IndexNum: 2, Name: "The Go Programing Language", Price: "68", Hardcover: true,
			Language: []string{"English", "Chinese", "Japanese"}},
		{IndexNum: 3, Name: "Effective Java", Price: "59", Hardcover: true,
			Language: []string{"English", "Chinese"}},
	}
	//data, err := json.Marshal(books)
	data, err := json.MarshalIndent(books, "", "	")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)
	//fmt.Println(data)

	//解码json字符串
	//var jsonstr = "[{\"num\":1,\"Name\":\"Thinking in Java\",\"Price\":\"78\",\"Language\":[\"English\",\"Chinese\",\"Japanese\"]},{\"num\":2,\"Name\":\"The Go Programing Language\",\"Price\":\"68\",\"favorite\":true,\"Language\":[\"English\",\"Chinese\",\"Japanese\"]},{\"num\":3,\"Name\":\"Effective Java\",\"Price\":\"59\",\"favorite\":true,\"Language\":[\"English\",\"Chinese\"]}]"

	var bookD []Book
	if err := json.Unmarshal(data, &bookD); err != nil {
		log.Fatalf("JSON unmarshaling failed: %v", err)
	}
	fmt.Println(bookD)
}
