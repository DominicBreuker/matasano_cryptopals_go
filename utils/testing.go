package utils

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"
)

func MustReadBase64File(f string) []byte {
	data := mustReadTestData(f)
	return mustBase64Decode(data)
}

func mustReadTestData(f string) []byte {
	data, err := ioutil.ReadFile(f)
	if err != nil {
		panic(fmt.Sprintf("Test file could not be read: %v", err))
	}
	return data
}

func mustBase64Decode(b []byte) []byte {
	b, err := base64.StdEncoding.DecodeString(string(b))
	if err != nil {
		panic(fmt.Sprintf("Test data could not be base64-decoded: %v", err))
	}
	return b
}

// MustGetScanner returns a scanner to read a test file line by line. Do not forget to close the file afterwards
func MustGetScanner(f string) (*bufio.Scanner, *os.File) {
	file, err := os.Open(f)
	if err != nil {
		panic(fmt.Sprintf("Test file could not be opened: %v", err))
	}
	scanner := bufio.NewScanner(file)
	return scanner, file
}
