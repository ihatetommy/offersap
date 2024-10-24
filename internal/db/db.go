package db

import (
	"database/sql"
	"log"

	"OffersApp/config"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	cfg := config.LoadConfig()

	connStr := "user=" + cfg.DBUser +
		" dbname=" + cfg.DBName +
		" password=" + cfg.DBPassword +
		" host=" + cfg.DBHost +
		" port=" + cfg.DBPort +
		" sslmode=require"
	
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal(err)
	}
}
