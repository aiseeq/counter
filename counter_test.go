package counter

import (
	"testing"
)

func TestNew(t *testing.T) {
	c := New()
	if c == nil {
		t.Error("new failed")
	}
}

func TestCounter_Inc(t *testing.T) {
	c := New()
	c.Inc("test_key")
	c.Wait()
	if c.data["test_key"] != 1 {
		t.Error("inc failed")
	}
}

func TestCounter_Do(t *testing.T) {
	c := New()
	c.Do(func(d Data) {
		d["test_key"] = 100
	})
	c.Wait()
	if c.data["test_key"] != 100 {
		t.Error("do failed")
	}
}

func ExampleCounter_Inc() {
	// Inits queues and makes map
	c := New()
	// Will increment in order of queue
	c.Inc("test_key")
	// Will execute in order of queue
	c.Do(func(d Data) {
		// Do something with Data atomically
		print(d["test_key"])
	})
}
