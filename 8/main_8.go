package main

import (
	"fmt"
	"time"
)

type WaitGroup struct {
	semaphore chan struct{}
	counter   int
	waiters   []chan struct{}
}

func NewWaitGroup() *WaitGroup {
	wg := &WaitGroup{
		semaphore: make(chan struct{}, 1),
	}
	wg.semaphore <- struct{}{}
	return wg
}

func (wg *WaitGroup) Add(delta int) {
	wg.lock()
	defer wg.unlock()

	next := wg.counter + delta
	if next < 0 {
		panic("negative WaitGroup counter")
	}

	wg.counter = next
	if wg.counter != 0 {
		return
	}

	for _, waiter := range wg.waiters {
		close(waiter)
	}
	wg.waiters = nil
}

func (wg *WaitGroup) Done() {
	wg.Add(-1)
}

func (wg *WaitGroup) Wait() {
	wg.lock()
	if wg.counter == 0 {
		wg.unlock()
		return
	}

	waiter := make(chan struct{})
	wg.waiters = append(wg.waiters, waiter)
	wg.unlock()

	<-waiter
}

func (wg *WaitGroup) lock() {
	<-wg.semaphore
}

func (wg *WaitGroup) unlock() {
	wg.semaphore <- struct{}{}
}

func main() {
	wg := NewWaitGroup()

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			time.Sleep(time.Duration(workerID) * 100 * time.Millisecond)
			fmt.Printf("worker %d done\n", workerID)
		}(i)
	}

	wg.Wait()
	fmt.Println("all workers done")
}
