package postgres

import (
	"database/sql"
	"fmt"
	"subscription/internal/config"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"go.uber.org/zap"
)

func Connect(cfg *config.Config, logger *zap.Logger) *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=%s",
		cfg.Db.Host, cfg.Db.Port, cfg.Db.User, cfg.Db.Password,
		cfg.Db.Name, cfg.Db.Sslmode)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		logger.Error("Cannot open db connection", zap.String("Error", err.Error()))
	}

	err = db.Ping()
	if err != nil {
		logger.Error("Cannot ping db connection", zap.String("Error", err.Error()))
	}

	err = goose.Up(db, "migrations")
	if err != nil {
		panic(err)
	}

	return db
}
