package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(w, "Hello: %s\n", r.URL.Path)
	})
	http.ListenAndServe(":80", nil)
}
