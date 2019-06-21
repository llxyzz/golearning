package part2

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"testing"
	"time"
)

const IssueUrl = "https://api.github.com/search/issues"

type IssueSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Item
}

type Item struct {
	Id        int
	Title     string
	Number    int
	HTMLURL   string `json:"html_url"`
	State     string
	CreatedAt time.Time `json:"created_at"`
	Score     float64
	User      *User
	Body      string
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

func SearchIssue(key []string) (*IssueSearchResult, error) {

	query := url.QueryEscape(strings.Join(key, " "))

	res, err := http.Get(IssueUrl + "?q=" + query)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		res.Body.Close()
		log.Fatalf("search is failed:%v", err)
		return nil, fmt.Errorf("search is failed : %s", res.Status)
	}

	var result IssueSearchResult

	//读取输入流数据，并且将其存储在&result指向的值中
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		res.Body.Close()
		return nil, err
	}

	res.Body.Close()
	return &result, err
}

func TestSearch(t *testing.T) {
	key := []string{"go", "java"}

	result, err := SearchIssue(key)
	if err != nil {
		log.Fatalf("the error is %v", err)
	}

	fmt.Println("the total result is :" + strconv.Itoa(result.TotalCount))

	//data, err := json.MarshalIndent(result, "", "	")
	//if err != nil {
	//	log.Fatalf("JSON marshaling failed: %s", err)
	//}
	//fmt.Printf("%s\n", data)

	for _, v := range result.Items {
		fmt.Printf("%s\n", v.HTMLURL)
	}

	const template = `the total issue is : {{.TotalCount}}
		{{range .Items}}------------
		

			
		`

}
