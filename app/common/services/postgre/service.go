package postgre

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"geektest/internal/logs"
)

func getConnection() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, pass, dbName)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	return db	
}

func selectData(query string) (*sql.Rows, error) {
	db := getConnection()
	defer db.Close()
	rows, err := db.Query(query)
	if err != nil {
		logs.Error("failed to run postgre select query", err)
		return nil, err
	}

	return rows, nil
}