package db

import (
	"database/sql"
	"fmt"
)

func CreateDatabase() (*sql.DB, error) {
	servername := "localhost:3306"
	user := "root"
	password := "aardvark"
	dbName := "churrios"

	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true&multiStatements=true", user, password, serverName, dbName)
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}

	if err := migrateDatabase(db); err != nil {
		return db, err
	}

	return db, nil

}
