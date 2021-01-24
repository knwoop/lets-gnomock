package db

import (
	"fmt"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"github.com/knwoop/lets-gnomock/src/config"
)

func NewClient(env *config.Env) (*sqlx.DB, error) {
	c := mysql.Config{
		DBName:               env.MysqlDatabase,
		User:                 env.MysqlUser,
		Passwd:               env.MysqlPass,
		Addr:                 env.MysqlHost,
		Net:                  "tcp",
		ParseTime:            true,
		AllowNativePasswords: true,
	}
	dsn := c.FormatDSN()
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open db(%s): %w", dsn, err)
	}
	return db, nil
}
