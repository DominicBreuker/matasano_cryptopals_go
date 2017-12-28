package p0x0a

import (
	"reflect"
	"strings"
	"testing"

	"github.com/dominicbreuker/matasano_cryptopals_go/utils"
)

func TestCBCDecrypt(t *testing.T) {
	ct := utils.MustReadBase64File("data.txt")
	key := []byte("YELLOW SUBMARINE")
	iv := make([]byte, 16)

	pt, _ := Decrypt(ct, key, iv)

	if !reflect.DeepEqual(pt[:33], []byte("I'm back and I'm ringin' the bell")) {
		t.Fatalf("Beginning of plaintext not correctly decrypted: %s", pt[:33])
	}
	if !reflect.DeepEqual(pt[len(pt)-23:], []byte("Play that funky music \n")) {
		t.Fatalf("End of plaintext not correctly decrypted: %sb", pt[len(pt)-23:])
	}
}

func TestCBCDecryptEmptyCiphertext(t *testing.T) {
	ct := []byte("")
	key := []byte("YELLOW SUBMARINE")
	iv := make([]byte, 16)

	pt, _ := Decrypt(ct, key, iv)

	if !reflect.DeepEqual(pt, []byte("")) {
		t.Fatalf("plaintext not correctly decrypted: %s", pt)
	}
}

func TestCBCDecryptInvalidKey(t *testing.T) {
	ct := utils.MustReadBase64File("data.txt")
	key := []byte("this.is.a.key.that.has.an.invalid.length")
	iv := make([]byte, len(key))

	_, err := Decrypt(ct, key, iv)

	if !strings.Contains(err.Error(), "Could not create AES cipher from key") {
		t.Fatalf("Should return correct error but gave: %s", err.Error())
	}
}

func TestCBCDecryptInvalidIV(t *testing.T) {
	ct := utils.MustReadBase64File("data.txt")
	key := []byte("YELLOW SUBMARINE")
	iv := make([]byte, len(key)+1)

	_, err := Decrypt(ct, key, iv)

	if !strings.Contains(err.Error(), "IV does not equal block size") {
		t.Fatalf("Should return correct error but gave: %s", err.Error())
	}
}

func TestCBCDecryptWrongKey(t *testing.T) {
	ct := utils.MustReadBase64File("data.txt")
	key := []byte("WRONG KEY xxxxxz")
	iv := make([]byte, len(key))

	_, _ = Decrypt(ct, key, iv)
	t.Skip("TODO: check if decrypt was successful and return error if not")
}

func TestCBCEncryptThenDecrypt(t *testing.T) {
	pt := []byte("This is a sample plaintext")
	key := []byte("YELLOW SUBMARINE")
	iv := make([]byte, 16)

	ct, _ := Encrypt(pt, key, iv)
	pt2, _ := Decrypt(ct, key, iv)

	if !reflect.DeepEqual(pt2, []byte("This is a sample plaintext")) {
		t.Fatalf("Did not recover plaintext '%s', got instead : '%s'", pt, pt2)
	}
}
