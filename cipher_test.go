package rum_test

import (
	"testing"

	"github.com/sidan-lab/rum"
)

func TestDecryptWithCipher(t *testing.T) {
	encryptedData := `{"iv":"XRAGv22SYgpZiGhy","salt":"5YowN2Txol1ejcvt9gJB1A==","ciphertext":"SUJcKVu5/yLVXvcVRI0xLTT+HN0j0JQc2YGL4uwmdErIAa4ZwTkfaKP3VNlclBeXoRfRqCRw9ioYZLSrZOsUlSKRDIGkrfHamZw3Nt+bTwWgzAecWmLOeU8Ks1ou6iQa1K9Yqt2+zJi6rDJfkEFEZJBOjC0iFnmeIMemYVD5UexqIkTlGZcKzwW57WU4HPKHpri/PhupcPRVpbZaNurCTB9tfnDLsr83zgHqSFILOdnSwvUaMA=="}`

	t.Run("correct password should decrypt successfully", func(t *testing.T) {
		password := "testing123456"

		result, err := rum.DecryptWithCipher(encryptedData, password)

		if err != nil {
			t.Errorf("expected no error, got: %v", err)
		}
		if result == "" {
			t.Error("expected non-empty result, got empty string")
		}
		t.Logf("Decrypted successfully, result length: %d", len(result))
	})

	t.Run("incorrect password should fail", func(t *testing.T) {
		password := "wrongPassword"

		result, err := rum.DecryptWithCipher(encryptedData, password)

		if err == nil {
			t.Error("expected error with wrong password, got nil")
		}
		if result != "" {
			t.Errorf("expected empty result on error, got: %s", result)
		}
		t.Logf("Failed as expected with error: %v", err)
	})
}
