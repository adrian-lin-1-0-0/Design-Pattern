package pubsub

import (
	"sync"
)

type Subscriber struct {
	Topic     string
	MessageCh chan *Message
}

type Hub struct {
	subscribers map[string][]*Subscriber
	mu          sync.Mutex
	bufferLen   int
}

type HubOptions struct {
	BufferLen int
}

func NewHub(ho *HubOptions) *Hub {
	bufferLen := 10
	if ho != nil && ho.BufferLen > 0 {
		bufferLen = ho.BufferLen
	}
	return &Hub{
		subscribers: make(map[string][]*Subscriber),
		bufferLen:   bufferLen,
	}
}

func (h *Hub) Subscribe(topic string) *Subscriber {
	sub := &Subscriber{
		Topic:     topic,
		MessageCh: make(chan *Message, h.bufferLen),
	}

	h.mu.Lock()
	defer h.mu.Unlock()

	h.subscribers[topic] = append(h.subscribers[topic], sub)
	return sub
}

func (h *Hub) Unsubscribe(sub *Subscriber) {
	h.mu.Lock()
	defer func() {
		h.mu.Unlock()
		close(sub.MessageCh)
	}()

	subs := h.subscribers[sub.Topic]
	for i, s := range subs {
		if s == sub {
			subs = append(subs[:i], subs[i+1:]...)
			h.subscribers[sub.Topic] = subs
			break
		}
	}
}

func (h *Hub) Publish(topic string, payload interface{}) {

	msg := &Message{
		Topic:   topic,
		Payload: payload,
	}

	h.mu.Lock()
	defer h.mu.Unlock()

	subs := h.subscribers[topic]
	for _, sub := range subs {
		go func(sub *Subscriber) {
			select {
			case sub.MessageCh <- msg:
			default:
			}
		}(sub)
	}
}
