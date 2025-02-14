MIGRATION_PATH = ./cmd/migrate/migrations


.PHONY: migrate-create
migration:
	@migrate create -seq -ext sql -dir ${MIGRATION_PATH} $(filter-out $@,${MAKECMDGOALS})