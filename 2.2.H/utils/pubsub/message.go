package pubsub

type Message struct {
	Topic   string
	Payload interface{}
}
