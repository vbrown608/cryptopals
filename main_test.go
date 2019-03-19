package main

import (
	"encoding/hex"
	"testing"
)

func TestRepeatingKeyXOR(t *testing.T) {
	plainText := []byte(`Burning 'em, if you ain't quick and nimble
I go crazy when I hear a cymbal`)
	expected := "0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f"
	encrypted := RepeatingKeyXOR(plainText, []byte("ICE"))
	cipherText := hex.EncodeToString(encrypted)
	if cipherText != expected {
		t.Errorf("Expected:\t%s\nGot:\t%s\n", expected, encrypted)
	}
}
