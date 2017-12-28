package p0x07

import (
	"crypto/aes"
	"fmt"

	"github.com/dominicbreuker/matasano_cryptopals_go/p0x09"
)

// Decrypt decrytps the ciphertext c with key k
func Decrypt(ct, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("Could not create AES cipher from key: %v", err)
	}

	bs := block.BlockSize()

	pt := make([]byte, len(ct))
	for i := 0; i < len(ct); i += bs {
		block.Decrypt(pt[i:i+bs], ct[i:i+bs])
	}
	pt = p0x09.StripPad(pt)

	return pt, nil
}

// Encrypts the plaintext p with key k
func Encrypt(pt, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("Could not create AES cipher from key: %v", err)
	}

	bs := block.BlockSize()

	pt = p0x09.Pad(pt, bs)
	ct := make([]byte, len(pt))
	for i := 0; i < len(pt); i += bs {
		block.Encrypt(ct[i:i+bs], pt[i:i+bs])
	}

	return ct, nil
}
