

postgres:
	docker run --name postgres12 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -p 5432:5432 -d postgres:12-alpine
postgres-start:
	docker start postgres12
create-db:
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank
drop-db:
	docker exec -it postgres12 dropdb simple_bank
migrate-up:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up
migrate-up1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up 1
migrate-down:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down
migrate-down1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down 1
test:
	go test -v -cover ./...
sqlc:
	docker run --rm -v $(PWD):/src -w /src kjconroy/sqlc generate
	make mocks
server:
	make postgres-start
	go run main.go
mocks:
	mockgen -package mockDB -destination db/mock/store.go github.com/NicolasMartino/simplebank/db/sqlc Store
	
.PHONY: createdb postgres dropdb migrate-up migrate-down migrate-up1 migrate-down1 sqlc test server start-postgres mocks