.PHONY: sqlc createdb dropdb postgres  migrateup migratedown migrateup1 migratedown1 migratedrop server

sqlc:
	@sqlc generate
# Start PostgreSQL container
postgres:
	@docker run --name postgresDB -e POSTGRES_USER=root -e POSTGRES_PASSWORD=12345678 -p 5432:5432 -d postgres:latest

# Create simple_bank database
createdb:
	@docker exec -it postgresDB createdb -U root -O root simple_bank

# Drop simple_bank database
dropdb:
	@docker exec -it postgresDB dropdb simple_bank

# Run all migrations (up)
migrateup:
	@ migrate -path ./db/migrate -database "postgres://root:12345678@localhost:5432/simple_bank?sslmode=disable" up

# Rollback last migration 
migratedown:
	@ migrate -path ./db/migrate -database "postgres://root:12345678@localhost:5432/simple_bank?sslmode=disable" down 


# Run all migrations 1
migrateup1:
	@ migrate -path ./db/migrate -database "postgres://root:12345678@localhost:5432/simple_bank?sslmode=disable" up 1

# Rollback last migration 1
migratedown1:
	@ migrate -path ./db/migrate -database "postgres://root:12345678@localhost:5432/simple_bank?sslmode=disable" down 1 


# Drop all tables in the database
migratedrop:
	@ migrate -path ./db/migrate -database "postgres://root:12345678@localhost:5432/simple_bank?sslmode=disable" drop -f

server:
	@go run main.go