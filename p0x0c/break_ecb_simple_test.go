package p0x0c

import (
	"math/rand"
	"reflect"
	"testing"

	"github.com/dominicbreuker/matasano_cryptopals_go/utils"
)

func TestNewOracle(t *testing.T) {
	oracle := getOracle()

	out := oracle([]byte("test-input"))
	if !reflect.DeepEqual(out, utils.MustDecodeHexString("18eb0562a27686b0718f0ee30d6a7e7246d499bbb60ecf035ffa22de96b1fa8e28156289ea97e85735fcb6b904835164a5e357e4980837539cfe0eedd2e5f6041c8ffc869772963511350f81d1262ee568c4f3d5af1b1dceafca5eacce959dfb379210ad0d604df0ad8d8ad878ca46403dc318b95a6451f3cb46df6b012debbcfb33989b2617e1a833c3c211d8e47f1a4e9c67bb6201bcc37e78e77697a67aad")) {
		t.Fatalf("Oracle did not return correct ciphertext: %#x", out)
	}
}

func TestDetectBlockSize(t *testing.T) {
	oracle := getOracle()

	bs, _ := detectBlockSize(oracle)
	if bs != 16 {
		t.Fatalf("Did not identify block size, got: %d", bs)
	}
}

func getOracle() Oracle {
	postfix := utils.MustReadBase64File("data.txt")
	r := rand.New(rand.NewSource(42))
	return NewOracle(r, postfix)
}
