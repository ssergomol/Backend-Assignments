package data_getter

import (
	"io"
	"os"

	"github.com/nats-io/stan.go"
)

type Publisher struct {
}

func (p *Publisher) PublishData() {
	const (
		clusterID = "test-cluster"
		clientID  = "publisher"
		dataFile  = "../../../task/model.json"
	)

	jsonFile, _ := os.Open(dataFile)
	defer jsonFile.Close()

	bytesData, _ := io.ReadAll(jsonFile)

	conn, _ := stan.Connect(clusterID, clientID)
	conn.Publish("channel", bytesData)
	conn.Close()
}
