setup:
	mkdir db_data; \
	mkdir db_data/postgres; \
	cd cmd; \
	openssl genrsa -out private_key.pem 4096

start:
	docker-compose up -d

migrate_create:
	migrate create -seq -ext sql -dir ./cmd/internal/database/migrate/migrations ${name}

migrate_up:
	migrate -path=./cmd/internal/database/migrate/migrations -database "postgresql://admin:secret@localhost:5432/social_db?sslmode=disable" -verbose up

migrate_down1:
	migrate -path=./cmd/internal/database/migrate/migrations -database "postgresql://admin:secret@localhost:5432/social_db?sslmode=disable" -verbose down 1