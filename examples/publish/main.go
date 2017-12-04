package main

import (
	"os"

	"github.com/jmu0/simpleMQTT"
)

func main() {
	args := os.Args[1:]
	if len(args) == 3 {
		m, _ := simpleMQTT.NewMqtt("tcp://"+args[0], nil)
		defer m.Destroy()
		m.Publish(args[1], args[2])
	} else {
		panic("<program> <host:port> <topic> <message>")
	}
}
