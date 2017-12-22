package p0x04

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/dominicbreuker/matasano_cryptopals_go/p0x03"
)

func TestDetect(t *testing.T) {
	file, _ := os.Open("data.txt")

	detector := New()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.TrimSuffix(scanner.Text(), "\n")
		b, _ := hex.DecodeString(s)
		detector.Update(b)
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Strange error: %v", err)
	}
	file.Close()

	actual, _, _ := p0x03.Decrypt(detector.Best)

	expected := []byte("Now that the party is jumping\n")
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("The ciphertext decrypts to '%s' but we found '%s'", expected, actual)
	}
}
