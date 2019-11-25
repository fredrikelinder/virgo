package server

import "io"

type subscriber struct {
	state interface{}
}

func (s *subscriber) Consume(message interface{}) {
	fn := message.(func(interface{}))
	fn(s.state)
}

func (s *subscriber) Close() error {
	state, ok := s.state.(io.Closer)
	if ok {
		return state.Close()
	}

	return nil
}
