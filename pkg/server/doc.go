// Package server is an attempt attempt at abstracting out the
// client-server pattern, much like gen_server in Erlang.
//
// A drawback with this package is its use of interface{},
// removing type safety. That could be addressed using a generator.
package server
