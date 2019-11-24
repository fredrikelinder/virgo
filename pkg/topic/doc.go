// Package topic provides primitives to use the pub/sub pattern.
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
package topic
