package main

import (
	"backend-assignments/l0/pkg/streaming"
)

func main() {
	streaming.PublishDataToChannel()
	streaming.GetDataFromChannel()
}
