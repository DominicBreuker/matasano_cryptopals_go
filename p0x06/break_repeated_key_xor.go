package p0x06

import (
	"errors"
	"fmt"
	"math"
	"math/bits"

	"github.com/dominicbreuker/matasano_cryptopals_go/p0x03"
)

func Break(c []byte, min, max int) []byte {
	ks := guessKeysize(c, min, max)

	k := make([]byte, ks)
	blks := getBlocks(c, ks)
	for i, b := range blks {
		_, _, k[i] = p0x03.Decrypt(b)
	}

	return k
}

// getBlocks groups ciphertext into blocks, each containing all bytes encrpted with the same key byte
func getBlocks(c []byte, ks int) map[int][]byte {
	blocks := make(map[int][]byte, ks)

	for i, b := range c {
		blocks[i%ks] = append(blocks[i%ks], b)
	}

	return blocks
}

func guessKeysize(c []byte, min, max int) int {
	bestSize := min
	bestScore := math.MaxFloat64

	for k := min; k <= max; k++ {
		score := scoreKeysize(c, k)
		if score < bestScore {
			bestSize, bestScore = k, score
		}
	}

	return bestSize
}

// scoreKeysize ranks key sizes by computing edit distances between XORs of consecutive chunks of length keysize
// Idea is:
// - for correct keysize, XORing two chunks cancels out the key. When computing edit distance, the edit distance of the plaintext will be computed. Assuming the plaintext is not random, it should be small.
// - for wrong keysize, XORing two chunks does not cancel out the key. Hence all edit distances will contain randomness from the encryption. This drives up edit distances.
func scoreKeysize(c []byte, k int) float64 {
	var sum, n int
	prev := c[:k]

	for i := k; i < len(c); i += k {
		sum += mustGetEditDistance(prev, c[i:i+k])
		n += 1
	}

	return float64(sum) / float64(n) / float64(k)
}

func editDistance(b1, b2 []byte) (int, error) {
	if len(b1) != len(b2) {
		return 0, errors.New("Cannot compare buffers if different length")
	}
	n := len(b1)

	var d int
	for i := 0; i < n; i++ {
		d += bits.OnesCount(uint(b1[i] ^ b2[i]))
	}
	return d, nil
}

func mustGetEditDistance(b1, b2 []byte) int {
	d, err := editDistance(b1, b2)
	if err != nil {
		panic(fmt.Sprintf("Error computing edit distance: %v", err))
	}
	return d
}
