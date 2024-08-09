package db

import (
	"database/sql"
	"log"
	"lottery-api/internal/config"

	_ "github.com/lib/pq"
)

var Conn *sql.DB

func Connect() {
	var err error
	dsn := config.GetDSN()
	Conn, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	// if err := DB.Ping(); err != nil {
	// 	log.Fatalf("Failed to ping the database: %v", err)
	// }
}
