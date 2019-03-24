package main

func main() {
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
