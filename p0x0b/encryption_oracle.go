package p0x0b

import (
	"crypto/rand"
	"fmt"
)

func randomKey(n int) ([]byte, error) {
	key := make([]byte, n)
	_, err := rand.Read(key)
	if err != nil {
		return nil, fmt.Errorf("Error generating random key: %v", err)
	}
	return key, nil
}
