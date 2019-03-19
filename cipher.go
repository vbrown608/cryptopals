package main

func RepeatingKeyXOR(in []byte, key []byte) []byte {
	r := make([]byte, len(in))
	for i := 0; i < len(in); i++ {
		r[i] = in[i] ^ key[i%len(key)]
	}
	return r
}
