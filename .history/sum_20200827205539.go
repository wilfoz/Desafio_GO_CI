package main

import (
	"fmt"
)

func main() {

	res := sum(2, 3)
	fmt.Println("2 + 3 =", res)
}

func sum(a int, b int) int {
	return a + b
}
