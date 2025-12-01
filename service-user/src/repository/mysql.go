package repository

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/payments/service-user/src"
	"github.com/payments/service-user/src/config"
)

func dsn() string {
	username := config.C.GetString("database.sql_username")
	password := config.C.GetString("database.sql_password")
	hostname := config.C.GetString("database.sql_host")
	dbName := config.C.GetString("database.sql_dbname")
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
}

func OpenConnection() (*sql.DB, error) {
	log.Println(dsn())
	db, err := sql.Open("mysql", dsn())
	if err != nil {
		log.Printf("Error %s when opening DB", err)
		return nil, err
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(20)
	db.SetConnMaxLifetime(time.Minute * 5)

	log.Printf("Connected to DB %s successfully\n", config.C.GetString("database.sql_dbname"))

	return db, nil
}

func NewRepository(client *sql.DB) src.RepositoryList {
	return &Repository{
		db: client,
	}

}
