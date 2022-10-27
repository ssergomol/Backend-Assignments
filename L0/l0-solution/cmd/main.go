package main

import (
	"backend-assignments/l0/pkg/apiserver"
	"backend-assignments/l0/pkg/streaming"
	"log"
)

func main() {
	// Retrieve data from streaming server
	streaming.PublishDataToChannel()
	streaming.GetDataFromChannel()

	config := apiserver.NewConfig()
	server := apiserver.NewServer(config)

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}

}
