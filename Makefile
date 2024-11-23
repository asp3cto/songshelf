# Load environment variables from dev.env file
#ifneq (,$(wildcard dev.env))
#    include dev.env
#    export
#else
#    $(error dev.env file not found)
#endif

LOCAL_BIN:=$(CURDIR)/bin

# Install golangci-lint
install-golangci-lint:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.61.0

# Lint the Go code
lint:
	$(LOCAL_BIN)/golangci-lint run ./... --config .golangci.pipeline.yaml

include dev.env
export

initdb:
	PGPASSWORD=$(PG_PASSWORD) psql -U $(PG_USER) -p $(PG_PORT) -d postgres -c "DROP DATABASE IF EXISTS $(PG_DBNAME)"
	PGPASSWORD=$(PG_PASSWORD) psql -U $(PG_USER) -p $(PG_PORT) -d postgres -c "CREATE DATABASE $(PG_DBNAME)"
	# Run goose migrations
	goose -dir sql/migrations postgres $(PG_DSN) up
	PGPASSWORD=$(PG_PASSWORD) psql  -U $(PG_USER) -p $(PG_PORT) -d $(PG_DBNAME) -f sql/fixtures/001_init.sql

# Generate SQLC code
sqlc:
	rm -r internal/database
	mkdir internal/database
	sqlc generate
