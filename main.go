package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

var lastServedIndex int

type Server struct {
	Url string `json:"url"`
}

func (s *Server) IsHealthy() bool {
	if res, err := http.Head(s.Url); err == nil {
		return res.StatusCode == http.StatusOK
	} else {
		return false
	}
}

func (s *Server) ReverseProxy() *httputil.ReverseProxy {

	serverUrl, err := url.Parse(s.Url)
	if err != nil {
		log.Fatalln(err.Error())
	}
	return httputil.NewSingleHostReverseProxy(serverUrl)
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		availableServers := yamlReader()
		for lastServedIndex < len(availableServers) {
			server := availableServers[lastServedIndex]
			if server.IsHealthy() {
				server.ReverseProxy().ServeHTTP(w, r)
				counter(&lastServedIndex, len(availableServers))
				return
			}
			counter(&lastServedIndex, len(availableServers))
		}
	})

	fmt.Printf("load balancer server started at port %s", os.Args[1])
	log.Fatalln(http.ListenAndServe(os.Args[1], nil))

}

func yamlReader() []*Server {
	var servers map[string][]*Server
	yamlFile, err := ioutil.ReadFile("./servers.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &servers)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return servers["servers"]
}

func counter(lastServedIndex *int, totalServersCount int) {
	if *lastServedIndex == totalServersCount-1 {
		*lastServedIndex = 0
	} else {
		*lastServedIndex++
	}
}
