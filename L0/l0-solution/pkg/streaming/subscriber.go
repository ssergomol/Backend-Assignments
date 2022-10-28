package streaming

import (
	"backend-assignments/l0/pkg/models"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/nats-io/stan.go"
)

type JSONstructure struct {
	models.Order
	Delivery models.Delivery `json:"delivery"`
	Payment  models.Payment  `json:"payment"`
	Items    []models.Item   `json:"items"`
}

func GetDataFromChannel() {
	const (
		clusterID = "test-cluster"
		clientID  = "subscriber"
	)

	conn, _ := stan.Connect(clusterID, clientID)
	defer conn.Close()

	wg := sync.WaitGroup{}
	wg.Add(1)
	channel := make(chan []byte)

	sub, _ := conn.Subscribe("channel", func(m *stan.Msg) {
		fmt.Printf("Received a message: %v\n", string(m.Data))
		channel <- m.Data
		wg.Done()
	}, stan.StartWithLastReceived())
	var data []byte
	data = <-channel
	UnmarshalJSON(data)

	wg.Wait()

	sub.Unsubscribe()
}

func UnmarshalJSON(rawData []byte) {
	data := JSONstructure{}
	json.Unmarshal(rawData, &data)
	data.Delivery.OrderUID = data.OrderUID
	data.Payment.OrderUID = data.OrderUID
	for i := range data.Items {
		data.Items[i].OrderUID = data.OrderUID
	}
}
