# go_loadbalancer
GoLang gateway that distributes the incoming traffic to the BE servers.
# how to use it
1. specify your BE/FE server url's in service.yaml
2. start the load balancer just by runnig `go run main.go` --default port would be 8000, you can also allowed to pass the port number as an argument.
3. we are using the Round Robin technique to distribute the load among all the serverson a cyclical basis
#
--you can add or remove the servers from service.yaml at any given of time ----

