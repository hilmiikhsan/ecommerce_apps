package database

import (
	"fmt"
	"time"

	"github.com/ecommerce/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func ConnectSQLXPostgres(cfg config.DB) (db *sqlx.DB, err error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name,
	)

	db, err = sqlx.Open("postgres", dsn)
	if err != nil {
		return
	}

	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetConnMaxLifetime(time.Duration(cfg.MaxConnLifetime) * time.Second)

	if err = db.Ping(); err != nil {
		return
	}

	return
}
