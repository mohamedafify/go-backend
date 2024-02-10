package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type Database interface {
	DB() (*sql.DB, error)
	ConnectionString(map[string]any) error
}

func Exec(c *gin.Context, db Database, query string, args []any) error {
	conn, err := db.DB()
	if err != nil {
		log.Println(fmt.Sprintf("error in creating db: %v", err.Error()))
		return err
	}
	txOptions := &sql.TxOptions{
		Isolation: sql.LevelDefault,
		ReadOnly:  false,
	}
	transaction, err := conn.BeginTx(c, txOptions)
	if err != nil {
		log.Println(fmt.Sprintf("error in begining transaction: %v", err.Error()))
		return err
	}
	_, err = transaction.ExecContext(c, query, args...)
	if err != nil {
		log.Println(fmt.Sprintf("error in executing transaction: %v", err.Error()))
		transaction.Rollback()
		return err
	}
	transaction.Commit()
	return nil
}
