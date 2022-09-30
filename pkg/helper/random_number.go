package helper

import "math/rand"

const numbBytes = "1234567890"

func RandNumbBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = numbBytes[rand.Intn(len(numbBytes))]
	}
	return string(b)
}
