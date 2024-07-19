package wterminal

import "github.com/eclipse/paho.golang/paho"

type Event struct {
	EventCode int    `json:"eventCode" form:"eventCode"`
	Number    string `json:"number" form:"number"`
	Value     string `json:"value" form:"value"`
}

type MqttTerminal interface {
	Run()
	Stop()
	GetDevcieSn() string
	GenerateClientId() string
	GenerateOnlieTopic(clientID string) string
	GenerateOnlineMessage(clientID string) string
	GenerateLastWillMessage(clientID string) string
	GenerateLastWillTopic(clientID string) string
	GetReceivedChan() chan *paho.Publish
	GenerateOnClientError(err error)
	GenerateSubscribeOptions() []paho.SubscribeOptions
	Separated() bool
}
