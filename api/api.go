package api

import (
	"fmt"
	"log"
	"net/http"
	"path"
)

type Vue string

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func api() {
	vueAppHandler := http.FileServer(Vue("app/dist/"))

	http.Handle("/", vueAppHandler)

  http.HandleFunc("/hello", hello)
  http.HandleFunc("/headers", headers)

	log.Println("Listening on port 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func (v Vue) Open(name string) (http.File, error) {
	if ext := path.Ext(name); name != "/" && (ext == "" || ext == ".html") {
		name = "index.html"
	}
	return http.Dir(v).Open(name)
}
