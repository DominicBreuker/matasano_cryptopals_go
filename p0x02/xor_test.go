package p0x02

import (
	"encoding/hex"
	"reflect"
	"testing"
)

func TestXOR(t *testing.T) {
	b1, _ := hex.DecodeString("1c0111001f010100061a024b53535009181c")
	b2, _ := hex.DecodeString("686974207468652062756c6c277320657965")

	actual, _ := XOR(b1, b2)

	expected, _ := hex.DecodeString("746865206b696420646f6e277420706c6179")
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("XOR takes the XOR of two byte sequences, but delivered: actual %v / expected %v", actual, expected)
	}
}

func TestXORDifferentLength(t *testing.T) {
	b1, _ := hex.DecodeString("1c0111001f010100061a024b5353500918")
	b2, _ := hex.DecodeString("686974207468652062756c6c277320657965")

	_, err := XOR(b1, b2)

	expected := "Can't XOR buffers of different length"
	if err.Error() != expected {
		t.Fatalf("XORed hex strings must be of the same length")
	}
}
