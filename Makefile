setup:
	mkdir db_data; \
	mkdir db_data/postgres; \
	cd cmd; \
	openssl genrsa -out private_key.pem 4096

start:
	docker-compose up -d