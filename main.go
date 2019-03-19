package main

func main() {
}

func BestCharXOR(in []byte) (best []byte, max float64) {
	for i := 0; i < 128; i++ {
		xored := RepeatingKeyXOR([]byte(in), []byte{byte(i)})
		score := EnglishScore(string(xored))
		if score > max {
			max = score
			best = xored
		}
	}
	return
}
