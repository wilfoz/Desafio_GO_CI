package main

import (
	"testing"
)

func TestSomaGo(t *testing.T) {
	got := sum(2, 3)
	want := 5

	if got != want {
		t.Errorf("Error!!")
	}

}
