package p0x0b

import (
	"fmt"
	"math/rand"
	"reflect"

	"github.com/dominicbreuker/matasano_cryptopals_go/p0x07"
	"github.com/dominicbreuker/matasano_cryptopals_go/p0x0a"
)

type Oracle func([]byte) ([]byte, bool)

func NewOracle(r *rand.Rand) Oracle {
	key := randomBytes(r, 16)
	useECB := r.Intn(2) == 0

	return func(in []byte) ([]byte, bool) {
		prefix := randomBytes(r, 10-r.Intn(6))
		postfix := randomBytes(r, 10-r.Intn(6))

		pt := prefix
		pt = append(pt, in...)
		pt = append(pt, postfix...)

		if useECB {
			return mustEncryptECB(pt, key), true
		} else {
			return mustEncryptCBC(pt, key, r), false
		}
	}
}

// DetectECB decides if the oracle uses ECB or CBC by making sure two entire blocks with the same plaintext are encrypted under the oracle.
// If ECB is used, both will have the same ciphertext. If CBC is used, the odds of that happening are practically zero.
func DetectECB(oracle Oracle) error {
	in := []byte("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
	out, usedECB := oracle(in)

	detectedECB := reflect.DeepEqual(out[16:32], out[32:48])
	if usedECB == detectedECB {
		return nil
	}

	return fmt.Errorf("Failed to detect: was %t but guessed %t", false, true)
}

func randomBytes(r *rand.Rand, n int) []byte {
	key := make([]byte, n)
	r.Read(key)
	return key
}

func mustEncryptECB(pt, key []byte) []byte {
	ct, err := p0x07.Encrypt(pt, key)
	if err != nil {
		panic(fmt.Sprintf("Unexpected error: %v", err))
	}
	return ct
}

func mustEncryptCBC(pt, key []byte, r *rand.Rand) []byte {
	iv := randomBytes(r, len(key))
	ct, err := p0x0a.Encrypt(pt, key, iv)
	if err != nil {
		panic(fmt.Sprintf("Unexpected error: %v", err))
	}
	return ct
}
