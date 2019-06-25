package part3

import (
	"fmt"
	"golang.org/x/net/html"
)

//Go函数的返回值变量能被提前声明，并且作用于整个函数的区块内
func f(x, y int) (z int, a string) {
	z = x + y
	a = "OK"
	return
}

//获取页面的a标签的链接
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

func outline(stack []string, n *html.Node) {

	if n.Type == html.ElementNode {
		stack = append(stack, n.Data)
		fmt.Println(stack)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}

var depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		depth++
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}

//遍历节点之前，通过前置和后置函数处理添加空格
//pre函数和end函数都是可选的
func forEachElement(n *html.Node, pre, end func(n *html.Node)) {

	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachElement(c, pre, end)
	}

	if end != nil {
		end(n)
	}
}
