package main

import (
  "fmt"
  "strconv"
  "os"
)

func main() {

	valueA,_ := strconv.ParseFloat(os.Args[1],64);
	valueB,_ := strconv.ParseFloat(os.Args[2],64);

	fmt.Println(fmt.Sprintf("%f", sum(valueA,valueB)))
}

func sum(valueA float64,valueB float64) float64{
	return valueA + valueB;
}
