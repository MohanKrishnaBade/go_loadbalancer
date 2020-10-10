# go_loadbalancer
GoLang gateway that distributes the incoming traffic to the BE servers.

# how to use
1. specify your BE/FE server url's in service.yaml
2. start the load balancer just by runnig `go run main.go` --default port would be 8000, you can also allowed to pass the port number as an argument.
3. we are using the Round Robin technique to distribute the load among all the servers on a cyclical basis
#
--you can add or remove the servers in service.yaml at any given of time ----
# how to test
1. I've created a simple bakend server, You can spin up mutiple BE servers just by passing portnumber as an arguments.
   1. ` cd backendservers && go run main.go :8081 server1`
   1. ` cd backendservers && go run main.go :8082 server2`
2. run load balancer server 
3. add BE server url's in the server.yaml file
4. you can test it by hitting the loader blancer URL mutiple times.

# make commands to simpliy the process even further
1. `make load_server` -> to start the load balancer
2. `make server1 && make server2 && ...` -> to spin up the BE servers on different ports
3. `make load_test` to perform a load test against the load balancer.

# sample output
![Alt text](asserts/img1.png?raw=true "Title")

   


