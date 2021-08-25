package main

import "fmt"

func main() {
	mess := make(chan string)

	go func() {
		mess <- "ping"
	}()

	msg := <-mess
	fmt.Println(msg)
}
