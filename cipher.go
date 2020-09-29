package main

import (
	"crypto/aes"
	"crypto/cipher"

	"golang.org/x/crypto/argon2"
)

const (
	nonceSize = 12
	saltSize  = 44
)

func newCipher(password, salt []byte) (cipher.AEAD, error) {

	kdf := argon2.IDKey(password, salt, 4, 32*1024, 4, 44)

	block, err := aes.NewCipher(kdf[:32])
	//fmt.Println(kdf)
	if err != nil {
		return nil, err
	}

	aead, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	return aead, nil
}

func decrypt(password, ciphertext []byte) ([]byte, error) {

	nonce := ciphertext[:nonceSize]
	salt := ciphertext[nonceSize : saltSize+nonceSize]
	ciphertext = ciphertext[saltSize+nonceSize:]
	cipher, err := newCipher(password, salt)
	if err != nil {
		return nil, err
	}
	return cipher.Open(nil, nonce, ciphertext, nil)

}

func encrypt(password, plaintext []byte) ([]byte, error) {
	nonce, err := generateRandomBytes(12)
	salt, err := generateRandomBytes(44)
	cipher, err := newCipher(password, salt)
	if err != nil {
		return nil, err
	}

	prependedData := append(nonce, salt...)
	return cipher.Seal(prependedData, nonce, plaintext, nil), nil
}
