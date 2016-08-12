BIN_DIR = bin
CLIENT_DIR = ./client
SERVER_DIR = ./server

.PHONY: \
	default \
	clean \
	build_client \
	copy_client \
	build_server \
	copy_server

default: \
	clean \
	build_client \
	build_server \
	copy_client \
	copy_server

clean:
	rm -rf $(BIN_DIR) && mkdir -p $(BIN_DIR)
build_client:
	cd $(CLIENT_DIR) && $(MAKE)
copy_client:
	cp -r $(CLIENT_DIR)/$(BIN_DIR)/* $(BIN_DIR)/
build_server:
	cd $(SERVER_DIR) && $(MAKE)
copy_server:
	cp -r $(SERVER_DIR)/$(BIN_DIR)/* $(BIN_DIR)/
