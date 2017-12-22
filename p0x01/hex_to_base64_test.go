package p0x01

import (
	"strings"
	"testing"
)

func TestHexToBase64(t *testing.T) {
	s := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"

	actual, _ := HexToBase64(s)

	expected := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"
	if actual != expected {
		t.Fatalf("Converts hex strings to standard base64, but delivered: actual %v / expected %v", actual, expected)
	}
}

func TestHexToBase64InvalidHexString(t *testing.T) {
	s := "asdlfjslndsldzfndjfa"

	_, err := HexToBase64(s)

	expected := "Can't decode hex string: "
	if !strings.Contains(err.Error(), expected) {
		t.Fatalf("Argument must be valid hex string")
	}
}
