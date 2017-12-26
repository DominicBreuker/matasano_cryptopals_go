package p0x08

import (
	"fmt"
)

// Detector can detect AES in ECB mode by counting duplicate blocks
type Detector struct {
	Best []byte

	bs    int
	score float64
}

func New(bs int) *Detector {
	return &Detector{
		bs: bs,
	}
}

func (d *Detector) Update(b []byte) bool {
	s := score(b, d.bs)
	if s > d.score {
		d.Best = b
		d.score = s

		return true
	}

	return false
}

// score counts the number of duplicate blocks in the buffer, normalized by buffer length
func score(c []byte, bs int) float64 {
	blockSet := make(map[string]bool)
	dups := 0
	for i := 0; i < len(c); i += bs {
		block := fmt.Sprintf("%b", c[i:i+bs])
		if _, ok := blockSet[block]; ok {
			dups++
		}
		blockSet[block] = true
	}
	return float64(dups) / float64(len(c)) * float64(bs)
}
