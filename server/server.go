package main

import (
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

	if req.Value == nil {
		return nil, grpc.Errorf(codes.InvalidArgument, "Value is a required field")
	}

	key, err := randomKey()
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, grpc.ErrorDesc(err))
	}

	encrypted, err := encrypt(key, req.Value)
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, grpc.ErrorDesc(err))
	}

	if ok := kvStore.add(hash(req.Id), encrypted); !ok {
		return nil, grpc.Errorf(codes.AlreadyExists, "Provided Id already exists")
	}
	return &c.EncryptResponse{
		Key: key,
	}, nil
}

func (s *server) Decrypt(ctx context.Context, req *c.DecryptRequest) (*c.DecryptResponse, error) {

	ciphertext, ok := kvStore.get(hash(req.Id))
	if !ok {
		return nil, grpc.Errorf(codes.NotFound, "Provided id not found in store")
	}

	decrypted, err := decrypt(req.Key, ciphertext)
	if err != nil {
		return nil, grpc.Errorf(codes.Internal, grpc.ErrorDesc(err))
	}

	return &c.DecryptResponse{
		Decrypted: decrypted,
	}, nil
}
