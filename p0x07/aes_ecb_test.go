package p0x07

import (
	"reflect"
	"testing"

	"github.com/dominicbreuker/matasano_cryptopals_go/utils"
)

func TestDecrypt(t *testing.T) {
	c := utils.MustReadBase64File("data.txt")
	k := []byte("YELLOW SUBMARINE")

	Decrypt(c, k)
	actual := c[:33] // inplace decrypt, take only first few bytes

	expected := []byte("I'm back and I'm ringin' the bell")
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Wrong plaintext, got:\n%s\n", actual)
	}
}

func TestDecryptWrongKeysize(t *testing.T) {
	c := []byte("this.is.the.ciphertext")
	k := []byte("keys.may.only.be.16.24.or.32.bytes.long")

	err := Decrypt(c, k)

	if err == nil {
		t.Fatalf("Decrypt must not accept invalid keys")
	}
}
