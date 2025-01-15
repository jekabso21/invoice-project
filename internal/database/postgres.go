package database

import (
	"database/sql"
	"log"
)

var PostgresDB *sql.DB

func ConnectPostgres() {
	log.Println("Connected to PostgreSQL successfully!")
}

func ExecuteInsert(query string, args ...interface{}) (int64, error) {
	result, err := PostgresDB.Exec(query, args...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func ExecuteSelect(query string, args []interface{}, dest ...interface{}) error {
	return PostgresDB.QueryRow(query, args...).Scan(dest...)
}

func ExecuteUpdate(query string, args ...interface{}) (int64, error) {
	result, err := PostgresDB.Exec(query, args...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func ExecuteDelete(query string, args ...interface{}) (int64, error) {
	result, err := PostgresDB.Exec(query, args...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}
