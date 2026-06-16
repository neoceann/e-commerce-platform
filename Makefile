# Makefile
.PHONY: proto clean-proto sqlc swag run build

AUTH_DIR=./auth-service
STORE_DIR=./store
SQLC_DIR=$(AUTH_DIR)/sqlc

build:
	cd $(AUTH_DIR) && go build -o build/auth ./cmd/auth/
	cd $(STORE_DIR) && go build -o build/store ./cmd/api/

run: build
	cp .env $(AUTH_DIR)/build
	cp .env $(STORE_DIR)/build
	cd $(AUTH_DIR) && ./build/auth
	cd $(STORE_DIR) && ./build/store

dbuild: proto
	docker-compose build

up: build
	docker-compose up -d

down:
	docker-compose down

proto:
	protoc --go_out=$(AUTH_DIR) --go_opt=module=auth-service \
	       --go-grpc_out=$(AUTH_DIR) --go-grpc_opt=module=auth-service \
	       $(AUTH_DIR)/proto/auth.proto

	cp -r $(AUTH_DIR)/internal/grpc/pb $(STORE_DIR)/internal/grpc/

clean-proto:
	rm -rf $(AUTH_DIR)/internal/grpc/pb/*.go
	rm -rf $(STORE_DIR)/internal/grpc/pb/*.go

swag:
	swag init -g store/cmd/api/main.go -o store/docs

sqlc:
	cd $(SQLC_DIR) && sqlc generate