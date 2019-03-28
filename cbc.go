package main

import (
	"crypto/aes"
)

func encryptCBC(pt, iv, key []byte) []byte {
	blockSize := len(key)
	pt = pad(pt, blockSize)
	c, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	ct := make([]byte, len(pt))
	last := iv
	for i := 0; i < len(ct); i += blockSize {
		xor(pt[i:i+blockSize], last, ct[i:i+blockSize])
		c.Encrypt(ct[i:i+blockSize], ct[i:i+blockSize])
		last = ct[i : i+blockSize]
	}
	return ct
}

func decryptCBC(ct, iv, key []byte) []byte {
	blockSize := len(key)
	c, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	pt := make([]byte, len(ct))
	last := iv
	for i := 0; i < len(pt); i += blockSize {
		c.Decrypt(pt[i:i+blockSize], ct[i:i+blockSize])
		xor(pt[i:i+blockSize], last, pt[i:i+blockSize])
		last = ct[i : i+blockSize]
	}
	return pt
}
