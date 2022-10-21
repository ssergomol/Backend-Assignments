package streaming

import (
	"io"
	"log"
	"os"

	"github.com/nats-io/stan.go"
)

func PublishDataToChannel() {
	const (
		clusterID = "test-cluster"
		clientID  = "publisher"
		dataFile  = "../../task/model.json"
	)

	jsonFile, err := os.Open(dataFile)
	defer jsonFile.Close()
	if err != nil {
		log.Fatal(err)
	}

	bytesData, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Fatal(err)
	}

	conn, err := stan.Connect(clusterID, clientID)
	defer conn.Close()
	if err != nil {
		log.Fatal(err)
	}

	err = conn.Publish("channel", bytesData)
	if err != nil {
		log.Fatal(err)
	}
}
