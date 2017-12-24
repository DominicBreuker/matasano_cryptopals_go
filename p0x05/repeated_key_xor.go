package p0x05

import (
	"errors"
	"fmt"
)

// Encrypt returns the ciphertext for plaintext p under the repeated XOR key k
func Encrypt(p, k []byte) ([]byte, error) {
	var c []byte
	if len(k) == 0 {
		return c, errors.New("Key must have at least one byte")
	}

	c = make([]byte, len(p))
	for i := 0; i < len(p); i++ {
		c[i] = p[i] ^ k[i%len(k)]
	}
	return c, nil
}

func MustEncrypt(p, k []byte) []byte {
	c, err := Encrypt(p, k)
	if err != nil {
		panic(fmt.Sprintf("Unexpected error: %v", err))
	}
	return c
}
