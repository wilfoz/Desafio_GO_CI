package main
import (
  "fmt"
  "strconv"
  "os"
)
func main() {
	valor1,_ := strconv.ParseFloat(os.Args[1],64);
	valor2,_ := strconv.ParseFloat(os.Args[2],64);
	fmt.Println(fmt.Sprintf("%f", soma(valor1,valor2)))
}
func soma(valor1 float64,valor2 float64) float64{
	//value := os.Args[1]
	return valor1 + valor2;
}
