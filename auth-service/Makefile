CONFIG_FILE=config/config.yml

PG_DSN=$(shell yq e '.pg_dsn' $(CONFIG_FILE))

.PHONY: migration migrateup migrateup1 migratedown migratedown1

migration:
	migrate create --ext sql --dir migrations -seq $(MIGRATION_FILE_NAME)

migrateup:
	migrate -path migrations -database $(PG_DSN) -verbose up

migrateup1:
	migrate -path migrations -database $(PG_DSN) -verbose up 1

migratedown:
	migrate -path migrations -database $(PG_DSN) -verbose down

migratedown1:
	migrate -path migrations -database $(PG_DSN) -verbose down 1