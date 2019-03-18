package main

import "encoding/hex"

func FixedXOR(s1 string, s2 string) string {
	// Assume len(s1) == len(s2)
	b1 := HexToBytes(s1)
	b2 := HexToBytes(s2)
	r := make([]byte, len(s1)/2)
	for i := 0; i < len(b1); i++ {
		r[i] = b1[i] ^ b2[i]
	}
	return hex.EncodeToString(r)
}
