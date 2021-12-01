package main

import "testing"

func TestSlow(t *testing.T) {
	sum, _ := slow(100)
	if sum != 5050 {
		t.Errorf("Expected sum of all numbers between 1 and 100 to be 5050, but got %d", sum)
	}
}

func TestFast(t *testing.T) {
	sum, _ := fast(100)
	if sum != 5050 {
		t.Errorf("Expected sum of all numbers between 1 and 100 to be 5050, but got %d", sum)
	}
}
