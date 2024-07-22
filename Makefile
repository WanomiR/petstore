build:
	@cd backend
	@swag init -g ./cmd/api/main.go
	@cd ..
	@docker compose up --build --force-recreate

read-env:
	@. ./read-env.sh

# make read-env && make migration_create
migration_create:
	@migrate create -ext sql -dir postgres/migration/ -seq init_mg

migration_up:
	@migrate -path postgres/migration/ -database "postgresql://$(DB_USER):$(DB_PASSWORD)@$(HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" -verbose up

migration_down:
	@migrate -path postgres/migration/ -database "postgresql://$(DB_USER):$(DB_PASSWORD)@$(HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" -verbose down

migration_fix:
	@migrate -path postgres/migration/ -database "postgresql://$(DB_USER):$(DB_PASSWORD)@$(HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" force $(VERSION)

# make postgres_dump CONT_NAME=users-migration-postgres-1
postgres_dump:
	@docker exec -t $(CONT_NAME) pg_dump -U $(DB_USER) -h $(DB_HOST) -d $(DB_NAME) > ./postgres/create_tables.sql