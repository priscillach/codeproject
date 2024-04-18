package utils

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestInt2String(t *testing.T) {
	fmt.Println(Int2String(0))
	fmt.Println(Int2String(123))
}

func TestSync(t *testing.T) {
	g := sync.WaitGroup{}
	g.Add(10)
	for i := 0; i < 10; i++ {
		goroutineNum := i
		go func() {
			defer g.Done()
			lock.Lock()
			defer lock.Unlock()
			startTime := time.Now()
			fmt.Printf("gorotine: %v, time: %v\n start", goroutineNum, startTime.Second())
			time.Sleep(time.Duration(2+rand.Intn(5)) * time.Second)
			endTime := time.Now()
			fmt.Printf("gorotine: %v, time: %v end, cost: %v\n", goroutineNum, endTime.Second(), FormatDuration(endTime.Sub(startTime)))
		}()
	}
	g.Wait()
}
