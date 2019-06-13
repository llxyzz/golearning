package part0

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func RunWebServer() {
	//handler := func(w http.ResponseWriter, r *http.Request) {
	//
	//}
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request)  {
	lissajous(w)
}

/*func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "%s <---> %s <---> %s\n", r.Method, r.URL, r.Proto)

	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%s]=%s\n", k, v)
	}

	fmt.Fprintf(w, "Host=%s\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr=%s\n", r.RemoteAddr)

	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	for k, v := range r.Form {
		fmt.Fprintf(w,"Form[%s]=%s", k, v)
	}
}
*/
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "访问次数是：%d\n", count)
	mu.Unlock()
}
