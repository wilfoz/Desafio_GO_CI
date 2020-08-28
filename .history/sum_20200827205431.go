package main

import (
	"fmt"
)

func main() {

	res := sum(1, 2)
	fmt.Println("1+2 =", res)
}

func sum(a int, b int) int {
	return a + b
}
