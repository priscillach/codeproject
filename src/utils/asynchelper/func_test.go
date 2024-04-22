package asynchelper

import (
	"fmt"
	"leetcode/src/utils/timehelper"
	"math/rand"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestSyncWaitGroup(t *testing.T) {
	g := sync.WaitGroup{}
	g.Add(10)
	for i := 0; i < 10; i++ {
		goroutineNum := i
		go func() {
			defer g.Done()
			GlobalLock.Lock()
			defer GlobalLock.Unlock()
			startTime := time.Now()
			fmt.Printf("gorotine: %v, time: %v\n start", goroutineNum, startTime.Second())
			time.Sleep(time.Duration(2+rand.Intn(5)) * time.Second)
			endTime := time.Now()
			fmt.Printf("gorotine: %v, time: %v end, cost: %v\n", goroutineNum, endTime.Second(), timehelper.FormatDuration(endTime.Sub(startTime)))
		}()
	}
	g.Wait()
}

func TestAtomicCounter(t *testing.T) {
	var num int32
	g := sync.WaitGroup{}
	g.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer g.Done()
			atomic.AddInt32(&num, 1)
		}()
	}
	g.Wait()
	fmt.Println(num)

	num = 0
	for i := 0; i < 100; i++ {
		go func() {
			num++
		}()
	}
	fmt.Println(num)
}

func TestSyncLock(t *testing.T) {
	syncLock := NewSyncLockByeKey()
	syncLock.AcquireSyncLock("key1")
	fmt.Println("Lock acquired for key1")

	go func() {
		syncLock.AcquireSyncLock("key1")
		fmt.Println("goroutine Lock acquired for key1")
		// Do some work...
		syncLock.ReleaseSyncLock("key1")
		fmt.Println("goroutine Lock released for key1")
	}()

	time.Sleep(10 * time.Second)
	fmt.Println("Lock released for key1")
	syncLock.ReleaseSyncLock("key1")
	// Wait for goroutines to complete
	fmt.Scanln()
}

func TestSyncUnlock(t *testing.T) {
	syncLock := NewSyncLockByeKey()
	syncLock.AcquireSyncLock("key1")
	syncLock.ReleaseSyncLock("key1")
	syncLock.ReleaseSyncLock("key1")
}

func TestSyncUnlock2(t *testing.T) {
	syncLock := NewSyncLockByeKey()
	syncLock.AcquireSyncLock("key1")
	syncLock.ReleaseSyncLock("key2")
}

func TestSyncLock2(t *testing.T) {
	syncLock := NewSyncLockByeKey()
	go func() {
		syncLock.AcquireSyncLock("key1")
		fmt.Println("Lock acquired for key1-1")
		time.Sleep(time.Duration(2+rand.Intn(5)) * time.Second)
		syncLock.ReleaseSyncLock("key1")
		fmt.Println("Lock released for key1-1")
	}()
	go func() {
		syncLock.AcquireSyncLock("key2")
		fmt.Println("Lock acquired for key2-1")
		time.Sleep(time.Duration(2+rand.Intn(5)) * time.Second)
		syncLock.ReleaseSyncLock("key2")
		fmt.Println("Lock released for key2-1")
	}()

	time.Sleep(100 * time.Millisecond)
	syncLock.AcquireSyncLock("key1")
	fmt.Println("Lock acquired for key1-2")
	syncLock.AcquireSyncLock("key2")
	fmt.Println("Lock acquired for key2-2")
	time.Sleep(time.Duration(2+rand.Intn(5)) * time.Second)

	syncLock.ReleaseSyncLock("key2")
	fmt.Println("Lock released for key2-2")

	syncLock.ReleaseSyncLock("key1")
	fmt.Println("Lock released for key1-2")

	syncLock.AcquireSyncLock("key1")
	fmt.Println("Lock acquired for key1-3")
	time.Sleep(time.Duration(2+rand.Intn(5)) * time.Second)

	syncLock.ReleaseSyncLock("key1")
	fmt.Println("Lock released for key1-3")
}
