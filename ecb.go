package main

import (
	"crypto/aes"
	"crypto/rand"
	"fmt"
	"net/url"
)

// Electronic Cookbook (ECB) mode encryption

func encryptECB(ct, key []byte) []byte {
	cipher, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	ct = pad(ct, 16)
	for i := 0; i < len(ct); i += 16 {
		cipher.Encrypt(ct[i:i+16], ct[i:i+16])
	}
	return ct
}

func decryptECB(ct, key []byte) []byte {
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

var fixedECBKey = make([]byte, 16)

func fixedECB(ct []byte) []byte {
	return encryptECB(ct, fixedECBKey)
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

func BreakECB(secret []byte) []byte {
	rand.Read(fixedECBKey)
	var result []byte
	for i := 0; i < len(secret); i++ {
		b := getECBByte(result, secret)
		result = append(result, b)
	}
	return result
}

func getECBByte(known []byte, ct []byte) byte {
	targetLen := roundUpToMultiple(len(known)+1, 16)
	padded := append(make([]byte, targetLen-len(known)-1), known...)
	unknown := ct[len(known)]
	want := fixedECB(append(padded, unknown))
	for i := 0; i < 256; i++ {
		got := fixedECB(append(padded, byte(i)))
		if bytesEq(got, want) {
			return byte(i)
		}
	}
	panic("No match found")
	return byte(0)
}
