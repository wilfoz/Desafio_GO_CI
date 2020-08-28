package main

import (
	"fmt"
)

func main() {

	res := sum(4, 5)
	fmt.Println("4 + 5 =", res)
}

func sum(a int, b int) int {
	return a + b
}
