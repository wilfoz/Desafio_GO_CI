package main
import (
  "testing"
)
func TestSomaGo(t *testing.T){
	got := soma(2,3)
	want:= 5.000000

	if got != want{
		t.Errorf("somou errado")
	}

}
