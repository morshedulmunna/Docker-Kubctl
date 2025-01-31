air_install: go install github.com/air-verse/air@latest

goose: go install github.com/pressly/goose/v3/cmd/goose@latest

migrate: go get -u github.com/golang-migrate/migrate/v4

create_migrate_file : migrate create -ext sql -dir migrations -seq init