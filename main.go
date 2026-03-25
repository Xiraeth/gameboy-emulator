package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	z := add(4, 5)

	fmt.Println(z)
}
