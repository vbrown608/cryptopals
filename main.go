package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Open("files/4.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		bytes := HexToBytes(scanner.Text())
		best, max := BestCharXOR(bytes)
		if max > 5 {
			fmt.Printf("%v\t%v\n", string(best), max)
		}
	}
}

func BestCharXOR(in []byte) (best []byte, max float64) {
	for i := 0; i < 128; i++ {
		xored := SingleCharXOR([]byte(in), byte(i))
		score := EnglishScore(string(xored))
		if score > max {
			max = score
			best = xored
		}
	}
	return
}

func SingleCharXOR(in []byte, c byte) []byte {
	r := make([]byte, len(in))
	for i := 0; i < len(in); i++ {
		r[i] = in[i] ^ c
	}
	return r
}
