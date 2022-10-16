postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=lg -e POSTGRES_PASSWORD=lg2022 -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=lg --owner=lg lets_game

dropdb:
	docker exec -it postgres12 dropdb lets_game

installmigrateubuntu:
	sudo curl -L https://packagecloud.io/golang-migrate/migrate/gpgkey | apt-key add -
	sudo echo "deb https://packagecloud.io/golang-migrate/migrate/ubuntu/ $(lsb_release -sc) main" > /etc/apt/sources.list.d/migrate.list
	sudo apt-get update
	sudo apt-get install -y migrate

migrateup:
	migrate -path db/migration -database "postgresql://lg:lg2022@localhost:5432/lets_game?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://lg:lg2022@localhost:5432/lets_game?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go clean -testcache
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test installmigrateubuntu