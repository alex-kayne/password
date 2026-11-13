package encrypter

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
	"os"
)

type Encrypter struct {
	key []byte
}

func NewEncrypter() (*Encrypter, error) {
	key, exitsts := os.LookupEnv("ENCRYPT_KEY")
	if !exitsts {
		return nil, errors.New("переменная ENCRYPT_KEY не задана в окружении")
	}
	return &Encrypter{key: []byte(key)}, nil
}

func (e *Encrypter) Encrypt(text string) (string, error) { // 1. Создаем блочный шифр AES
	block, err := aes.NewCipher(e.key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	ciphertext := gcm.Seal(nonce, nonce, []byte(text), nil)
	return string(ciphertext), nil
}

func (e *Encrypter) Decrypt(text string) (string, error) {
	block, err := aes.NewCipher(e.key)
	ciphertext := []byte(text)

	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return "", errors.New("зашифрованные данные слишком короткие")
	}

	nonce, encryptedData := ciphertext[:nonceSize], ciphertext[nonceSize:]

	plaintext, err := gcm.Open(nil, nonce, encryptedData, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}
