package topic

import (
	"io"
	"sync"
)

// template type Topic(Message)
type Message = interface{}

// Topic represents a unidirectional fan-in and fan-out
// communication channel.
//
// Publishers publish messages to the topic. Messages published
// by the same publisher are added to the topic in the same order
// they were published.
//
// Subscribers consume messages published to the topic. All subscribers
// receive the same messages in the same order.
//
// Subscribers within a subscription group distribute the messages
// across the subscribers in the group, all receiving different messages.
type Topic struct {
	queue               chan Message
	closing             chan struct{}
	publisherWaitGroup  sync.WaitGroup
	subscriberWaitGroup sync.WaitGroup
	subscriberCs        []chan Message
	m                   sync.Mutex
}

// New returns a new Topic.
func New() *Topic {
	t := &Topic{
		queue:   make(chan Message, 1),
		closing: make(chan struct{}),
	}

	go func(queue <-chan Message) {
		defer func() {
			subscriberCs := t.getSubscriberCs()
			for _, subscriberC := range subscriberCs {
				close(subscriberC)
			}
		}()

		for message := range queue {
			subscriberCs := t.getSubscriberCs()
			for _, subscriberC := range subscriberCs {
				subscriberC <- message
			}
		}
	}(t.queue)

	return t
}

// Publish returns false if the topic is closed and otherwise it
// publishes the message and returns true.
func (t *Topic) Publish(message Message) bool {
	t.publisherWaitGroup.Add(1)
	defer t.publisherWaitGroup.Done()

	if t.IsClosed() {
		return false
	}

	t.queue <- message

	return true
}

// NewPublisher return false if the topic is closed and otherwise it returns
// a new publisher and true.
func (t *Topic) NewPublisher(outboxLen int) (*Publisher, bool) {
	t.publisherWaitGroup.Add(1)

	if t.IsClosed() {
		t.publisherWaitGroup.Done()
		return nil, false
	}

	outbox := make(chan Message, outboxLen)

	// one goroutine to pull messages out of the outbox into the queue of the subscriberWaitGroup
	go func(queue chan<- Message, outbox chan Message) {
		defer t.publisherWaitGroup.Done()

		for message := range outbox {
			queue <- message
		}
	}(t.queue, outbox)

	return &Publisher{outbox: outbox}, true
}

// AddSubscriber returns false if the topic is closed and otherwise it
// starts a new subscriber goroutine and returns true. The subscriber's
// close method is called when the topic is closed.
func (t *Topic) AddSubscriber(inboxLen int, subscriber Subscriber) bool {
	t.subscriberWaitGroup.Add(1)

	if t.IsClosed() {
		t.subscriberWaitGroup.Done()
		return false
	}

	inbox := make(chan Message, inboxLen)

	// one goroutine to pull from inbox
	go func(inbox <-chan Message) {
		defer t.subscriberWaitGroup.Done()
		defer subscriber.Close()

		for message := range inbox {
			subscriber.Consume(message)
		}
	}(inbox)

	t.addSubscriberC(inbox)

	return true
}

// NewSubscriptionGroup returns false if the topic is closed and otherwise it
// returns a subscription group and true.
func (t *Topic) NewSubscriptionGroup(inboxLen int) (*SubscriptionGroup, bool) {
	g := &SubscriptionGroup{
		isClosed:            t.IsClosed,
		inbox:               make(chan Message, inboxLen),
		subscriberWaitGroup: &t.subscriberWaitGroup,
	}

	ok := t.AddSubscriber(inboxLen, g)
	if !ok {
		return nil, false
	}

	return g, true
}

// Close closes the topic and returns nil.
func (t *Topic) Close() error {
	close(t.closing)
	t.publisherWaitGroup.Wait()

	close(t.queue)
	t.subscriberWaitGroup.Wait()

	return nil
}

// CloseC returns a channel that is closed when the topic is closed.
func (t *Topic) CloseC() <-chan struct{} {
	closeC := make(chan struct{})

	go func(closing <-chan struct{}) {
		defer close(closeC)
		for range closing {
		}
	}(t.closing)

	return closeC
}

// IsClosed returns true if the topic has been closed, and false otherwise.
func (t *Topic) IsClosed() bool {
	select {
	case <-t.closing:
		return true
	default:
		return false
	}
}

func (t *Topic) getSubscriberCs() []chan Message {
	t.m.Lock()
	defer t.m.Unlock()
	return t.subscriberCs
}

func (t *Topic) addSubscriberC(subscriberC chan Message) {
	t.m.Lock()
	defer t.m.Unlock()
	t.subscriberCs = append(t.subscriberCs, subscriberC)
}

// Publisher represents a publisher.
type Publisher struct {
	outbox chan Message
}

// Publish publishes the given value to the topic.
func (p *Publisher) Publish(message Message) {
	p.outbox <- message
}

// Close closes the publisher.
func (p *Publisher) Close() error {
	close(p.outbox)
	return nil
}

// Subscriber consumes messages published to the topic.
type Subscriber interface {
	io.Closer

	// Consume consumes a published message.
	Consume(message Message)
}

// SubscriptionGroup represents a group of subscribers
// that divides the messages between them. Each message
// is consumed by one subscriber in the group.
type SubscriptionGroup struct {
	isClosed            func() bool
	inbox               chan Message
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

	go func(inbox chan Message) {
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

func (g *SubscriptionGroup) Consume(value Message) {
	g.inbox <- value
}
