package main

import "fmt"

func ya(a []int) []int {
	a = append(a, 10, 10, 10)
	fmt.Println(a)
	return a
}

func main() {
	a := make([]int, 0, 2)
	fmt.Println(a)
    ya(a)
	fmt.Println(a)
    fmt.Println(len(a), cap(a))
}
