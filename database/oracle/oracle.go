package oracle

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/sijms/go-ora/v2"
)

type OracleDatabase struct {
	connectionString string
}

func (db *OracleDatabase) DB() (*sql.DB, error) {
	connStr := db.connectionString
	conn, err := sql.Open("oracle", connStr)
	if err != nil {
		log.Println(fmt.Sprintf("error in opening connection: %v", err.Error()))
		return nil, err
	}
	err = conn.Ping()
	if err != nil {
		log.Println(fmt.Sprintf("error in pinging: %v", err.Error()))
		return nil, err
	}

	return conn, nil
}

func (db *OracleDatabase) ConnectionString(params map[string]any) error {
	ip, ok := params["ip"].(string)
	if !ok {
		errorString := "ip is required and must be a string"
		log.Println(errorString)
		return errors.New(errorString)
	}
	port, ok := params["port"].(int)
	if !ok {
		errorString := "port is required and must be a number"
		log.Println(errorString)
		return errors.New(errorString)
	}
	service, ok := params["service"].(string)
	if !ok {
		errorString := "service is requried and must be as string"
		log.Println(errorString)
		return errors.New(errorString)
	}
	user, ok := params["user"].(string)
	if !ok {
		errorString := "user is required and must be be a string"
		log.Println(errorString)
		return errors.New(errorString)
	}
	password, ok := params["password"].(string)
	if !ok {
		errorString := "password is required and must be a string"
		log.Println(errorString)
		return errors.New(errorString)
	}
	options, ok := params["options"].(map[string]string)
	if !ok {
		options = nil
	}
	connStr := go_ora.BuildUrl(ip, port, service, user, password, options)
	db.connectionString = connStr
	return nil
}
