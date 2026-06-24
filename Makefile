APP_NAME := mq
CMD_DIR := ./cmd/mq
BIN_DIR := ./bin

.PHONY: build run clean

build:
	@mkdir -p $(BIN_DIR)
	@go build -o $(BIN_DIR)/$(APP_NAME) $(CMD_DIR)

run:
	@go run $(CMD_DIR)

clean:
	@rm -rf $(BIN_DIR)