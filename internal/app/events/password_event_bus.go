package events

import "sync"

type EventBus struct {
	subscribers map[string][]chan interface{}
	mu          sync.RWMutex
}

func NewEventBus() *EventBus {
	return &EventBus{
		subscribers: make(map[string][]chan interface{}),
	}
}

func (eb *EventBus) Subscribe(topic string) <-chan interface{} {
	eb.mu.Lock()
	defer eb.mu.Unlock()

	ch := make(chan interface{}, 1)
	eb.subscribers[topic] = append(eb.subscribers[topic], ch)
	return ch
}

func (eb *EventBus) Publish(topic string, data interface{}) {
	eb.mu.RLock()
	defer eb.mu.RUnlock()

	if chans, found := eb.subscribers[topic]; found {
		for _, ch := range chans {
			go func(c chan interface{}) {
				c <- data
			}(ch)
		}
	}
}
