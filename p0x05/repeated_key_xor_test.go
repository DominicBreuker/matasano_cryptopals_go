package p0x05

import (
	"reflect"
	"testing"

	"github.com/dominicbreuker/matasano_cryptopals_go/utils"
)

func TestEncrypt(t *testing.T) {
	s := []byte(`Burning 'em, if you ain't quick and nimble
I go crazy when I hear a cymbal`)
	key := []byte("ICE")

	actual, _ := Encrypt(s, key)

	expected := utils.MustDecodeHexString("0b3637272a2b2e63622c2e69692a23693a2a3c6324202d623d63343c2a26226324272765272a282b2f20430a652e2c652a3124333a653e2b2027630c692b20283165286326302e27282f")
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Encrypt returns the ciphertext for the given key, but instead gave %s", actual)
	}
}

func TestEncryptEmptyKey(t *testing.T) {
	s := []byte("The plaintext")
	key := []byte("")

	_, err := Encrypt(s, key)

	expected := "Key must have at least one byte"
	if err.Error() != expected {
		t.Fatalf("If empty key is passed an error is returned")
	}
}

func TestEncryptEmptyCiphertext(t *testing.T) {
	s := []byte("")
	key := []byte("ABC")

	actual, _ := Encrypt(s, key)

	expected := []byte("")
	if !reflect.DeepEqual(actual, expected) {
		t.Fatal("Encrypting the empty plaintext gives an empty ciphertext")
	}
}

func TestMustEncrypt(t *testing.T) {
	s := []byte("The plaintext")
	key := []byte("ABC")

	actual := MustEncrypt(s, key)

	expected := utils.MustDecodeHexString("152a2661322f202b2d35273b35")
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Can encrypt this message correctly, but gave: %#x", actual)
	}
}

func TestMustEncryptPanicsOnError(t *testing.T) {
	s := []byte("The plaintext")
	key := []byte("")

	defer func() {
		if r := recover(); r == nil {
			t.Fatal("The code did not panic")
		}
	}()

	MustEncrypt(s, key)
}
