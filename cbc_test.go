package main

import (
	"encoding/base64"
	"io/ioutil"
	"strings"
	"testing"
)

func TestDecryptCBC(t *testing.T) {
	in, err := ioutil.ReadFile("files/10.txt")
	if err != nil {
		t.Fatal(err)
	}
	ct, err := base64.StdEncoding.DecodeString(string(in))
	if err != nil {
		t.Fatal(err)
	}
	key := []byte("YELLOW SUBMARINE")
	iv := make([]byte, 16)
	decrypted := decryptCBC(ct, iv, key)
	if !strings.Contains(string(decrypted), "Spaghetti with a spoon! Come on and say it!") {
		t.Error(string(decrypted))
	}
}

func TestEncryptCBC(t *testing.T) {
	pt := `I grow old ... I grow old ...
I shall wear the bottoms of my trousers rolled.`
	key := []byte("YELLOW SUBMARINE")
	iv := make([]byte, 16)
	ct := encryptCBC([]byte(pt), iv, key)
	newPT := string(decryptCBC(ct, iv, key))
	if !strings.Contains(newPT, pt) {
		t.Error(newPT)
	}
}
