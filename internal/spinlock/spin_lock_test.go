package spinlock

import (
	"sync"
	"sync/atomic"
	"testing"
)

func TestSpinLock(t *testing.T) {
	type LockTest struct {
		acquired atomic.Bool
	}

	lock := SpinLock{}
	lockTest := LockTest{}

	// event signalling thread1 has acquired the lock
	waitAcquired := sync.WaitGroup{}
	waitAcquired.Add(1)

	// event signalling test-1 conditions have been set
	// - thread1 has acquired lock,
	// - thread2 has attempted to acquire lock.
	test1Ready := sync.WaitGroup{}
	test1Ready.Add(1)

	// after testing blocking, complete this barrier to continue with test2
	test1Done := sync.WaitGroup{}
	test1Done.Add(1)

	// event signalling test-2 conditions have been set
	// - thread1 has acquired lock
	// - thread2 has attempted to acquire lock
	// - thread1 has released lock
	// - thread2 has acquired lock
	test2Ready := sync.WaitGroup{}
	test2Ready.Add(1)

	// after testing release from blocking, complete this barrier to continue with test3
	test2Done := sync.WaitGroup{}
	test2Done.Add(1)

	// event signalling test-3 conditions have been set
	// - thread1 has acquired lock
	// - thread2 has attempted to acquire lock
	// - thread1 has released lock
	// - thread2 has acquired lock
	// - thread2 has released lock
	test3Ready := sync.WaitGroup{}
	test3Ready.Add(1)

	go func(lock *SpinLock) {
		lock.Acquire()
		waitAcquired.Done()

		// wait for blocking test to complete, then release lock
		test1Done.Wait()
		lock.Release()
	}(&lock)

	go func(lock *SpinLock, lockTest *LockTest) {
		// after thread-1 acquires lock, attempt to aquire lock
		waitAcquired.Wait()
		test1Ready.Done()
		lock.Acquire()

		// record the lock was acquired, this thread is unblocked
		lockTest.acquired.Store(true)
		test2Ready.Done()
		test2Done.Wait()

		// release the lock, we can now test lock is freed
		lock.Release()
		test3Ready.Done()
	}(&lock, &lockTest)

	test1Ready.Wait()
	if !lock.lock.Load() {
		t.Error("Expected lock.Acquire() to set lock from thread1")
		return
	}
	if lockTest.acquired.Load() {
		t.Error("Expected lock.Acquire() to block until it was released. Instead it was acquired")
		return
	}
	test1Done.Done()

	test2Ready.Wait()
	if !lockTest.acquired.Load() {
		t.Error("After releasing lock, expected 2nd thread to stop blocking, and acquire lock")
		return
	}
	test2Done.Done()

	test3Ready.Wait()
	if lock.lock.Load() {
		t.Error("Expected lock to be released once 2nd thread was finished with it")
		return
	}
}
