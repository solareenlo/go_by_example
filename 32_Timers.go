package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(2 * time.Second)

	<-timer1.C
	fmt.Println("Timer 1 fired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()

	timer3 := time.NewTimer(time.Second)
	go func() {
		<-timer3.C
		fmt.Println("Timer 3 fired")
	}()
	stop3 := timer3.Stop()
	if stop3 {
		fmt.Println("Timer 3 Stopped")
	}

	time.Sleep(2 * time.Second)
}
