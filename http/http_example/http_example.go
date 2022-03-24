// Package http_example Package http_example by chenaiquan<like.aiquan@qq.com> create on 2021/9/30 17.58
package http_example

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hello", helloHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// handler echoes r.URL.Path
func indexHandler(w http.ResponseWriter, req *http.Request) {
	_, err := fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	if err != nil {
		return
	}
}

// handler echoes r.URL.Header
func helloHandler(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		_, err := fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		if err != nil {
			return
		}
	}
}
