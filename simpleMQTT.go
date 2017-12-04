//Package simpleMQTT Simplifies paho.mqtt.golang
package simpleMQTT

import (
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

//NewMqtt gives new mqtt
func NewMqtt(url string, handler Handler) (Mqtt, error) {
	var m = Mqtt{URL: url}
	var err error
	m.on = handler
	err = m.connect()
	if err != nil {
		return Mqtt{}, err
	}
	err = m.subscribe("#")
	if err != nil {
		return Mqtt{}, err
	}
	return m, nil
}

//Handler function type
type Handler func(m *Mqtt, topic, msg string)

//Mqtt client
type Mqtt struct {
	URL    string
	token  MQTT.Token
	client MQTT.Client
	on     Handler
}

func (m *Mqtt) messageHandler(client MQTT.Client, msg MQTT.Message) {
	if m.on != nil {
		m.on(m, msg.Topic(), string(msg.Payload()))
	}
}

//Connect opens connection
func (m *Mqtt) connect() error {
	opts := MQTT.NewClientOptions().AddBroker(m.URL)
	opts.SetClientID("simpleMQTT" + time.Now().Format(time.UnixDate))
	opts.SetDefaultPublishHandler(m.messageHandler)
	m.client = MQTT.NewClient(opts)
	m.token = m.client.Connect()
	m.token.Wait()
	return m.token.Error()
}

//subscribe subscribes to topic
func (m *Mqtt) subscribe(topic string) error {
	m.token = m.client.Subscribe(topic, 0, nil)
	m.token.Wait()
	return m.token.Error()
}

//Publish publishes message to broker
func (m *Mqtt) Publish(topic, message string) error {
	m.token = m.client.Publish(topic, 0, false, message)
	m.token.Wait()
	return m.token.Error()
}

//Destroy unsubscribes and closes the connection
func (m *Mqtt) Destroy() error {
	m.token = m.client.Unsubscribe("go-mqtt/sample")
	m.token.Wait()
	m.client.Disconnect(250)
	return m.token.Error()
}
