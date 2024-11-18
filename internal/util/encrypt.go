package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"io"
	"log"
)

func Encrypt(data []byte, keystr []byte) ([]byte, error) {

	// keystr := make([]byte, 32)
	// if _, err := rand.Reader.Read(keystr); err != nil {
	// 	log.Println("Error on Creating KEY")
	// }
	// log.Println("KEY: " + string(keystr))

	key, _ := hex.DecodeString(string(keystr))
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Println("Error on Creating AES CIPHER")
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Println("Error on Creating GCM")
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		log.Println("Error on Creating NONCE")
	}

	cipherText := gcm.Seal(nonce, nonce, data, nil)

	return []byte(hex.EncodeToString(cipherText)), nil
}
