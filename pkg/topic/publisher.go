package topic

// Publisher represents a publisher.
type Publisher struct {
	outbox chan Value
}

// Publish publishes the given value to the topic.
func (p *Publisher) Publish(message Value) {
	p.outbox <- message
}

// Close closes the publisher.
func (p *Publisher) Close() error {
	close(p.outbox)
	return nil
}
