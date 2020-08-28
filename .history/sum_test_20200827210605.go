package main

import (
	"testing"
)

func TestSumGo(t *testing.T) {
	got := sum(5, 5)
	want := 10

	if got != want {
		t.Errorf("Error!!")
	}

}
