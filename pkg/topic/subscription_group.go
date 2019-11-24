package topic

import (
	"sync"
)

// SubscriptionGroup represents a group of subscribers
// that divides the messages between them. Each message
// is consumed by one subscriber in the group.
type SubscriptionGroup struct {
	isClosed            func() bool
	inbox               chan interface{}
	subscriberWaitGroup *sync.WaitGroup
}

// AddSubscriber returns false if the topic is closed and otherwise it
// starts a new subscriber goroutine and returns true. The subscriber's
// close method is called when the topic is closed.
func (g *SubscriptionGroup) AddSubscriber(inboxLen int, subscriber Subscriber) bool {
	g.subscriberWaitGroup.Add(1)

	if g.isClosed() {
		g.subscriberWaitGroup.Done()
		return false
	}

	go func(inbox chan interface{}) {
		defer g.subscriberWaitGroup.Done()
		defer subscriber.Close()

		for message := range inbox {
			subscriber.Consume(message)
		}
	}(g.inbox)

	return true
}

func (g *SubscriptionGroup) Close() error {
	close(g.inbox)
	return nil
}

func (g *SubscriptionGroup) Consume(value interface{}) {
	g.inbox <- value
}
