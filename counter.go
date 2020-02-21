// Package counter lets you atomically count arbitrary quantity of values and do something with them
package counter

import "sync"

type Data map[string]uint32
type Counter struct {
	data Data
	inc  chan string
	do   chan func(Data)
	wg   sync.WaitGroup
}

// worker listens two queues and either increments counter, or executes function for gathered data
func (c *Counter) worker() {
	for {
		select {
		case key := <-c.inc:
			c.data[key]++
		case f := <-c.do:
			f(c.data)
		}
		c.wg.Done()
	}
}

// Inc increments counter for key
func (c *Counter) Inc(key string) {
	c.wg.Add(1)
	c.inc <- key
}

// Do executes function for Data
func (c *Counter) Do(f func(Data)) {
	c.wg.Add(1)
	c.do <- f
}

// Wait can be used to wait while queue will become empty
func (c *Counter) Wait() {
	c.wg.Wait()
}

// New returns new counter ready to use
func New() *Counter {
	c := Counter{
		data: make(Data),
		inc:  make(chan string, 100),
		do:   make(chan func(Data), 100),
	}
	go c.worker()
	return &c
}
