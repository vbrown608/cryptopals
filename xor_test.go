package main

import (
	"bufio"
	"encoding/base64"
	"encoding/hex"
	"io/ioutil"
	"math"
	"os"
	"strings"
	"testing"
)

// Challenge 5
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

func TestBestCharXOR(t *testing.T) {
	file, _ := os.Open("files/4.txt")
	defer file.Close()

	min := math.Inf(1)
	var bestText []byte
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		bytes := HexToBytes(scanner.Text())
		text, _, score := BestCharXOR(bytes)
		if score < min {
			min = score
			bestText = text
		}
	}
	expected := "Now that the party is jumping\n"
	if string(bestText) != expected {
		t.Errorf("Expected %s, got %s", expected, bestText)
	}
}

// Challenge 6
func TestBreakRepeatingKeyXOR(t *testing.T) {
	in, err := ioutil.ReadFile("files/6.txt")
	if err != nil {
		t.Fatal(err)
	}
	inBytes, err := base64.StdEncoding.DecodeString(string(in))
	if err != nil {
		t.Fatal(err)
	}
	keySize := BestKeySize(inBytes)
	key := SolveBlocks(inBytes, keySize)
	plainText := string(RepeatingKeyXOR(inBytes, key))
	if !strings.Contains(plainText, "You're weakenin' fast, YO!") {
		t.Error(plainText)
	}
}
