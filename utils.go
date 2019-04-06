package main

import (
	"crypto/rand"
	"math/big"
)

// Round n up to the nearest multiple of m
func roundUpToMultiple(n, m int) int {
	if rem := n % m; rem > 0 {
		return n + m - rem
	}
	return n
}

func bytesEq(a, b []byte) bool {
	if len(a) != len(b) {
		panic("Block lengths should match")
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func randInt(max int) int {
	i, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		panic(err)
	}
	return int(i.Int64())
}

// Implements PKCS#7 padding
func pad(in []byte, size int) []byte {
	target := len(in)
	rem := len(in) % size
	if rem > 0 {
		target += size - rem
	}
	r := make([]byte, target)
	for i := 0; i < target; i++ {
		if i < len(in) {
			r[i] = in[i]
		} else {
			r[i] = '\x04'
		}
	}
	return r
}
