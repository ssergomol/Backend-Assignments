package main

import (
	"backend-assignments/l0/pkg/apiserver"
	"backend-assignments/l0/pkg/streaming"
	"sync"

	"github.com/sirupsen/logrus"
)

func main() {
	// Retrieve data from streaming server
	streaming.PublishDataToChannel()
	data := streaming.GetDataFromChannel()

	config := apiserver.NewConfig()
	server := apiserver.NewServer(config)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		if err := server.Start(); err != nil {
			logrus.Fatal(err)
		}
		wg.Done()
	}()

	server.UpdateDatabase(data)
	wg.Wait()

}
