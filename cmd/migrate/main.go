package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
)

func main() {
	dir := flag.String("dir", "./migrations", "directory with migration files")
	dsn := flag.String("dsn", os.Getenv("MIGRATE_DSN"), "database DSN")
	cmd := flag.String("cmd", "up", "migration command: up, down")
	flag.Parse()

	if *dsn == "" {
		log.Fatal("MIGRATE_DSN environment variable or --dsn flag required")
	}

	m, err := migrate.New(
		fmt.Sprintf("file://%s", *dir),
		fmt.Sprintf("mysql://%s", *dsn),
	)
	if err != nil {
		log.Fatalf("failed to create migrate instance: %v", err)
	}

	switch *cmd {
	case "up":
		err = m.Up()
	case "down":
		err = m.Down()
	default:
		log.Fatalf("unknown command: %s", *cmd)
	}

	if err != nil && err != migrate.ErrNoChange {
		log.Fatalf("migration failed: %v", err)
	}

	log.Printf("migration %s successful", *cmd)
}
