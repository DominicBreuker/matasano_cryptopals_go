package p0x09

func Pad(b []byte, bs int) []byte {
	m := bs - len(b)%bs

	pad := make([]byte, m)
	for i := 0; i < m; i++ {
		pad[i] = byte(m)
	}

	return append(b, pad...)
}

func StripPad(b []byte) []byte {
	if len(b) == 0 {
		return b
	}
	pad := int(b[len(b)-1])
	return b[:len(b)-pad]
}
