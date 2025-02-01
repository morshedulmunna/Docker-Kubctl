air_install: go install github.com/air-verse/air@latest

goose: go install github.com/pressly/goose/v3/cmd/goose@latest

migrate: go get -u github.com/golang-migrate/migrate/v4

create_migrate_file : migrate create -ext sql -dir migrations -seq init

postgres:
	docker run --name postgres16 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:16

install_sqlc:
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

sqlc:
	sqlc generate