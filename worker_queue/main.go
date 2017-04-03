package main

import (
	"flag"
	"fmt"
	"net/http"
)

var (
	NWorkers = flag.Int("n", 4, "The number of workers to start")
	HTTPAddr = flag.String("http", "127.0.0.1.3000", "Address to listen for HTTP requests on")
)

func main() {
	//parse the command line flags
	flag.Parse()

	//start the dispathcer
	fmt.Println("Starting the dispatcher")
	StartDispatcher(*NWorkers)

	//register our collector as an HTTP handler function
	fmt.Println("Registering the collector")
	http.HandleFunc("/work", Collector)

	//start the http server
	fmt.Println("HTTP server listening on", *HTTPAddr)
	if err := http.ListenAndServe(*HTTPAddr, nil); err != nil {
		fmt.Println(err.Error())
	}
}

/*
go build -o queued *.go

./queued -n 20

#output with error "listen tcp: address 127.0.0.1.8080: missing port in address"
...
Starting Worker 17
Starting Worker 18
Starting Worker 19
Starting Worker 20
Registering the collector
HTTP server listening on 127.0.0.1.8080
listen tcp: address 127.0.0.1.8080: missing port in address
*/
