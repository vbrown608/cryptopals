package main

import (
	"io/ioutil"
	"testing"
)

func TestDetectOracleCipher(t *testing.T) {
	pt, err := ioutil.ReadFile("files/11.txt")
	if err != nil {
		t.Fatal(err)
	}
	ct := encryptionOracle(pt)
	detectOracleCipher(pt, ct)
}
