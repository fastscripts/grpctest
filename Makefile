APP_NAME: grpctest

export PATH := $(PATH):$(GOPATH)/bin

.PHONY: proto-gen

proto-gen:
	/usr/local/bin/protoc \
	 --go_out=gen/go \
	 --go_opt=paths=source_relative \
	 --go-grpc_out=gen/go \
	 --go-grpc_opt=paths=source_relative \
	 proto/hello/v1/hello.proto

	/usr/local/bin/protoc \
	 --go_out=gen/go \
	 --go_opt=paths=source_relative \
	 --go-grpc_out=gen/go \
	 --go-grpc_opt=paths=source_relative \
	 proto/user/v1/user.proto
	


# Migration targets
migrate-up: ## Run migrations up.
	@if [ -z "$(dsn)" ]; then \
		echo "Usage: make migrate-up dsn=\"postgres://postgres:password@localhost:5432/boilerplate?sslmode=disable\""; \
		exit 1; \
	fi

	@echo "Running migrations up on DSN=$(dsn)"
	@migrate -path $(MIGRATIONS_DIR) -database $(dsn) up

migrate-down: ## Run migrations down (be careful!).
	@if [ -z "$(dsn)" ]; then \
		echo "Usage: make migrate-down dsn=\"postgres://postgres:password@localhost:5432/boilerplate?sslmode=disable\""; \
		exit 1; \
	fi

	@echo "Running migrations down on $(DATABASE_NAME) with DSN=$(DATABASE_DSN)"
	@migrate -path $(MIGRATIONS_DIR) -database $(dsn) down

migrate-new: ## Create a new migration file. 
	@if [ -z "$(name)" ]; then \
		echo "Usage: make migrate-new name=add_users_table"; \
		exit 1; \
	fi

	@migrate create -ext sql -dir $(MIGRATIONS_DIR) -seq $(name)

seed: ## Run all seeders.
	go run cmd/cli/main.go seed $(dsn)