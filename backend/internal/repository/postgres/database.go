package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"subscription/internal/config"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

func Connect(cfg *config.Config) *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=%s",
		cfg.Db.Host, cfg.Db.Port, cfg.Db.User, cfg.Db.Password,
		cfg.Db.Name, cfg.Db.Sslmode)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}

	if err := goose.Up(db, "migrations"); err != nil {
		panic(err)
	}

	return db
}
