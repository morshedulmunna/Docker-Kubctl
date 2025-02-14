MIGRATION_PATH = ./cmd/migrate/migrations
DB_MIGRATOR_ADDR = postgres://postgres:postgres@localhost/pxomartdb?sslmode=disable

.PHONY: migrate-create
migration:
	@migrate create -seq -ext sql -dir ${MIGRATION_PATH} $(filter-out $@,${MAKECMDGOALS})


.PHONY: migrate-up
migrate-up:
	@migrate -path=${MIGRATION_PATH} -database=$(DB_MIGRATOR_ADDR) up

.PHONY: migrate-down
migrate-down:
	@migrate -path=${MIGRATION_PATH} -database=$(DB_MIGRATOR_ADDR) down $(filter-out $@,$(MAKECMDGOALS))