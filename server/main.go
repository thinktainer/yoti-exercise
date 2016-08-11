package main

import (
	"log"
	"net"

	c "github.com/thinktainer/yoti-exercise/crypt_contracts"
	"google.golang.org/grpc"
)

const (
	listenAddress = ":50001"
)

func main() {

	lis, err := net.Listen("tcp", listenAddress)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	c.RegisterCryptServer(s, &server{})
	s.Serve(lis)
}
