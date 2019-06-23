package part2

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"testing"
	"text/template"
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

func TestSearchIssue(t *testing.T) {
	key := []string{"go", "java"}

	result, err := SearchIssue(key)
	if err != nil {
		log.Fatalf("the error is %v", err)
	}

	fmt.Println("the total result is :" + strconv.Itoa(result.TotalCount))

	//为了便于输出打印显示，只选取前3条items的数据
	result.Items = result.Items[:3]
	data, err := json.MarshalIndent(result, "", "	")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)

	//for _, v := range result.Items {
	//	fmt.Printf("%s\n", v.HTMLURL)
	//}
}

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

func TestTextTemplate(t *testing.T) {
	const templateText = `the total issue is : {{.TotalCount}}
		{{range .Items}}
		------------
		Numbers:	{{.Number}}
		User:		{{.User.Login}}
		Title:		{{.Title | printf "%.64s"}}
		Age:		{{.CreatedAt | daysAgo}} days ago
		{{end}}`

	var report = template.Must(template.New("issueList").Funcs(template.
		FuncMap{"daysAgo": daysAgo}).Parse(templateText))

	key := []string{"go", "java"}

	result, err := SearchIssue(key)
	//为了便于输出打印显示，只选取前3条items的数据
	result.Items = result.Items[:3]
	if err != nil {
		log.Fatalf("the error is %v", err)
	}

	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatalf("the error is %v", err)
	}

	fmt.Println()
}
