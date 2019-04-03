package main

import (
	"crypto/rand"
	"math/big"
)

func main() {
}

func encryptionOracle(pt []byte) []byte {
	// Generate a random key.
	key := make([]byte, 16)
	rand.Read(key)
	frontPad := make([]byte, randInt(6)+5)

	// Append 5-10 random bytes to front and back.
	rand.Read(frontPad)
	pt = append(frontPad, pt...)
	backPad := make([]byte, randInt(6)+5)
	rand.Read(backPad)
	pt = append(pt, backPad...)

	// Choose a random cipher.
	cipher := randInt(2)
	if cipher == 0 {
		return aes128ECB(pt, key)
	} else {
		iv := make([]byte, 16)
		rand.Read(iv)
		return encryptCBC(pt, iv, key)
	}
}

func randInt(max int) int {
	i, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		panic(err)
	}
	return int(i.Int64())
}

func detectOracleCipher(pt, ct []byte) int {
	return countRepeats(ct)
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
