package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	load()
}

func load() {
	for i := 0; i < 100000; i++ {
		resp, err := http.Get("http://localhost:8000")
		if err != nil {
			log.Fatalln(err)
		}
		defer resp.Body.Close()

		body, _ := ioutil.ReadAll(resp.Body)
		log.Println(i, string(body))
	}
}
