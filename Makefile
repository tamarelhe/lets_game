DB_URL=postgresql://lg:lg2022@localhost:5432/lets_game?sslmode=disable

network:
	docker network create lets-game-network

postgres:
	docker run --name postgres12 --network lets-game-network -p 5432:5432 -e POSTGRES_USER=lg -e POSTGRES_PASSWORD=lg2022 -d postgres:14-alpine

createdb:
	docker exec -it postgres12 createdb --username=lg --owner=lg lets_game

dropdb:
	docker exec -it postgres12 dropdb lets_game

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

migrateup1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

migratedown1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1

sqlc:
	sqlc generate

test:
	go clean -testcache
	go test -v -cover ./...

server:
	go run main.go

.PHONY: network postgres createdb dropdb migrateup migratedown migrateup1 migratedown1 sqlc test server