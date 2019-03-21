package main

import "testing"

func TestHammingDistance(t *testing.T) {
	a := []byte("this is a test")
	b := []byte("wokka wokka!!!")
	if r := HammingDistance(a, b); r != 37 {
		t.Errorf("Expected Hamming distance to be %d, got %d", 37, r)
	}
}
