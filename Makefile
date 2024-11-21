setup:
	mkdir db_data; \
	mkdir db_data/postgres; \
	cd cmd; \
	openssl genrsa -out private_key.pem 4096

start:
	docker-compose up -d

migrate_create:
	migrate create -seq -ext sql -dir ./cmd/internal/database/migrate/migrations ${name}