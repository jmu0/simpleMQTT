package main

import (
	"log"
	"os"
	"time"

	"github.com/jmu0/simpleMQTT"
)

func handler(m *simpleMQTT.Mqtt, topic, message string) {
	log.Println("TOPIC:", topic, ", MESSAGE:", message)
}
func main() {
	args := os.Args[1:]
	if len(args) == 1 {
		m, _ := simpleMQTT.NewMqtt("tcp://"+args[0], handler)
		m.Publish("subscriber", "I am listening!")
		for {
			time.Sleep(10 * time.Second)
			m.Publish("subscriber", "I am still alive!")
		}
	} else {
		panic("<program> <host:port>")
	}

}
