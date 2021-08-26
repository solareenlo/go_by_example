package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

type readOp struct {
	key int
	res chan int
}

type writeOp struct {
	key int
	val int
	res chan bool
}

func main() {
	var readOps, writeOps uint64
	reads := make(chan readOp)
	writes := make(chan writeOp)

	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.res <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.res <- true
			}
		}
	}()

	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := readOp{
					key: rand.Intn(5),
					res: make(chan int),
				}
				reads <- read
				<-read.res
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := writeOp{
					key: rand.Intn(5),
					val: rand.Intn(100),
					res: make(chan bool),
				}
				writes <- write
				<-write.res
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)

	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)
}
