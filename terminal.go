package terminal

// 事件内容由后续再约定，此处先进行声明
type Event struct {
	EventCode int    `json:"eventCode" form:"eventCode"` //指令类别代码
	Number    string `json:"number" form:"number"`       // 具体事件的为标识
	Value     string `json:"value" form:"value"`         //指令中的可用值
}

type MqttTerminal interface {
	// XAuth(data map[string]interface{}) (metadata map[string]interface{}, err error)
	Run()
	Stop()
	GetDevcieSn() string
	GenerateClientId() string
	GenerateOnlieTopic(clientID string) string
	GenerateOnlineMessage(clientID string) string
	GenerateLastWillMessage(clientID string) string
	GenerateLastWillTopic(clientID string) string
	Separated() bool //是否接收发送分离
}

type MqttTerminalSender interface {
	Sende(topic string, payload []byte, qos byte, retained bool, userProperties map[string]string)
}

type MqttTerminalRecHandler func()

type MqttTerminalReceiver interface {
	Receive(topic string, recHandler MqttTerminalRecHandler)
}
