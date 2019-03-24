package main

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"io/ioutil"
	"os"
	"testing"
)

func TestAES128ECB(t *testing.T) {
	in, err := ioutil.ReadFile("files/7.txt")
	if err != nil {
		t.Fatal(err)
	}
	ct, err := base64.StdEncoding.DecodeString(string(in))
	if err != nil {
		t.Fatal(err)
	}
	key := []byte("YELLOW SUBMARINE")
	pt := aes128ECB(ct, key)
	if !bytes.Contains(pt, []byte("Supercalafragilisticexpialidocious")) {
		t.Fatal(string(pt))
	}
}

func TestDecryptECB(t *testing.T) {
	file, _ := os.Open("files/8.txt")
	defer file.Close()

	var max int
	var detected []byte
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ct := HexToBytes(scanner.Text())
		count := countRepeats(ct)
		if count > max {
			max = count
			detected = ct
		}
	}
	if hex.EncodeToString(detected) != "d880619740a8a19b7840a8a31c810a3d08649af70dc06f4fd5d2d69c744cd283e2dd052f6b641dbf9d11b0348542bb5708649af70dc06f4fd5d2d69c744cd2839475c9dfdbc1d46597949d9c7e82bf5a08649af70dc06f4fd5d2d69c744cd28397a93eab8d6aecd566489154789a6b0308649af70dc06f4fd5d2d69c744cd283d403180c98c8f6db1f2a3f9c4040deb0ab51b29933f2c123c58386b06fba186a" {
		t.Fatal()
	}
}
