package main

import (
	"crypto/rand"
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

func detectOracleCipher(pt, ct []byte) int {
	return countRepeats(ct)
}
