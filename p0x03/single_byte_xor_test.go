package p0x03

import (
	"encoding/hex"
	"math"
	"reflect"
	"testing"
)

func TestScore(t *testing.T) {
	b := []byte("abcd")

	actual := score(b)

	expected := 0.141225
	if math.Abs(actual-expected) > 0.0001 {
		t.Fatalf("Score of 'abcd' is %v, but was %v", expected, actual)
	}
}

func TestDecrypt(t *testing.T) {
	b, _ := hex.DecodeString("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")

	actual, _, _ := Decrypt(b)

	expected := []byte("Cooking MC's like a pound of bacon")
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Correct plaintext not recovered! Got '%s' but should '%s'", actual, expected)
	}
}
