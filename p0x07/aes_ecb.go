package p0x07

import (
	"crypto/aes"
	"fmt"
)

// Decrypt decrytps the ciphertext c inplace with key k
func Decrypt(c, k []byte) error {
	block, err := aes.NewCipher(k)
	if err != nil {
		return fmt.Errorf("Could not create AES cipher from key: %v", err)
	}

	bs := block.BlockSize()

	for i := 0; i < len(c); i += bs {
		block.Decrypt(c[i:i+bs], c[i:i+bs])
	}
	fmt.Println(block)

	return nil
}
