.PHONY: db_up clean build

BIN_FILE=lineServer
all: clean db_up build run
build:
	@go build -o "${BIN_FILE}" main.go
clean:
	@go clean
	@rm -f ./"${BIN_FILE}"
	@docker rm -f mongo
run:
	./"${BIN_FILE}"
db_up:
	docker-compose up -d
help:
	@echo "make: rebuild and run"
	@echo "make build: build binary file"
	@echo "make clean: delete the target file"
	@echo "make run: run lineBot"
