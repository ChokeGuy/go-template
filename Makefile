# Makefile
ENV := $(CURDIR)/.env

# Environment variables for project
include $(ENV)

server:
	go run cmd/main.go

sqlc:
	sqlc generate

migrateup:
	migrate -path db/migrations -database "$(POSTGRES_URL)" -verbose up

migrateup1:
	migrate -path db/migrations -database "$(POSTGRES_URL)" -verbose up 1

migratedown:
	migrate -path db/migrations -database "$(POSTGRES_URL)" -verbose down

migratedown1:
	migrate -path db/migrations -database "$(POSTGRES_URL)" -verbose down 1

migrateforce:
	migrate -path db/migrations -database "$(POSTGRES_URL)" -verbose force 1

.PHONY: server sqlc migrateup migrateup1 migratedown migratedown1 migrateforce

