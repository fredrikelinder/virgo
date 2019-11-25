package server

import (
	"github.com/fredrikelinder/virgo/pkg/topic"
)

// New returns a new server with state equal to the
// return value of the given factory.
func New(factory func() interface{}) *Server {
	t := topic.New()
	t.AddSubscriber(1, &subscriber{state: factory()})
	return &Server{topic: t}
}

// Server represents a server.
type Server struct {
	topic *topic.Topic
}

// Async evaluates the given function using the state of the server
// as its only argument, it may return before the function has been evaluated.
func (s *Server) Async(fn func(interface{})) {
	s.topic.Publish(fn)
}

// Sync evaluates the given function using the state of the server
// as its only argument, it return after the function has been evaluated.
func (s *Server) Sync(fn func(interface{}) interface{}) interface{} {
	ch := make(chan interface{}, 1)
	s.Async(func(message interface{}) {
		defer close(ch)
		ch <- fn(message)
	})
	return <-ch
}
