package main

import (
	"crypto/rand"
	"fmt"
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
	padS := size - (len(in) % size)
	target += padS
	r := make([]byte, target)
	for i := 0; i < target; i++ {
		if i < len(in) {
			r[i] = in[i]
		} else {
			r[i] = byte(padS)
		}
	}
	return r
}

func unpad(in []byte) ([]byte, error) {
	padS := int(in[len(in)-1])
	for i := 0; i < padS; i++ {
		if in[len(in)-padS+i] != byte(padS) {
			return []byte{}, fmt.Errorf("Invalid padding")
		}
	}
	return in[0 : len(in)-padS], nil
}
