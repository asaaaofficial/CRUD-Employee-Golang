package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// InitDatabase initializes and returns a database connection
func InitDatabase() (*sql.DB, error) {
	dsn := "root:@tcp(localhost:3308)/crud_employee?parseTime=true"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Printf("Database connection failed: %v", err)
		return nil, err
	}

	// Cek koneksi
	if err := db.Ping(); err != nil {
		log.Printf("Database ping failed: %v", err)
		return nil, err
	}

	// Atur koneksi pool (opsional)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)

	log.Println("Database connected successfully")
	return db, nil
}
