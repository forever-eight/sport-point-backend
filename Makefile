.PHONY: migrate
migrate:
	migrate -path ./schema -database "postgresql://db-user:db-pass@localhost:5432/db-name?sslmode=disable" -verbose up

.PHONY: migrate-down
migrate-down:
	migrate -path ./schema -database "postgresql://db-user:db-pass@localhost:5432/db-name?sslmode=disable" -verbose drop