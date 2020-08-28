package main

import (
	"testing"
)

func TestSumGo(t *testing.T) {
	got := sum(2, 3)
	want := 5

	if got != want {
		t.Errorf("Error!!")
	}

}
