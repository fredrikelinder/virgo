package topic

import "io"

// Subscriber consumes messages published to the topic.
type Subscriber interface {
	io.Closer

	// Consume consumes a published message.
	Consume(message interface{})
}
