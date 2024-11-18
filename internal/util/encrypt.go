package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"log"
	"io"
)

func encrypt(data []byte, keystr string) ([]byte, error) {
	key, err := hex.DecodeString(keystr)
	if err != nil {
		log.Println("Error on decoding Keystring")
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Println("Error on creating aes cipher")
	}
	
	ciphertext := make([]byte, aes.BlockSize+len(data))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return []byte(""), err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], data)

	return ciphertext, nil
}
