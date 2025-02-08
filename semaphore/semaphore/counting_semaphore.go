package semaphore

import (
	"context"
	"errors"
	"sync"
)

type CountingSemaphore struct {
	mu    sync.Mutex
	count int64
	limit int64
}

func NewCountingSemaphore(limit int64) Semaphore {
	return &CountingSemaphore{limit: limit}
}

func (c *CountingSemaphore) Acquire(ctx context.Context, n int64) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.count+n > c.limit {
		return errors.New("semaphore limit exceeded")
	}
	c.count += n
	return nil
}

func (c *CountingSemaphore) TryAcquire(n int64) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.count+n > c.limit {
		return false
	}
	c.count += n
	return true
}

func (c *CountingSemaphore) Release(n int64) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.count >= n {
		c.count -= n
	}
}
