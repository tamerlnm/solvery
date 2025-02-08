package semaphore

import (
	"context"
	"errors"
	"sync"
)

type MutexSemaphore struct {
	mu sync.Mutex
}

func NewMutexSemaphore() Semaphore {
	return &MutexSemaphore{}
}

func (m *MutexSemaphore) Acquire(ctx context.Context, n int64) error {
	if n != 1 {
		return errors.New("mutex semaphore only supports n = 1")
	}
	m.mu.Lock()
	return nil
}

func (m *MutexSemaphore) TryAcquire(n int64) bool {
	if n != 1 {
		return false
	}
	return m.mu.TryLock()
}

func (m *MutexSemaphore) Release(n int64) {
	if n != 1 {
		return
	}
	m.mu.Unlock()
}
