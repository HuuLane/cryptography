package otp

// one-time pad

import "math/rand"

func Encrypt(plaintext []byte) ([]byte, []byte) {
	key := make([]byte, len(plaintext))
	rand.Read(key)

	ciphertext := make([]byte, len(plaintext))
	for i := 0; i < len(plaintext); i++ {
		ciphertext[i] = plaintext[i] ^ key[i]
	}
	return key, ciphertext
}

func Decrypt(key []byte, ciphertext []byte) []byte {
	if len(key) != len(ciphertext) {
		panic("len of key and ciphertext must be equal")
	}
	plaintext := make([]byte, len(key))
	for i := 0; i < len(ciphertext); i++ {
		plaintext[i] = ciphertext[i] ^ key[i]
	}
	return plaintext
}
