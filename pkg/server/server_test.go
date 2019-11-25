package server

import "fmt"

type xserver struct {
	s *Server
}

func xnew() *xserver {
	factory := func() interface{} { return &xstate{} }
	return &xserver{s: New(factory)}
}

func (x *xserver) add(i int) {
	x.s.Async(async(func(s *xstate) {
		s.add(i)
	}))
}

func (x *xserver) sum() int {
	return x.s.Sync(sync(func(s *xstate) interface{} {
		return s.sum()
	})).(int)
}

func (x *xserver) Close() error {
	fmt.Println("xserver.Close: <nil>")
	return nil
}

type xstate struct {
	value int
}

func (x *xstate) add(i int) {
	x.value += i
}

func (x *xstate) sum() int {
	return x.value
}

func (x *xstate) Close() error {
	return nil
}

func async(fn func(*xstate)) func(interface{}) {
	return func(x interface{}) {
		fn(x.(*xstate))
	}
}

func sync(fn func(*xstate) interface{}) func(interface{}) interface{} {
	return func(x interface{}) interface{} {
		return fn(x.(*xstate))
	}
}

func Example() {
	x := xnew()
	x.add(2)
	fmt.Println("sum:", x.sum())
	pie(x.Close())

	// Output:
	// sum: 2
	// xserver.Close: <nil>
}

func pie(err error) {
	if err != nil {
		fmt.Println("err:", err)
	}
}
