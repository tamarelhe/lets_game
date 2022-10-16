postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=lg -e POSTGRES_PASSWORD=lg2022 -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=lg --owner=lg lets_game

dropdb:
	docker exec -it postgres12 dropdb lets_game

migrateup:
	migrate -path db/migration -database "postgresql://lg:lg2022@localhost:5432/lets_game?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://lg:lg2022@localhost:5432/lets_game?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go clean -testcache
	go test -v -cover ./...

server:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server