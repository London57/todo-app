toml=config/dev.toml
DB_USER := $(shell grep -E 'user\s*=' $(toml) | cut -d '"' -f 2)
DB_PASS := $(shell grep -E 'password\s*=' $(toml) | cut -d '"' -f 2)
DB_NAME := $(shell grep -E 'database\s*=' $(toml) | cut -d '"' -f 2)
SSLMODE := $(shell grep -E 'sslmode\s*=' $(toml) | cut -d '"' -f 2)

run-postgres:	
	docker run --rm --name postgres \
		-e POSTGRES_USER=$(DB_USER) \
		-e POSTGRES_PASSWORD=$(DB_PASS) \
		-e POSTGRES_DB=$(DB_NAME) \
		-e SSLMODE=${SSLMODE} \
		-p 5432:5432 \
		-d postgres

.PHONY: run-postgres