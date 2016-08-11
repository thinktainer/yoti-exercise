package main

import (
	"encoding/hex"
	c "github.com/thinktainer/yoti-exercise/crypt_contracts"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

type server struct{}

func (s *server) Encrypt(ctx context.Context, req *c.EncryptRequest) (*c.EncryptResponse, error) {
	if req.Id == "" {
		return nil, grpc.Errorf(codes.InvalidArgument, "Id is a required field")
	}

	if req.Value == "" {
		return nil, grpc.Errorf(codes.InvalidArgument, "Value is a required field")
	}

	key, err := randomKey()
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, grpc.ErrorDesc(err))
	}

	encrypted, err := encrypt(key, []byte(req.Value))
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, grpc.ErrorDesc(err))
	}
	hashedKey := hex.EncodeToString(hash(key))
	kvStore.add(string(hash([]byte(req.Id))), encrypted)
	return &c.EncryptResponse{
		Key: hashedKey,
	}, nil
}

func (s *server) Decrypt(ctx context.Context, req *c.DecryptRequest) (*c.DecryptResponse, error) {

	ciphertext, ok := kvStore.get(string(hash([]byte(req.Id))))
	if !ok {
		return nil, grpc.Errorf(codes.NotFound, "Provided id not found in store")
	}

	key, err := hex.DecodeString(req.Key)
	if err != nil {
		return nil, grpc.Errorf(codes.InvalidArgument, grpc.ErrorDesc(err))
	}

	decrypted, err := decrypt(key, ciphertext)
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, grpc.ErrorDesc(err))
	}

	return &c.DecryptResponse{
		Decrypted: string(decrypted),
	}, nil
}
