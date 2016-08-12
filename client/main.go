package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	c "github.com/thinktainer/yoti-exercise/crypt_contracts"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"os"
	"strings"
)

const (
	serverAddress = ":50001"
)

type CommandType int

const (
	EncryptCommand CommandType = iota
	DecryptCommand
	Unrecognized
)

type Command struct {
	Type    CommandType
	Payload []string
}

func parse(input string) *Command {
	tokens := strings.Split(input, " ")
	if len(tokens) < 3 {
		return &Command{
			Type: Unrecognized,
		}
	}

	if strings.ToUpper(tokens[0]) == "STORE" {
		if len(tokens[1:]) < 2 {
			return &Command{
				Type: Unrecognized,
			}
		}

		return &Command{
			Type:    EncryptCommand,
			Payload: tokens[1:],
		}
	}

	if strings.ToUpper(tokens[0]) == "RETRIEVE" {
		if len(tokens[1:]) < 2 {
			return &Command{
				Type: Unrecognized,
			}
		}

		return &Command{
			Type:    DecryptCommand,
			Payload: tokens[1:],
		}
	}

	return &Command{
		Type: Unrecognized,
	}
}

func handle(client clientContainer, command *Command) (result interface{}, err error) {
	log.Println("Handling command: %v", &command)
	switch command.Type {
	case DecryptCommand:
		decoded, err := hex.DecodeString(command.Payload[1])
		if err != nil {
			return nil, err
		}
		key := []byte(decoded)
		res, err := client.Retrieve(command.Payload[0], key)
		if err != nil {
			return nil, err
		}
		return res, nil
	case EncryptCommand:
		res, err := client.Store(command.Payload[0], []byte(command.Payload[1]))
		if err != nil {
			return nil, err
		}

		return res, nil

	default:
		return nil, fmt.Errorf("Unhandled command type.")
	}

	panic("You shall not pass!")
}

func main() {

	log.Println("Establishing connection")
	conn, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	cryptClient := c.NewCryptClient(conn)
	ccont := clientContainer{
		client: cryptClient,
		ctx:    context.Background(),
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		command := parse(text)
		if command.Type == Unrecognized {
			log.Println("Unrecognized input: %s", text)
			continue
		}

		res, err := handle(ccont, command)
		if err != nil {
			panic(err)
		}

		switch res := res.(type) {
		case string:
			fmt.Printf("Got string response from server: [%s]\n", res)
		case []byte:
			switch {
			case command.Type == EncryptCommand:
				fmt.Printf("Got key response from server: [%s]\n", hex.EncodeToString(res))
			case command.Type == DecryptCommand:
				fmt.Printf("Got decrypt response from server: [%s]\n", string(res))
			default:
				fmt.Printf("Got unknown response from server: [%v]\n", res)
			}
		default:
			fmt.Printf("Got unknown response from server: [%v]\n", res)
		}
	}
}
