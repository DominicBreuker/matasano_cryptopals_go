package p0x0b

import (
	"reflect"
	"testing"
)

func TestRandomKey(t *testing.T) {
	key1, _ := randomKey(16)
	key2, _ := randomKey(16)

	// TODO: find way to fix seed for proper test
	if reflect.DeepEqual(key1, key2) {
		t.Fatalf("Should return random keys but returned the same key in two calls: %b", key1)
	}
}
