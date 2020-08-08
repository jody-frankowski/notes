package main

import "sync"

type Counter struct {
	mu    sync.Mutex
	value int
}

func NewCounter() *Counter {
	return &Counter{}
}

func (counter *Counter) Inc() {
	counter.mu.Lock()
	counter.value++
	counter.mu.Unlock()

}

func (counter *Counter) Value() int {
	return counter.value
}
