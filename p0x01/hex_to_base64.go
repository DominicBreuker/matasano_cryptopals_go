package p0x01

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func HexToBase64(s string) (string, error) {
	decoded, err := hex.DecodeString(s)
	if err != nil {
		return "", fmt.Errorf("Can't decode hex string: %v", err)
	}

	return base64.StdEncoding.EncodeToString(decoded), nil
}
