# Test Interview KC

A Go (Golang) project for interview/testing purposes, featuring:
- Clean architecture
- GORM for database access
- Fiber web framework
- Zap logging
- Table-driven unit tests
- GoMock for mocking
- Golang Migrate for DB migrations
- Docker Compose for easy MySQL setup

## Prerequisites
- [Go](https://golang.org/dl/) 1.20+
- [Docker](https://www.docker.com/get-started)
- [Make](https://www.gnu.org/software/make/)

## Quick Start

### 1. Clone the repository
```sh
git clone https://github.com/ahmadfarras/test-interview-kc.git
cd test-interview-kc
```

### 2. Start MySQL with Docker Compose
This will create a MySQL 8.0 container with the `testdb` database and auto-create tables from `migrations/init.sql`.

```sh
docker compose up -d
```

- MySQL will be available at `localhost:3306` (user: `root`, password: `password`, db: `testdb`).
- To reset the DB (drop all data and recreate tables):
  ```sh
  docker compose down -v
  docker compose up -d
  ```

### 3. Run the API
```sh
go run ./cmd/api/main.go
```

### 4. Run Tests
```sh
make test
# or
# go test ./...
```

## Migrations
- Create a new migration:
  ```sh
  make migrate-create
  ```
- Edit the generated `.up.sql` and `.down.sql` files in `migrations/`.
- Apply migrations:
  ```sh
  make migrate-up
  ```
- Rollback migrations:
  ```sh
  make migrate-down
  ```
