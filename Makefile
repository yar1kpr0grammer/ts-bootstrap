APP := tsBootstrap
BIN_DIR := bin
MAIN := ./cmd/tsBootstrap

.DEFAULT_GOAL := build

.PHONY: fmt tidy build build-linux build-windows build-mac run clean

# --- format & deps ---
fmt:
	go fmt ./...

tidy:
	go mod tidy

# --- build ---
build: fmt tidy build-linux build-windows build-mac

build-linux:
	mkdir -p $(BIN_DIR)
	GOOS=linux GOARCH=amd64 go build -o $(BIN_DIR)/$(APP)-linux $(MAIN)

build-windows:
	mkdir -p $(BIN_DIR)
	GOOS=windows GOARCH=amd64 go build -o $(BIN_DIR)/$(APP).exe $(MAIN)

build-mac:
	mkdir -p $(BIN_DIR)
	GOOS=darwin GOARCH=amd64 go build -o $(BIN_DIR)/$(APP)-mac $(MAIN)

# --- run ---
run: fmt
	go run $(MAIN)

# --- clean ---
clean:
	rm -rf $(BIN_DIR)