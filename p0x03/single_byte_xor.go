package p0x03

import (
	"github.com/dominicbreuker/matasano_cryptopals_go/p0x02"
)

func score(b []byte) float64 {
	var result float64
	var i int
	s := string(b)
	for _, r := range s {
		result += scores[r]
		i++
	}
	return result / float64(i)
}

func decryptWithKey(b []byte, key byte) []byte {
	n := len(b)
	k := make([]byte, n, n)
	for i := 0; i < n; i++ {
		k[i] = key
	}

	decrypted, err := p0x02.XOR(b, k)
	if err != nil {
		panic("XOR cannot fail since both buffers must have same length!")
	}
	return decrypted
}

func Decrypt(b []byte) ([]byte, float64, byte) {
	best, plain, key := 0.0, []byte{}, byte(0)
	for i := uint8(0); i < 255; i++ {
		p := decryptWithKey(b, byte(i))
		score := score(p)
		if score > best {
			best, plain, key = score, p, byte(i)
		}
	}
	return plain, best, key
}
