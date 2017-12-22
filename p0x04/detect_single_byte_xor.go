package p0x04

import "github.com/dominicbreuker/matasano_cryptopals_go/p0x03"

type Detector struct {
	Best []byte

	score float64
}

func New() *Detector {
	return &Detector{}
}

func (d *Detector) Update(b []byte) bool {
	_, s, _ := p0x03.Decrypt(b)
	if s > d.score {
		d.Best = b
		d.score = s

		return true
	}

	return false
}
