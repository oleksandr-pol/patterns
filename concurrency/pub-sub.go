package concurrency

import (
	"errors"
	"sync"
)

type PubSub struct {
	mu     sync.RWMutex
	closed bool
	subs   map[string][]chan string
}

func CreatePubSub() *PubSub {
	ps := &PubSub{}

	subs := make(map[string][]chan string)
	ps.subs = subs

	return ps
}

func (ps *PubSub) Subscribe(topic string) <-chan string {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	ch := make(chan string, 2)

	ps.subs[topic] = append(ps.subs[topic], ch)

	return ch
}

func (ps *PubSub) Publish(topic string, msg string) error {
	ps.mu.RLock()
	defer ps.mu.RUnlock()

	if ps.closed {
		return errors.New("connection is closed")
	}

	// will mutex work in go routines?
	for _, ch := range ps.subs[topic] {
		go func(ch chan string) {
			ch <- msg
		}(ch)
	}

	return nil
}

func (ps *PubSub) Close() {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	if !ps.closed {
		ps.closed = true
		for _, subs := range ps.subs {
			for _, ch := range subs {
				close(ch)
			}
		}
	}
}
