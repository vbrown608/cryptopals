package main

import "testing"

// Challenge 9
func testPad(t *testing.T) {
	padded := pad([]byte("YELLOW SUBMARINE"), 20)
	if string(padded) != "YELLOW SUBMARINE\x04\x04\x04\x04" {
		t.Error(padded)
	}
}
