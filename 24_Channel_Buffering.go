package main

import "fmt"

func main() {
	mess := make(chan string, 2)

	mess <- "hello"
	mess <- "world"

	fmt.Println(<-mess)
	fmt.Println(<-mess)
}
