load_server:
	go run main.go :8000
server1:
	go run backendServers/main.go :8080 server1
server2:
	go run backendServers/main.go :8081 server2
server3:
	go run backendServers/main.go :8082 server3
server4:
	go run backendServers/main.go :8083 server4
load_test:
	go run loadTest/main.go
