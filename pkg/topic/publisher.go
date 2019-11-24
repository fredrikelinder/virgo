package topic

// Publisher represents a publisher.
type Publisher struct {
	outbox chan interface{}
}

// Publish publishes the given value to the topic.
func (p *Publisher) Publish(message interface{}) {
	p.outbox <- message
}

// Close closes the publisher.
func (p *Publisher) Close() error {
	close(p.outbox)
	return nil
}
