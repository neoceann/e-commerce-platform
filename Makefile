# Makefile
AUTH_DIR=./auth-service
STORE_DIR=./store
SQLC_DIR=$(AUTH_DIR)/sqlc

docker-build: proto docker-clean
	docker-compose build

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

proto:
	protoc --go_out=$(AUTH_DIR) --go_opt=module=auth-service \
	       --go-grpc_out=$(AUTH_DIR) --go-grpc_opt=module=auth-service \
	       $(AUTH_DIR)/proto/auth.proto

	cp -r $(AUTH_DIR)/internal/grpc/pb $(STORE_DIR)/internal/grpc/

clean-proto:
	rm -rf $(AUTH_DIR)/internal/grpc/pb/*.go
	rm -rf $(STORE_DIR)/internal/grpc/pb/*.go

docker-clean:
	docker-compose down -v
	docker system prune -f
	
swag:
	swag init -g store/cmd/api/main.go -o store/docs

sqlc:
	cd $(SQLC_DIR) && sqlc generate


.PHONY: proto clean-proto sqlc swag docker-build docker-up docker-down docker-clean