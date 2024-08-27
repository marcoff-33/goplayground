package main

import (
	"fmt"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			http.ServeFile(w, r, "static/index.html")
		} else {
			fmt.Fprintf(w, "Hello from %s\n", r.URL.Path)
		}
	})
	
	
	http.ListenAndServe("127.0.0.1:8000", nil)
}