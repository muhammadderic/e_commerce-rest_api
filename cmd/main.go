package main

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/muhammadderic/ecomrest/cmd/api"
	"github.com/muhammadderic/ecomrest/cmd/db"
)

func main() {
	db, err := db.NewMySQLStorage(mysql.Config{
		User:                 "root",
		Passwd:               "yourpassword",
		Addr:                 "127.0.0.1:3306",
		DBName:               "yourdb",
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)

	server := api.NewAPIServer(":8080", nil)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB connection established successfully!")
}
