package main

import (
	"math"
)

func xor(in, key, r []byte) []byte {
	if len(in) != len(r) {
		panic("Mismatched lengths in XOR")
	}
	for i := 0; i < len(in); i++ {
		r[i] = in[i] ^ key[i%len(key)]
	}
	return r
}

func xorWithReturn(in, key []byte) []byte {
	r := make([]byte, len(in))
	return xor(in, key, r)
}

func BestCharXOR(in []byte) ([]byte, byte, float64) {
	var bestKey byte
	var bestOutput []byte
	min := math.Inf(1)
	for i := 0; i < 128; i++ {
		xored := xorWithReturn([]byte(in), []byte{byte(i)})
		score := EnglishScore(string(xored))
		if score < min {
			min = score
			bestKey = byte(i)
			bestOutput = xored
		}
	}
	return bestOutput, bestKey, min
}

func SolveBlocks(in []byte, keySize int) []byte {
	key := make([]byte, keySize)
	blockSize := len(in) / keySize
	for i := 0; i < keySize; i++ {
		block := make([]byte, blockSize)
		// Build block.
		for j := 0; j < blockSize; j++ {
			block[j] = in[j*keySize+i]
		}
		_, bestChar, _ := BestCharXOR(block)
		key[i] = bestChar
	}
	return key
}
