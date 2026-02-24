package kudomodels

import "sync"

var (
	ID uint64
	mu sync.Mutex
)

func nextIdGenerator() {
	mu.Lock()
	defer mu.Unlock()

	ID++
}
