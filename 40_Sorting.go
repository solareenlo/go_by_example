package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"c", "b", "a"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	ints := []int{10, 2, 3}
	sort.Ints(ints)
	fmt.Println("Ints:   ", ints)

	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)
}
