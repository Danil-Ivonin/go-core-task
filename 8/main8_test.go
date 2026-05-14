package main

import (
	"sync/atomic"
	"testing"
	"time"
)

func TestWaitGroupWaitReturnsWhenCounterIsZero(t *testing.T) {
	t.Parallel()

	wg := NewWaitGroup()

	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
	case <-time.After(100 * time.Millisecond):
		t.Fatal("Wait() blocked with zero counter")
	}
}

func TestWaitGroupWaitBlocksUntilDone(t *testing.T) {
	t.Parallel()

	wg := NewWaitGroup()
	wg.Add(1)

	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		t.Fatal("Wait() returned before Done()")
	case <-time.After(20 * time.Millisecond):
	}

	wg.Done()

	select {
	case <-done:
	case <-time.After(100 * time.Millisecond):
		t.Fatal("Wait() did not return after Done()")
	}
}

func TestWaitGroupWaitsForSeveralWorkers(t *testing.T) {
	t.Parallel()

	wg := NewWaitGroup()
	const workers = 10
	wg.Add(workers)

	var completed atomic.Int64
	for range workers {
		go func() {
			defer wg.Done()
			completed.Add(1)
		}()
	}

	wg.Wait()

	if got := completed.Load(); got != workers {
		t.Fatalf("completed workers = %d, want %d", got, workers)
	}
}

func TestWaitGroupReleasesMultipleWaiters(t *testing.T) {
	t.Parallel()

	wg := NewWaitGroup()
	wg.Add(1)

	const waiters = 5
	released := make(chan struct{}, waiters)
	for range waiters {
		go func() {
			wg.Wait()
			released <- struct{}{}
		}()
	}

	select {
	case <-released:
		t.Fatal("waiter was released before Done()")
	case <-time.After(20 * time.Millisecond):
	}

	wg.Done()

	for range waiters {
		select {
		case <-released:
		case <-time.After(100 * time.Millisecond):
			t.Fatal("not all waiters were released")
		}
	}
}

func TestWaitGroupDonePanicsWhenCounterBecomesNegative(t *testing.T) {
	t.Parallel()

	wg := NewWaitGroup()

	defer func() {
		if recover() == nil {
			t.Fatal("Done() did not panic when counter became negative")
		}
	}()

	wg.Done()
}
