Counter library
---

Simple atomic counter. It lets you count arbitrary quantity of values and do something with them.

Usage example
---

```go
package main

import "github.com/aiseeq/counter"

func main() {
	// Init queues and makes map
	c := counter.New()
	// Increment key in order of queue
	c.Inc("test_key")
	// Execute func in order of queue
	c.Do(func(d counter.Data) {
		// Do something with Data atomically
		print(d["test_key"])
	})
}
```