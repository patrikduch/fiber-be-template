package database

import (
    "database/sql"
    "log"
    "os"
    
    "fiber-be-template/ent"  // Import your Ent package
    
    _ "github.com/lib/pq"
)

// Your existing SQL DB connection
var DB *sql.DB

// Add a new variable for the Ent client
var EntClient *ent.Client

// Your existing Init function
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
    
    // Initialize Ent client with the same DSN
    EntClient, err = ent.Open("postgres", dsn)
    if err != nil {
        log.Fatalf("Failed to create Ent client: %v", err)
    }
    
    // Optional: Run migrations (remove if you prefer to manage migrations separately)
    // if err := EntClient.Schema.Create(context.Background()); err != nil {
    //     log.Fatalf("Failed to create schema: %v", err)
    // }
    
    log.Println("Ent client initialized")
}

// Add a cleanup function to close connections
func Close() {
    if DB != nil {
        DB.Close()
    }
    
    if EntClient != nil {
        EntClient.Close()
    }
}