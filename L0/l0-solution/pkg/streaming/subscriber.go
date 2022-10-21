package streaming

import (
	"fmt"
	"sync"

	"github.com/nats-io/stan.go"
)

func GetDataFromChannel() {
	const (
		clusterID = "test-cluster"
		clientID  = "subscriber"
	)

	conn, _ := stan.Connect(clusterID, clientID)
	defer conn.Close()

	wg := sync.WaitGroup{}
	wg.Add(1)

	sub, _ := conn.Subscribe("channel", func(m *stan.Msg) {
		fmt.Printf("Received a message: %v\n", string(m.Data))
		wg.Done()
	}, stan.StartWithLastReceived())

	wg.Wait()
	sub.Unsubscribe()
}
