package db

import (
	"database/sql"
	"fmt"
	"log"
)

type Database interface {
	DB() (*sql.DB, error)
	SetConnectionString(map[string]any) error
}

func Exec(db Database, query string, args []any) error {
	conn, err := db.DB()
	if err != nil {
		log.Println(fmt.Sprintf("error in creating db: %v", err.Error()))
		return err
	}
	transaction, err := conn.Begin()
	if err != nil {
		log.Println(fmt.Sprintf("error in begining transaction: %v", err.Error()))
		return err
	}
	_, err = transaction.Exec(query, args...)
	if err != nil {
		log.Println(fmt.Sprintf("error in executing transaction: %v", err.Error()))
		transaction.Rollback()
		return err
	}
	transaction.Commit()
	return nil
}
