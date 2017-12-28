package p0x07

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/dominicbreuker/matasano_cryptopals_go/utils"
)

func TestDecrypt(t *testing.T) {
	ct := utils.MustReadBase64File("data.txt")
	key := []byte("YELLOW SUBMARINE")

	pt := MustDecrypt(ct, key)

	if !reflect.DeepEqual(pt[:33], []byte("I'm back and I'm ringin' the bell")) {
		t.Fatalf("Wrong beginning of plaintext, got:\n%s\n", pt[:33])
	}
	if !reflect.DeepEqual(pt[len(pt)-23:], []byte("Play that funky music \n")) {
		t.Fatalf("Wrong end of plaintext, got:\n%b\n", pt[len(pt)-23:])
	}
}

func TestDecryptWrongKeysize(t *testing.T) {
	ct := []byte("this.is.the.ciphertext")
	key := []byte("keys.may.only.be.16.24.or.32.bytes.long")

	_, err := Decrypt(ct, key)

	if err == nil {
		t.Fatalf("Decrypt must not accept invalid keys")
	}
}

func TestECBEncryptThenDecrypt(t *testing.T) {
	pt := []byte("This is a sample plaintext")
	key := []byte("YELLOW SUBMARINE")

	ct := MustEncrypt(pt, key)
	pt2 := MustDecrypt(ct, key)

	if !reflect.DeepEqual(pt2, []byte("This is a sample plaintext")) {
		t.Fatalf("Did not recover plaintext '%s', got instead : '%s'", pt, pt2)
	}
}

// Test helper

func MustDecrypt(ct, key []byte) []byte {
	pt, err := Decrypt(ct, key)
	if err != nil {
		panic(fmt.Sprintf("Could not decrypt: %v", err))
	}
	return pt
}

func MustEncrypt(pt, key []byte) []byte {
	ct, err := Encrypt(pt, key)
	if err != nil {
		panic(fmt.Sprintf("Could not encrypt: %v", err))
	}
	return ct
}
