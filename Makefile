MIGRATE_DSN ?= root:password@tcp(127.0.0.1:3306)/testdb?multiStatements=true
MIGRATE_DIR ?= ./migrations
MIGRATE_BIN ?= go run ./cmd/migrate/main.go

.PHONY: migrate-up migrate-down migrate-create test

migrate-up:
	$(MIGRATE_BIN) --cmd up --dir $(MIGRATE_DIR) --dsn $(MIGRATE_DSN)

migrate-down:
	$(MIGRATE_BIN) --cmd down --dir $(MIGRATE_DIR) --dsn $(MIGRATE_DSN)

migrate-create:
	@read -p "Enter migration name: " name; \
	timestamp=$$(date +%Y%m%d%H%M%S); \
	touch $(MIGRATE_DIR)/$${timestamp}_$${name}.up.sql; \
	touch $(MIGRATE_DIR)/$${timestamp}_$${name}.down.sql; \
	echo "Created migration $${timestamp}_$${name}"

test:
	go test ./... -v
