package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var state = make(map[int]int)
	var mutex = &sync.Mutex{}
	var readOps, writeOps uint64

	for r := 0; r < 10; r++ {
		go func() {
			sum := 0
			for {
				k := rand.Intn(5)
				mutex.Lock()
				sum += state[k]
				mutex.Unlock()
				fmt.Print(sum, " ")
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	for w := 0; w < 10; w++ {
		go func() {
			k := rand.Intn(5)
			v := rand.Intn(100)
			mutex.Lock()
			state[k] = v
			mutex.Unlock()
			atomic.AddUint64(&writeOps, 1)
			time.Sleep(time.Millisecond)
		}()
	}

	time.Sleep(time.Second)

	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)

	mutex.Lock()
	fmt.Println("state:", state)
	mutex.Unlock()
}