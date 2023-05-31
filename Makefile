include .env

# Variables for binding arguments from cmd
MIGRATION_FILE_NAME?=default_migration_file_name

.PHONY:

server:
	go run main.go server

client:
	go run main.go client

proto:
	rm -f pkg/pb/*.go
	protoc --proto_path=pkg/proto --go_out=pkg/pb --go_opt=paths=source_relative \
	--go-grpc_out=pkg/pb --go-grpc_opt=paths=source_relative \
	pkg/proto/*.proto