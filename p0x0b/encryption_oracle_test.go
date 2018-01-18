package p0x0b

import (
	"math/rand"
	"reflect"
	"testing"

	"github.com/dominicbreuker/matasano_cryptopals_go/utils"
)

func TestNewOracle(t *testing.T) {
	var tests = []struct {
		in     []byte
		out    []byte
		method bool
	}{
		{[]byte("test-input"), utils.MustDecodeHexString("9ed29a475b1ba2bc86457102e33e75153706d5ecb71a8e61298cf003dde4eaf5"), true},
		{[]byte("another-test-input"), utils.MustDecodeHexString("f770e6c5d2505796b5b6b7fd398cf158a7f37fd85cf5cd3d3df307783d77d10d"), false},
	}
	r := rand.New(rand.NewSource(42))

	for _, tt := range tests {
		oracle := NewOracle(r)
		out, method := oracle(tt.in)

		if !reflect.DeepEqual(out, tt.out) || method != tt.method {
			t.Fatalf("Should return %#x(%t) but gave %#x(%t)", tt.out, tt.method, out, method)
		}
	}
}

func TestDetectEBC(t *testing.T) {
	r := rand.New(rand.NewSource(42))

	for i := 0; i < 100; i++ {
		oracle := NewOracle(r)
		if err := DetectECB(oracle); err != nil {
			t.Fatalf("Failed to detect encryption method: %v", err)
		}
	}
}

func TestRandomBytes(t *testing.T) {
	r := rand.New(rand.NewSource(42))

	key := randomBytes(r, 16)

	if !reflect.DeepEqual(key, utils.MustDecodeHexString("538c7f96b164bf1b97bb9f4bb472e89f")) {
		t.Fatalf("Should return correct key but gave: %#x", key)
	}
}
