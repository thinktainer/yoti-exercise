package main

import (
	c "github.com/thinktainer/yoti-exercise/crypt_contracts"
	"golang.org/x/net/context"
)

type clientContainer struct {
	client c.CryptClient
	ctx    context.Context
}

func (cont *clientContainer) Store(id string, payload []byte) (aesKey []byte, err error) {
	req := &c.EncryptRequest{
		Id:    id,
		Value: payload,
	}

	result, err := cont.client.Encrypt(cont.ctx, req)
	if err != nil {
		return nil, err
	}
	return result.Key, nil
}

func (cont *clientContainer) Retrieve(id string, aesKey []byte) (payload []byte, err error) {
	req := &c.DecryptRequest{
		Id:  id,
		Key: aesKey,
	}
	result, err := cont.client.Decrypt(cont.ctx, req)
	if err != nil {
		return nil, err
	}
	return result.Decrypted, nil
}
