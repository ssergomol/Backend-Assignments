package main

import (
	"fmt"
	"log"

	stan "github.com/nats-io/stan.go"
)

func main() {
	const (
		clusterID = "test-cluster"
		clientID  = "client-0"
	)

	conn, err := stan.Connect(clusterID, clientID)
	if err != nil {
		log.Fatal(err)
	}

	sub, err := conn.Subscribe("test-channel", func(m *stan.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	}, stan.StartWithLastReceived())

	// err = conn.Publish("test-channel", []byte("test message 2"))
	if err != nil {
		log.Fatal(err)
	}

	sub.Unsubscribe()
	conn.Close()
}
