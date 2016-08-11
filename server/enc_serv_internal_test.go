package main

import (
	"bytes"
	"encoding/hex"
	"testing"
)

func Test_CanEncryptAndDecryptOK(t *testing.T) {

	sampleMessage := "a sample message"
	t.Logf("[Encrypting]: \"%s\"", sampleMessage)

	plaintext := []byte(sampleMessage)
	key, err := randomKey()
	if err != nil {
		t.Error("Error generating random key: %v", err)
	}

	ciphertext, err := encrypt(key, plaintext)

	if err != nil {
		t.Error("Error encrypting: %v", err)
	}

	decryptedMessage, err := decrypt(key, ciphertext)
	if err != nil {
		t.Logf("Error decrypting: %v", err)
	}

	t.Logf("[Decrypted Message]: \"%s\"", decryptedMessage)
	t.Logf("[Generated Key]: %s", hex.EncodeToString(key))
}

func Test_StoreAddAndRetrieveReturnsExpectedKey(t *testing.T) {
	st := kvStore
	key, expected := "key", []byte("value")
	st.add(key, expected)
	actual, _ := st.get(key)
	if !bytes.Equal(expected, actual) {
		t.Fail()
	}
}
