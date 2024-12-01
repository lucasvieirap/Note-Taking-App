package tests

import (
	"github.com/lucasvieirap/Note-Taking-App/internal/util"
	"testing"
)

func TestEncryptDecrypt(t *testing.T) {
	originalMsg := "Hello, World"

	key := "12345678901234567890123456789012"

	encryptedMsg, err := util.Encrypt([]byte(originalMsg), []byte(key))
	if err != nil {
		t.Fatalf(`util.Encrypt(%s, %s) returned an error: %v`, originalMsg, key, err)
	}

	decryptedMsg, err := util.Decrypt([]byte(encryptedMsg), []byte(key))
	if err != nil {
		t.Fatalf(`util.Decrypt(%s, %s) returned an error: %v`, encryptedMsg, key, err)
	}

	if string(decryptedMsg) != originalMsg {
		t.Fatalf(`Decrypt(%s, %s) = %s`, encryptedMsg, key, decryptedMsg)
	}
}
