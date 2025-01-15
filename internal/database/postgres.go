package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var PostgresDB *sql.DB

// ConnectPostgres establishes a connection to the PostgreSQL database
func ConnectPostgres() {
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")

	// Format connection string
	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	var err error
	PostgresDB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Error connecting to PostgreSQL: %v\n", err)
	}

	// Verify connection
	if err = PostgresDB.Ping(); err != nil {
		log.Fatalf("PostgreSQL ping failed: %v\n", err)
	}
	log.Println("Connected to PostgreSQL successfully!")
}