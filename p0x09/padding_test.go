package p0x09

import (
	"reflect"
	"testing"
)

func TestPad(t *testing.T) {
	b := []byte("a.buffer.of.length.21")
	bs := 16

	b = Pad(b, bs)

	if !reflect.DeepEqual(b, []byte("a.buffer.of.length.21\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b\x0b")) {
		t.Fatalf("Should padd 11x byte 11 (0x0b), but got: %v", b)
	}
}

func TestPadMultipleOfBlockSize(t *testing.T) {
	b := []byte("a.buffer.of.leng") // len 16
	bs := 16

	b = Pad(b, bs)

	if !reflect.DeepEqual(b, []byte("a.buffer.of.leng\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10\x10")) {
		t.Fatalf("Should padd 16x byte 16 (0x10), but got: %v", b)
	}
}

func TestPadOnlyOne(t *testing.T) {
	b := []byte("a.buffer.of.len") // len 15
	bs := 16

	b = Pad(b, bs)

	if !reflect.DeepEqual(b, []byte("a.buffer.of.len\x01")) {
		t.Fatalf("Should padd 1x byte 1 (0x01), but got: %v", b)
	}
}

func TestStripPad(t *testing.T) {
	b := []byte("a.buffer.with.padding\x03\x03\x03")

	b = StripPad(b)

	if !reflect.DeepEqual(b, []byte("a.buffer.with.padding")) {
		t.Fatalf("Should remove padding bytes but got: %v", b)
	}
}

func TestStripPadEmptyString(t *testing.T) {
	b := []byte("")

	b = StripPad(b)

	if !reflect.DeepEqual(b, []byte("")) {
		t.Fatalf("Should return empty string but gave: %v", b)
	}
}
