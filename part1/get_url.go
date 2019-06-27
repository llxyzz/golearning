package part1

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func FetchUrlAll(urlArr []string) {

	start := time.Now()

	ch := make(chan string)

	//ch1 := make(chan string)

	for _, url := range urlArr {
		go fetch(url, ch)
		//go fetch(url, ch1)
	}

	for range urlArr {
		fmt.Println("-------->" + <-ch)
	}

	//fmt.Println("===================分割线==========")
	//for range urlArr {
	//	fmt.Println("-------->" + <-ch1)
	//}

	fmt.Printf("花费的总时间%.2fs", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)

	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()

	if err != nil {
		ch <- fmt.Sprintf("当访问%s的时候,%v\n", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.3f----%7d----%s", secs, nbytes, url)
}

func FetchUrl(urls []string) (s string) {

	for _, url := range urls {
		resp, err := http.Get(url)
		if err != nil {
			log.Fatalf("fetch failed : %v\n", err)
		}
		defer resp.Body.Close()
		byte, err := ioutil.ReadAll(resp.Body)
		//byte, err := ioutil.ReadAll(transform.NewReader(resp.Body, simplifiedchinese.HZGB2312.NewEncoder()))
		if err != nil {
			log.Fatalf("fetch reading content failed: %v\n", err)
		}

		//fmt.Printf("%s",byte)

		s += string(byte)
	}
	return
}
