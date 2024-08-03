package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

var conn *sql.DB

func InitConn() {
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRESO_PORT")
	user := os.Getenv("POSTGRES_USER")
	pass := os.Getenv("POSTGRES_PASSWORD")
	database := os.Getenv("POSTGRES_DATABASE")

	postgresConnectionString := fmt.Sprintf("host=%s port=%s user=%s "+"password=%s dbname=%s sslmode=disable TimeZone=%s", host, port, user, pass, database, "Asia/Tashkent")

	db, err := sql.Open("postgres", postgresConnectionString)
	if err != nil {
		log.Fatalf("can not connect postgres: %v", err)
	}

	db.SetMaxOpenConns(4096)
	db.SetMaxIdleConns(256)
	db.SetConnMaxIdleTime(2 * time.Minute)
	db.SetConnMaxLifetime(2 * time.Minute)

	// Ping the database to check if the connection is successful
	if err := db.Ping(); err != nil {
		log.Fatalf("error_ping_db: %v", err)
	}

	conn = db
}

func Close() {
	err := conn.Close()
	if err != nil {
		log.Printf("error_on_close_postgres_conn: %v", err)
	}
}
