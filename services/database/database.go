package database

import (
	"fmt"
	"pharos/services/conf"

	"github.com/go-pg/pg/v10"
)

func NewDBOptions(cfg conf.Config) *pg.Options {
	db_con := &pg.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.DbHost, cfg.DbPort),
		Database: cfg.DbName,
		User:     cfg.DbUser,
		Password: cfg.DbPassword,
	}
	// fmt.Printf("\nDB connection - %v\n", db_con)
	return db_con
}
