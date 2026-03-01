APP := tsBootstrap
PKG := .
BIN_DIR := bin

.DEFAULT_GOAL := build

.PHONY: fmt build linux windows clean run

fmt:
	go fmt ./...

build: linux windows

linux: fmt
	mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/$(APP) $(PKG)

windows: fmt
	mkdir -p $(BIN_DIR)
	GOOS=windows GOARCH=amd64 go build -o $(BIN_DIR)/$(APP).exe $(PKG)

run: linux
	./$(BIN_DIR)/$(APP)

clean:
	rm -rf $(BIN_DIR)
