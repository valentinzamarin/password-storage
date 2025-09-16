package events

import (
	"testing"
	"time"
)

func TestEventBus_SubscribeAndPublish(t *testing.T) {
	eb := NewEventBus()

	topic := "test_topic"
	ch := eb.Subscribe(topic)

	data := "hello world"

	eb.Publish(topic, data)

	select {
	case msg := <-ch:
		if msg != data {
			t.Errorf("unexpected event data, got %v, want %v", msg, data)
		}
	case <-time.After(time.Second):
		t.Fatal("timeout waiting for event")
	}
}

func TestEventBus_MultipleSubscribers(t *testing.T) {
	eb := NewEventBus()

	topic := "multi_topic"
	ch1 := eb.Subscribe(topic)
	ch2 := eb.Subscribe(topic)

	data := 12345
	eb.Publish(topic, data)

	for i, ch := range []<-chan interface{}{ch1, ch2} {
		select {
		case msg := <-ch:
			if msg != data {
				t.Errorf("subscriber %d: unexpected data, got %v, want %v", i+1, msg, data)
			}
		case <-time.After(time.Second):
			t.Errorf("subscriber %d: timeout waiting for event", i+1)
		}
	}
}
