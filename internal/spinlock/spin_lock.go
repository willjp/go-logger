package spinlock

import (
	"sync/atomic"
)

// A cheap/fast user-level lock
//
// Make sure to use same instance everywhere you intend to use the same lock.
type SpinLock struct {
	lock atomic.Bool
}

// Acquire the lock,
// or block in the current thread until the lock is released
func (this *SpinLock) Acquire() {
	for {
		// if lock was already true, block until released
		if !this.lock.Swap(true) {
			return
		}
	}
}

// Release the lock
func (this *SpinLock) Release() {
	this.lock.Store(false)
}
