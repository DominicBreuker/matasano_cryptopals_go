package p0x02

import (
	"errors"
)

func XOR(b1, b2 []byte) ([]byte, error) {
	if len(b1) != len(b2) {
		return nil, errors.New("Can't XOR buffers of different length")
	}
	n := len(b1)

	result := make([]byte, 0, n)
	for i := 0; i < n; i++ {
		result = append(result, b1[i]^b2[i])
	}
	return result, nil
}
