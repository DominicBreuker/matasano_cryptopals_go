package utils

import (
	"encoding/hex"
	"fmt"
)

func MustDecodeHexString(s string) []byte {
	b, err := hex.DecodeString(s)
	if err != nil {
		panic(fmt.Sprintf("Could not decode hex string: %v", err))
	}
	return b
}
