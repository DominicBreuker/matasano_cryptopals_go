package p0x03

const baseScore = 0.1

// https://en.wikipedia.org/wiki/Letter_frequency
var scores = func() map[rune]float64 {
	freq := map[rune]float64{
		'e': 12.70,
		't': 9.06,
		'a': 8.17,
		'o': 7.51,
		'i': 6.97,
		'n': 6.75,
		's': 6.33,
		'h': 6.09,
		'r': 5.99,
		'd': 4.25,
		'l': 4.03,
		'c': 2.78,
		'u': 2.76,
		'm': 2.41,
		'w': 2.36,
		'f': 2.23,
		'g': 2.02,
		'y': 1.97,
		'p': 1.93,
		'b': 1.29,
		'v': 0.98,
		'k': 0.77,
		'j': 0.15,
		'x': 0.15,
		'q': 0.10,
		'z': 0.07,
	}
	others := []rune{' ', '.', ',', ':', ';', '?', '(', ')', '[', ']', '{', '}', '#', '$', '*', '\''}

	var result = make(map[rune]float64)
	for r, f := range freq {
		result[r] = f/100. + baseScore
	}
	for _, r := range others {
		result[r] = baseScore
	}

	return result
}()
