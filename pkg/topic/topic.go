package topic

import (
	"sync"
)

// Topic represents a unidirectional fan-in and fan-out
// communication channel.
type Topic struct {
	queue               chan interface{}
	closing             chan struct{}
	publisherWaitGroup  sync.WaitGroup
	subscriberWaitGroup sync.WaitGroup
	subscriberCs        []chan interface{}
	m                   sync.Mutex
}

// New returns a new Topic.
func New() *Topic {
	t := &Topic{
		queue:   make(chan interface{}, 1),
		closing: make(chan struct{}),
	}

	go func(queue <-chan interface{}) {
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

// NewPublisher return false if the topic is closed and otherwise it returns
// a new publisher and true.
func (t *Topic) NewPublisher(outboxLen int) (*Publisher, bool) {
	t.publisherWaitGroup.Add(1)

	if t.IsClosed() {
		t.publisherWaitGroup.Done()
		return nil, false
	}

	outbox := make(chan interface{}, outboxLen)

	// one goroutine to pull messages out of the outbox into the queue of the subscriberWaitGroup
	go func(queue chan<- interface{}, outbox chan interface{}) {
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

	inbox := make(chan interface{}, inboxLen)

	// one goroutine to pull from inbox
	go func(inbox <-chan interface{}) {
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
		inbox:               make(chan interface{}, inboxLen),
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

func (t *Topic) getSubscriberCs() []chan interface{} {
	t.m.Lock()
	defer t.m.Unlock()
	return t.subscriberCs
}

func (t *Topic) addSubscriberC(subscriberC chan interface{}) {
	t.m.Lock()
	defer t.m.Unlock()
	t.subscriberCs = append(t.subscriberCs, subscriberC)
}