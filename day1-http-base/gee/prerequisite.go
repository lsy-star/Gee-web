package gee

import (
	"fmt"
	"net/http"
)

//func main() {
//	http.HandleFunc("/hello", helloHandler)
//	http.HandleFunc("/hi", hiHandler)
//	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
//		fmt.Fprintln(writer, "this is for any request!")
//	})
//
//	log.Fatal(http.ListenAndServe(":9100", nil))
//}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, this is helloHandler!!!")
}

func hiHandler(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		fmt.Fprintln(w, "Header[%q] = %q", k, v)
	}
}
