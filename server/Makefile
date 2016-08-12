.PHONY : default build test clean watch

BIN_DIR = ./bin

default : build

build : clean
	go build -o $(BIN_DIR)/server
test :
	go test -v -cover "-run=^Test(Server|Store)_"
clean :
	rm -rf $(BIN_DIR)
watch :
	nodemon -e go -x "make test || true" 
