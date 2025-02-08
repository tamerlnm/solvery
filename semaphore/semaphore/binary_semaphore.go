package semaphore

import (
	"context"
	"errors"
	"sync"
)

type BinarySemaphore struct {
	mu sync.Mutex
}

func NewBinarySemaphore() Semaphore {
	return &BinarySemaphore{}
}

func (b *BinarySemaphore) Acquire(ctx context.Context, n int64) error {
	if n != 1 {
		return errors.New("binary semaphore only supports n = 1")
	}
	b.mu.Lock()
	return nil
}

func (b *BinarySemaphore) TryAcquire(n int64) bool {
	if n != 1 {
		return false
	}
	return b.mu.TryLock()
}

func (b *BinarySemaphore) Release(n int64) {
	if n != 1 {
		return
	}
	b.mu.Unlock()
}
