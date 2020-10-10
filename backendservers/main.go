package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", HelloServer)
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "test route called form  %s", os.Args[2])
	})
	fmt.Printf("server started at port %s", os.Args[1])
	_ = http.ListenAndServe(os.Args[1], nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "<h1>Hello, %s!<h1>", os.Args[2])
}
