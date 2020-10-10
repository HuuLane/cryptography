package otp

import (
	"fmt"
	"math/big"
	"testing"
)

func TestOTP(t *testing.T) {
	s := "Hello, 你好世界"
	key, ciphertext := Encrypt([]byte(s))
	fmt.Printf("%s\n", ciphertext)
	fmt.Println(key)
	fmt.Printf("%s\n", Decrypt(key, ciphertext))
}

func TestBFOTP(t *testing.T) {
	s := "Hello, 你好世界"
	_, ciphertext := Encrypt([]byte(s))
	keyspace := big.NewInt(1)
	keyspace.Lsh(keyspace, uint(len(s)))

	fmt.Printf("input len: %d\n", len(s))
	fmt.Printf("ciphertext: %s\n", ciphertext)
	fmt.Println("keyspace: ", keyspace.String())

	// todo: how to list all possible keys
	// key := make([]byte, len(ciphertext))
	// for i := 0; i < 1000; i++ {
	// 	rand.Read(key)
	// 	fmt.Printf("%s\n", Decrypt(key, ciphertext))
	// }
	// fmt.Println(key)
}
