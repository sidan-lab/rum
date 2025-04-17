package rum

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"

	"golang.org/x/crypto/pbkdf2"
)

// CipherData represents the structure of encrypted data
type CipherData struct {
	IV         string `json:"iv"`
	Ciphertext string `json:"ciphertext"`
}

// EncryptWithCipher encrypts data using PBKDF2 key derivation and AES-GCM
func EncryptWithCipher(data, key string, initializationVectorSize int) (string, error) {
	// Create an initialization vector
	iv := make([]byte, initializationVectorSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	// Create a salt with the same length as the IV (all zeros)
	salt := make([]byte, initializationVectorSize)

	// Derive key using PBKDF2 with SHA-256
	derivedKey := pbkdf2.Key([]byte(key), salt, 100000, 32, sha256.New)

	// Create cipher block
	block, err := aes.NewCipher(derivedKey)
	if err != nil {
		return "", err
	}

	// Use GCM mode
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// Encrypt the data
	ciphertext := aesGCM.Seal(nil, iv, []byte(data), nil)

	// Create CipherData structure
	cipherData := CipherData{
		IV:         base64.StdEncoding.EncodeToString(iv),
		Ciphertext: base64.StdEncoding.EncodeToString(ciphertext),
	}

	// Marshal to JSON
	encryptedJSON, err := json.Marshal(cipherData)
	if err != nil {
		return "", err
	}

	return string(encryptedJSON), nil
}

// DecryptWithCipherPBKDF2 decrypts data using PBKDF2 key derivation
func DecryptWithCipher(encryptedDataJSON, key string) (string, error) {
	// Parse the JSON
	var cipherData CipherData
	if err := json.Unmarshal([]byte(encryptedDataJSON), &cipherData); err != nil {
		return "", err
	}

	// Decode base64
	iv, err := base64.StdEncoding.DecodeString(cipherData.IV)
	if err != nil {
		return "", err
	}
	ciphertext, err := base64.StdEncoding.DecodeString(cipherData.Ciphertext)
	if err != nil {
		return "", err
	}

	// Create a salt with the same length as the IV (all zeros)
	salt := make([]byte, len(iv))

	// Derive key using PBKDF2 with SHA-256
	derivedKey := pbkdf2.Key([]byte(key), salt, 100000, 32, sha256.New)

	// Create cipher block
	block, err := aes.NewCipher(derivedKey)
	if err != nil {
		return "", err
	}

	// Use GCM mode (default in the JS implementation)
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// Decrypt the data
	plaintext, err := aesGCM.Open(nil, iv, ciphertext, nil)
	if err != nil {
		return "", errors.New("decryption failed: " + err.Error())
	}

	return string(plaintext), nil
}
