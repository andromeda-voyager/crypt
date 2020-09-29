package main

import (
	"crypto/rand"
	"fmt"
	"math"
)

const (
	letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numbers = "1234567890"
	symbols = "`~!@#$%^&*()_+-=[]{}\\|;:'\",.<>/?"
)

func createPassword(passwordLength int) []byte {
	alphabet := letters + numbers + symbols
	maxValue := len(alphabet)
	mask := createBitMask(len(alphabet))
	password := make([]byte, passwordLength)
	potentialBytes := randCryptoReadBytes(10)

	pIndex := 0
	bytesIndex := 0

	for pIndex < passwordLength {
		if b := int(potentialBytes[bytesIndex]) & mask; b < maxValue {
			password[pIndex] = alphabet[b]
			pIndex++
		}
		bytesIndex++
		if bytesIndex >= len(potentialBytes) {
			potentialBytes = randCryptoReadBytes(10)
			bytesIndex = 0
		}
	}
	return password
}

func randCryptoReadBytes(n int) []byte {
	b := make([]byte, 20)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("error:", nil)
		return nil
	}
	return b
}
func createBitMask(alphabetLenght int) int {
	bitCount := uint(math.Floor(math.Log2(float64(alphabetLenght)))) + 1
	mask := 1<<bitCount - 1
	return mask
}
