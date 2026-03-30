package database

import (
	"context"
	"database/sql"
	"log"
	"os"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"fiber-be-template/ent"
	_ "github.com/lib/pq"
)

var DB *sql.DB
var EntClient *ent.Client

func Init() {
	dsn := os.Getenv("DATABASE_URL")
	var err error
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatalf("Failed to ping DB: %v", err)
	}

	log.Println("Database connection established")

	drv := entsql.OpenDB(dialect.Postgres, DB)
	EntClient = ent.NewClient(ent.Driver(drv)).Debug()

	log.Println("Running schema migration...")
	if err := EntClient.Schema.Create(context.Background()); err != nil {
		log.Fatalf("Schema migration failed: %v", err)
	}
	log.Println("Schema migration done ✓")

	Seed()
	log.Println("Seed done ✓")
}

func Close() {
	if DB != nil {
		if err := DB.Close(); err != nil {
			log.Printf("Error closing DB connection: %v", err)
		}
	}
	if EntClient != nil {
		if err := EntClient.Close(); err != nil {
			log.Printf("Error closing Ent client: %v", err)
		}
	}
}
