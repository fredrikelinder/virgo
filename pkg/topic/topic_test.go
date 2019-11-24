package topic

import (
	"fmt"
	"sort"
)

type subscriberSpy struct {
	messages   []interface{}
	closeCount int
}

func (s *subscriberSpy) Consume(message interface{}) {
	s.messages = append(s.messages, message)
}

func (s *subscriberSpy) Close() error {
	s.closeCount++
	return nil
}

func ExampleTopic() {
	t := New()

	s := &subscriberSpy{}
	fmt.Println("AddSubscriber:", t.AddSubscriber(0, s))

	p, _ := t.NewPublisher(0)
	p.Publish("message:0")
	p.Publish("message:1")
	pie(p.Close())

	pie(t.Close())

	fmt.Println("messages:", s.messages)
	fmt.Println("closeCount:", s.closeCount)

	// Output:
	// AddSubscriber: true
	// messages: [message:0 message:1]
	// closeCount: 1
}

func ExampleTopic_Close() {
	t := New()

	pie(t.Close())

	p, ok := t.NewPublisher(0)
	fmt.Println("NewPublisher:", p, ok)

	s := &subscriberSpy{}
	fmt.Println("AddSubscriber:", t.AddSubscriber(0, s))
	fmt.Println("closeCount:", s.closeCount)

	// Output:
	// NewPublisher: <nil> false
	// AddSubscriber: false
	// closeCount: 0
}

func ExampleSubscriptionGroup() {
	t := New()

	s0 := &subscriberSpy{}
	s1 := &subscriberSpy{}
	s2 := &subscriberSpy{}

	g, ok := t.NewSubscriptionGroup(0)
	if !ok {
		fmt.Println("no subscription group")
		return
	}
	fmt.Println("AddSubscriber:", g.AddSubscriber(0, s0))
	fmt.Println("AddSubscriber:", g.AddSubscriber(0, s1))
	fmt.Println("AddSubscriber:", t.AddSubscriber(0, s2))

	p, ok := t.NewPublisher(0)
	if !ok {
		fmt.Println("no publisher")
		return
	}
	p.Publish("message:1")
	p.Publish("message:2")
	p.Publish("message:3")
	p.Publish("message:4")
	p.Publish("message:5")

	pie(p.Close())
	pie(t.Close())

	// we must merge messages from s0 and s1 and sort them
	// to get a consistent test output
	var messages []interface{}
	messages = append(messages, s0.messages...)
	messages = append(messages, s1.messages...)
	sort.Slice(messages, func(i, j int) bool {
		return messages[i].(string) < messages[j].(string)
	})
	fmt.Println("messages:", messages)

	// s2 receives them in send order
	fmt.Println("messages:", s2.messages)

	fmt.Println("closeCount:", s0.closeCount)
	fmt.Println("closeCount:", s1.closeCount)
	fmt.Println("closeCount:", s2.closeCount)

	// Output:
	// AddSubscriber: true
	// AddSubscriber: true
	// AddSubscriber: true
	// messages: [message:1 message:2 message:3 message:4 message:5]
	// messages: [message:1 message:2 message:3 message:4 message:5]
	// closeCount: 1
	// closeCount: 1
	// closeCount: 1
}

func pie(err error) {
	if err != nil {
		fmt.Println("err:", err)
	}
}
