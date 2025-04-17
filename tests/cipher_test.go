package rum_test

import (
	"testing"

	"github.com/sidan-lab/rum"
)

func TestDecryptWithCipher(t *testing.T) {
	data := "solution solution solution solution solution solution solution solution solution solution solution solution solution solution solution solution solution solution solution solution solution solution solution solution"
	key := "01234567890123456789"

	encryptedData := `{"iv":"/bs1AzciZ1bDqT5W","ciphertext":"mh5pgH8ErqqH2KLLEBqqr8Pwm+mUuh9HhaAHslSD8ho6zk7mXccc9NUQAW8rb9UajCq8LYyANuiorjYD5N0hd2Lbe2n1x8AGRZrogyRKW6uhoFD1/FW6ofjgGP/kQRQSW2ZdJaDMbCxwYSdzxmaRunk6JRfybhfRU6kIxPMu41jhhRC3LbwZ+NnfBJFrg859hbuQgMQm8mqOUgOxcK8kKH54shOpGuLT4YBXhx33dZ//wT5VXrQ8kwIKttNk5h9MNKCacpRZSqU3pGlZ5oxucNEGos0IKTTXfbmwYx14uiERcXd32OP2"}`

	decryptedData, err := rum.DecryptWithCipher(encryptedData, key)
	if err != nil {
		t.Fatalf("Failed to decrypt: %v", err)
	}

	if decryptedData != data {
		t.Errorf("Decrypted data doesn't match original data.\nExpected: %s\nGot: %s", data, decryptedData)
	}
}

func TestEncryptAndDecryptWithCipher(t *testing.T) {
	data := "solution solution solution solution solution solution solution solution solution solution solution solution solution solution solution solution solution solution solution solution solution solution solution solution"
	key := "01234567890123456789"

	encryptedData, err := rum.EncryptWithCipher(data, key, 12)
	if err != nil {
		t.Fatalf("Failed to encrypt: %v", err)
	}

	decryptedData, err := rum.DecryptWithCipher(encryptedData, key)
	if err != nil {
		t.Fatalf("Failed to decrypt: %v", err)
	}

	if decryptedData != data {
		t.Errorf("Decrypted data doesn't match original data.\nExpected: %s\nGot: %s", data, decryptedData)
	}
}
