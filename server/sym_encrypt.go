package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"io"
)

// Generates a random key for aes encryption
func randomKey() ([]byte, error) {
	key := make([]byte, 24) // arbitrarily chosen length
	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		return nil, err
	}

	return key, nil
}

// Encrypt encrypts a plaintext string using AES CFB encryption
// It returns the randomly generated encryption key as a raw byte array
// and the ciphertext (encrypted string), with the IV at the start
func encrypt(encryption_key, plaintext []byte) (ciphertext []byte, err error) {
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(encryption_key)
	if err != nil {
		return nil, err
	}

	ciphertext = make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	mode := cipher.NewCFBEncrypter(block, iv)
	mode.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)
	return ciphertext, nil
}

// Decrypt decrypts a ciphertext using a specified key
func decrypt(key, ciphertext []byte) (decrypted []byte, err error) {

	iv := ciphertext[:aes.BlockSize]
	if len(iv) < aes.BlockSize {
		return nil, errors.New("ciphertext too short to contain IV")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	decrypted = []byte(ciphertext[aes.BlockSize:])
	mode := cipher.NewCFBDecrypter(block, iv)
	mode.XORKeyStream(decrypted, ciphertext[aes.BlockSize:])
	return decrypted, nil
}

func hash(key string) string {
	sha := sha256.New()
	hashed := sha.Sum([]byte(key))
	return string(hashed)
}
