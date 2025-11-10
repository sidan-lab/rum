package rum

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
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

func DecryptWithCipher(encryptedDataJSON string, password string) (string, error) {
	// Parse the encrypted data JSON
	var encData struct {
		IV         string  `json:"iv"`
		Salt       *string `json:"salt,omitempty"`
		Ciphertext string  `json:"ciphertext"`
	}

	err := json.Unmarshal([]byte(encryptedDataJSON), &encData)
	if err != nil {
		return "", fmt.Errorf("failed to parse encrypted data: %w", err)
	}

	// Decode IV from base64
	iv, err := base64.StdEncoding.DecodeString(encData.IV)
	if err != nil {
		return "", fmt.Errorf("failed to decode IV: %w", err)
	}

	// Decode ciphertext from base64
	ciphertext, err := base64.StdEncoding.DecodeString(encData.Ciphertext)
	if err != nil {
		return "", fmt.Errorf("failed to decode ciphertext: %w", err)
	}

	// Handle salt - support both new format (with salt) and legacy format (without salt)
	var salt []byte
	if encData.Salt != nil && *encData.Salt != "" {
		// New format: use the provided salt
		salt, err = base64.StdEncoding.DecodeString(*encData.Salt)
		if err != nil {
			return "", fmt.Errorf("failed to decode salt: %w", err)
		}
	} else {
		// Legacy format: use zero-filled salt of IV length for backward compatibility
		salt = make([]byte, len(iv))
	}

	// Derive cryptographic key from password using PBKDF2
	// Matches frontend: 100,000 iterations, SHA-256, 256-bit key
	derivedKey := pbkdf2.Key([]byte(password), salt, 100000, 32, sha256.New)

	// Create AES cipher
	block, err := aes.NewCipher(derivedKey)
	if err != nil {
		return "", fmt.Errorf("failed to create cipher: %w", err)
	}

	// Create GCM mode
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("failed to create GCM: %w", err)
	}

	// Decrypt the data
	plaintext, err := aesgcm.Open(nil, iv, ciphertext, nil)
	if err != nil {
		return "", fmt.Errorf("failed to decrypt (incorrect password or corrupted data): %w", err)
	}

	// Return the decrypted data as string
	return string(plaintext), nil
}
