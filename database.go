package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDatabase() error {
	var err error
	DB, err = sql.Open("mysql", "root:@tcp(localhost:3306)/crud-mahasiswa?parseTime=true")
	if err != nil {
		return fmt.Errorf("gagal koneksi ke database: %v", err)
	}

	err = DB.Ping()
	if err != nil {
		return fmt.Errorf("database tidak bisa diakses: %v", err)
	}

	return nil
}

func CloseDatabase() {
	if DB != nil {
		DB.Close()
	}
}
