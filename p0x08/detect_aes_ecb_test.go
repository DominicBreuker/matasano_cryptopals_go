package p0x08

import (
	"fmt"
	"math"
	"reflect"
	"strings"
	"testing"

	"github.com/dominicbreuker/matasano_cryptopals_go/utils"
)

func TestDetect(t *testing.T) {
	scanner, file := utils.MustGetScanner("data.txt")

	detector := New(16)
	for scanner.Scan() {
		s := strings.TrimSuffix(scanner.Text(), "\n")
		b := utils.MustDecodeHexString(s)
		detector.Update(b)
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Strange error: %v", err)
	}
	file.Close()

	if !reflect.DeepEqual(detector.Best[:16], utils.MustDecodeHexString("d880619740a8a19b7840a8a31c810a3d")) {
		t.Fatalf("Wrong candidate discovered: %#x\n", detector.Best[:16])
	}
}

func TestScore(t *testing.T) {
	b := []byte("12341234abcd")
	bs := 4

	score := score(b, bs)

	if math.Abs(score-0.333333) > 0.00001 {
		t.Fatalf("Score was wrong: %f", score)
	}
}

func TestScoreNoDupplicates(t *testing.T) {
	b := []byte("12345678abcd")
	bs := 4

	score := score(b, bs)

	if math.Abs(score) > 0.00001 {
		t.Fatalf("Score of %b was wrong: %f", b, score)
	}
}

func TestScoreAllDuplicates(t *testing.T) {
	b := []byte("123412341234")
	bs := 4

	score := score(b, bs)

	if math.Abs(score-0.666667) > 0.00001 {
		t.Fatalf("Score was wrong: %f", score)
	}
}
