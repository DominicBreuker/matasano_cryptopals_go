package p0x06

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"reflect"
	"testing"
)

func TestBreak(t *testing.T) {
	c := MustGetCiphertext()

	actual := Break(c, 2, 40)

	if string(actual) != "Terminator X: Bring the noise" {
		t.Fatalf("Key should be ... but was: %s", actual)
	}
}

func TestGetBlocks(t *testing.T) {
	b := []byte("a_sequence_of_bytes")
	ks := 3

	actual := getBlocks(b, ks)

	expected := map[int][]byte{
		0: []byte("aeeefys"),
		1: []byte("_qn__t"),
		2: []byte("sucobe"),
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Blocks not correctly computed: %+v", actual)
	}
}

func TestScoreKeySize(t *testing.T) {
	b := []byte("some_byte_sequence")
	k := 3

	actual := scoreKeysize(b, k)

	if math.Abs(actual-2.3333333) > 0.00001 {
		log.Fatalf("Score was %v", actual)
	}
}

func TestGuessKeysize(t *testing.T) {
	c := MustGetCiphertext()

	actual := guessKeysize(c, 3, 40)

	if actual != 29 {
		t.Fatalf("Best key size should be 10, but was: %d", actual)
	}
}

func TestEditDistance(t *testing.T) {
	b1 := []byte("this is a test")
	b2 := []byte("wokka wokka!!!")

	actual, _ := editDistance(b1, b2)

	expected := 37
	if actual != expected {
		t.Fatalf("Edit distance should be %d, but was %d", actual, expected)
	}
}

func TestEditDistanceDifferentBufferSize(t *testing.T) {
	b1 := []byte("this is a test")
	b2 := []byte("wokka wokk")

	_, err := editDistance(b1, b2)

	if err == nil {
		t.Fatalf("Can't compute edit distance of buffers with differing length, but got no error")
	}
}

// ##################
// ### Test utils ###
// ##################

func MustGetCiphertext() []byte {
	data := MustReadTestData()
	return MustBase64Decode(data)
}

func MustReadTestData() []byte {
	data, err := ioutil.ReadFile("data.txt")
	if err != nil {
		panic(fmt.Sprintf("Test file could not be read: %v", err))
	}
	return data
}

func MustBase64Decode(b []byte) []byte {
	b, err := base64.StdEncoding.DecodeString(string(b))
	if err != nil {
		panic(fmt.Sprintf("Test data could not be base64-decoded: %v", err))
	}
	return b
}
