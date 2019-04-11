package main

import "testing"

// Challenge 9
func TestPad(t *testing.T) {
	padded := pad([]byte("YELLOW SUBMARINE"), 2)
	if string(padded) != "YELLOW SUBMARINE\x02\x02" {
		t.Error(padded)
	}
	padded = pad([]byte("YELLOW SUBMARINE"), 20)
	if string(padded) != "YELLOW SUBMARINE\x04\x04\x04\x04" {
		t.Error(padded)
	}
}

// Challenge 15
func TestUnpad(t *testing.T) {
	tests := []struct {
		in []byte
		r  []byte
		e  bool
	}{
		{[]byte("ICE ICE BABY\x04\x04\x04\x04"), []byte("ICE ICE BABY"), false},
		{[]byte("ICE ICE BABY\x05\x05\x05\x05"), []byte{}, true},
		{[]byte("ICE ICE BABY\x01\x02\x03\x04"), []byte{}, true},
	}
	for _, tt := range tests {
		gotR, gotE := unpad(tt.in)
		if !bytesEq(tt.r, gotR) {
			t.Errorf("Unpadded output for %s was %s", tt.in, gotR)
		}
		if tt.e == (gotE == nil) {
			t.Errorf("Error for %s was %v", tt.in, gotE)
		}
	}
}
