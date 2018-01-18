package p0x0c

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"

	"github.com/dominicbreuker/matasano_cryptopals_go/p0x07"
)

const maxSize = 256

type Oracle func([]byte) []byte

func NewOracle(r *rand.Rand, postfix []byte) Oracle {
	key := randomBytes(r, 16)

	return func(in []byte) []byte {
		pt := append(in, postfix...)

		return mustEncryptECB(pt, key)
	}
}

//func Break(ct []byte, oracle Oracle) ([]byte, error) {
//	bs, _ := detectBlockSize(oracle)
//	//isECB := p0x0b.DetectECB
//}

// detectBlockSize detects the block size of the oracle's encryption function by increasing the input string one byte at a time.
// Once we see the output length increasing, we know the difference between this and the previous length is the block size.
func detectBlockSize(oracle Oracle) (int, error) {
	var tmp int
	for i := 1; i <= maxSize; i++ {
		s := strings.Repeat("A", i)
		l := len(oracle([]byte(s)))

		if tmp == 0 {
			tmp = l
		}
		if l > tmp {
			return l - tmp, nil
		}

	}
	return -1, fmt.Errorf("Could not identify block size. Is it larger than %d", maxSize)
}

func decryptSingleByte(ct []byte, bs int, oracle Oracle) (byte, error) {
	if bs < 2 {
		return byte('\x00'), fmt.Errorf("Block size must be at least 2")
	}
	prefix := bytes.Repeat([]byte("\x00"), bs-1)
	tmp := oracle(append(prefix, ct...))
	var i byte
	for i = 0; i < 256; i++ {
		tmp2 := oracle(append(prefix, i, ct...))
		fmt.Printf("%+v", tmp2)
	}
	return tmp[0], nil
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
