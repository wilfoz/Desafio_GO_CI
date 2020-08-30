package main

import (
	"fmt"
)

func main() {

	res := sum(5, 5)
	fmt.Println("5 + 5 =", res)
}

func sum(a int, b int) int {
	return a + b
}
