package main

import (
	"crypto/aes"
)

// Electronic Cookbook (ECB) mode encryption

func aes128ECB(ct, key []byte) []byte {
	cipher, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	ct = pad(ct, 16)
	for i := 0; i < len(ct); i += 16 {
		cipher.Decrypt(ct[i:i+16], ct[i:i+16])
	}
	return ct
}

func countRepeats(in []byte) (count int) {
	seen := map[string]bool{}
	for i := 0; i < len(in); i += 16 {
		block := string(in[i : i+16])
		if seen[block] {
			count++
		}
		seen[block] = true
	}
	return count
}
