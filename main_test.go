package main

import (
	"io/ioutil"
	"testing"
)

// Challenge 9
func TestPad(t *testing.T) {
	padded := pad([]byte("YELLOW SUBMARINE"), 20)
	if string(padded) != "YELLOW SUBMARINE\x04\x04\x04\x04" {
		t.Error(padded)
	}
}

func TestDetectOracleCipher(t *testing.T) {
	pt, err := ioutil.ReadFile("files/11.txt")
	if err != nil {
		t.Fatal(err)
	}
	ct := encryptionOracle(pt)
	detectOracleCipher(pt, ct)
}
