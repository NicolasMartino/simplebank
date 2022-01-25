

postgres:
	docker run --name postgres12 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -p 5432:5432 -d postgres:12-alpine
create-db:
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank
drop-db:
	docker exec -it postgres12 dropdb simple_bank
migrate-up:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up
migrate-down:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down
test:
	go test -v -cover ./...
sqlc:
	docker run --rm -v $(PWD):/src -w /src kjconroy/sqlc generate
	
.PHONY: createdb postgres dropdb migrate-up migrate-down sqlc test