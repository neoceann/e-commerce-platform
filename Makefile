# Makefile
.PHONY: proto clean-proto sqlc

AUTH_DIR=./auth-service
SQLC_DIR=$(AUTH_DIR)/sqlc

proto:
#	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
#	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	
	protoc --go_out=$(AUTH_DIR) --go_opt=module=auth-service \
	       --go-grpc_out=$(AUTH_DIR) --go-grpc_opt=module=auth-service \
	       $(AUTH_DIR)/proto/auth.proto

clean-proto:
	rm -rf $(AUTH_DIR)/internal/grpc/pb/*.go

swag:
	swag init -g store/cmd/api/main.go -o store/docs

sqlc:
	cd $(SQLC_DIR) && sqlc generate