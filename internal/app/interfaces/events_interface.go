package interfaces

type EventsBusInterface interface {
	Subscribe(topic string) <-chan interface{}
	Publish(topic string, data interface{})
}
