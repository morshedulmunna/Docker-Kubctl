air_install: go install github.com/air-verse/air@latest


migrate: go get -u github.com/golang-migrate/migrate/v4

create_migrate_file : migrate create -ext sql -dir migrations -seq init