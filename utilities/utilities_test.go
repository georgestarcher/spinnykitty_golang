package utilities

import (
	"log"
	"testing"
)

// Test Empty Mean slice input
func TestEmptyMean(t *testing.T) {

	var testSlice []int

	want := 0.0
	got := CalcMean(testSlice)

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
	log.Println(got)
}

// Test Valid Mean slice input
func TestMean(t *testing.T) {

	testSlice := make([]int, 5)
	testSlice[0] = 4
	testSlice[1] = 4
	testSlice[2] = 4
	testSlice[3] = 4
	testSlice[4] = 4

	want := 4.0
	got := CalcMean(testSlice)

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}
	log.Println(got)
}
