package p0x0a

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"

	"github.com/dominicbreuker/matasano_cryptopals_go/p0x09"
)

// Decrypt decrypts AES CBC ciphertext ct given a key and iv
func Decrypt(ct, key, iv []byte) ([]byte, error) {
	block, err := getBlockCipher(key, iv)
	if err != nil {
		return nil, fmt.Errorf("Can't decrypt: %v", err)
	}
	bs := block.BlockSize()
	if len(ct)%bs > 0 {
		return nil, fmt.Errorf("Can't decrypt ciphertext: lenght %d not a multiple of block size %d", len(ct), bs)
	}

	pt := decryptCBC(ct, iv, block)
	return p0x09.StripPad(pt), nil
}

// Encrypt decrypts AES CBC ciphertext ct given a key and iv
func Encrypt(pt, key, iv []byte) ([]byte, error) {
	block, err := getBlockCipher(key, iv)
	if err != nil {
		return nil, fmt.Errorf("Can't encrypt: %v", err)
	}

	padded := p0x09.Pad(pt, block.BlockSize())
	return encryptCBC(padded, iv, block), nil
}

// getBlockCipher creates a block cipher and verifies it can be used with the given ciphertext ct and IV iv
func getBlockCipher(key, iv []byte) (cipher.Block, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("Could not create AES cipher from key: %v", err)
	}
	bs := block.BlockSize()

	if len(iv) != bs {
		return nil, fmt.Errorf("IV does not equal block size: %d", len(iv))
	}

	return block, nil
}

func decryptCBC(ct, iv []byte, block cipher.Block) []byte {
	bs := block.BlockSize()

	prevBlock := iv
	pt := make([]byte, len(ct))
	for i := 0; i < len(ct); i += bs {
		block.Decrypt(pt[i:i+bs], ct[i:i+bs])
		for j := 0; j < bs; j++ {
			pt[i+j] ^= prevBlock[j]
		}
		prevBlock = ct[i : i+bs]
	}
	return pt
}

func encryptCBC(pt, iv []byte, block cipher.Block) []byte {
	bs := block.BlockSize()

	prevBlock := iv
	tmp := make([]byte, bs)
	ct := make([]byte, len(pt))
	for i := 0; i < len(pt); i += bs {
		for j := 0; j < bs; j++ {
			tmp[j] = pt[i+j] ^ prevBlock[j]
		}
		block.Encrypt(ct[i:i+bs], tmp)
		prevBlock = ct[i : i+bs]
	}
	return ct
}
