package main

import (
	"fmt"
	"log"
	"net/http"
	"path"
	"strconv"
	"strings"
)

type Frames string

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func (f Frames) Open(name string) (http.File, error) {
	if ext := path.Ext(name); name != "/" && (ext == "" || ext == ".html") {
		name = "test0001.png"
	}
	return http.Dir(f).Open(name)
}

func getFrames(w http.ResponseWriter, req *http.Request) {
	keys, ok := req.URL.Query()["key"]
	if !ok || len(keys[0]) < 1 {
	  fmt.Fprintf(w, "get all frames\n")
		return
	}

	key := keys[0]

	if strings.Contains(key, ",") {
		nums := strings.Split(key, ",")
		if len(nums) > 1 {
			if _, err := strconv.Atoi(nums[1]); err == nil {
				fmt.Fprintf(w, "get frames: " + nums[0] + "-" + nums[1])
				return
			} else {
				fmt.Fprintf(w, "get frames: " + nums[0] + "-end")
				return
			}
		}
	} else if _, err := strconv.Atoi(key); err == nil {
		fmt.Fprintf(w, "get just frame: " + key)
		return
  } else {
		fmt.Fprintf(w, "Key error: " + key)
		return
	}
}

func startServer() {
	framesHandler := http.FileServer(Frames("frames/"))

	http.Handle("/", framesHandler)

	http.HandleFunc("/frames", getFrames)

	log.Println("Server now listening on port 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
