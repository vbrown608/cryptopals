package main

import (
	"math"
	"math/bits"
)

// HammingDistance computes the number of differing bits.
func HammingDistance(a []byte, b []byte) (r int) {
	if len(a) != len(b) {
		panic("Length of inputs should be equal")
	}
	for i := 0; i < len(a); i++ {
		r += int(bits.OnesCount8(a[i] ^ b[i]))
	}
	return
}

func BestKeySize(in []byte) int {
	min := math.Inf(1)
	var bestKeySize int
	for keySize := 4; keySize <= 40; keySize++ {
		b1 := in[:keySize]
		b2 := in[keySize : 2*keySize]
		b3 := in[2*keySize : 3*keySize]
		b4 := in[3*keySize : 4*keySize]
		dTotal := HammingDistance(b1, b2) + HammingDistance(b2, b3) + HammingDistance(b3, b4)
		d := float64(dTotal) / float64(keySize)
		if d < min {
			min = d
			bestKeySize = keySize
		}
	}
	return bestKeySize
}

func EnglishScore(text string) float64 {
	var ignored int
	count := make([]int, len(letterFrequencies))
	for i := 0; i < len(text); i++ {
		c := text[i]
		if c >= 65 && c <= 90 {
			// uppercase A-Z
			count[c-65]++
		} else if c >= 97 && c <= 122 {
			// lowercase a-z
			count[c-97]++
		} else if c == 32 {
			// space
			count[26]++
		} else if c >= 32 && c <= 126 {
			// numbers and punctuation
			count[27]++
		} else if c == 9 || c == 10 || c == 13 {
			// TAB, CR, LF
			ignored++
		} else {
			// non-printable character
			return math.Inf(1)
		}
	}

	var chi2 float64
	length := len(text) - ignored
	for i := 0; i < len(letterFrequencies); i++ {
		observed := count[i]
		expected := float64(length) * letterFrequencies[i]
		diff := float64(observed) - expected
		chi2 += diff * diff / expected
	}
	return chi2
}

var letterFrequencies = []float64{
	0.0651738,
	0.0124248,
	0.0217339,
	0.0349835,
	0.1041442,
	0.0197881,
	0.0158610,
	0.0492888,
	0.0558094,
	0.0009033,
	0.0050529,
	0.0331490,
	0.0202124,
	0.0564513,
	0.0596302,
	0.0137645,
	0.0008606,
	0.0497563,
	0.0515760,
	0.0729357,
	0.0225134,
	0.0082903,
	0.0171272,
	0.0013692,
	0.0145984,
	0.0007836,
	0.1918182, // space
	0.0008,    // other printable ascii
}
