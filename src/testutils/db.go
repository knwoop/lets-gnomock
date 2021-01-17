package testutils

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"runtime"

	mysqld "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/orlangure/gnomock"
	"github.com/orlangure/gnomock/preset/mysql"
)

type TestDB struct {
	DB        *sqlx.DB
	container *gnomock.Container
}

func NewDBTestClient() (*TestDB, error) {
	b, err := getInitStmts()
	if err != nil {
		panic(err)
	}

	p := mysql.Preset(
		mysql.WithUser("test", "pass"),
		mysql.WithDatabase("test_database"),
		mysql.WithQueries(string(b)),
	)

	container, err := gnomock.Start(p)

	addr := container.DefaultAddress()
	c := mysqld.Config{
		DBName:               "test_database",
		User:                 "test",
		Passwd:               "pass",
		Addr:                 addr,
		Net:                  "tcp",
		ParseTime:            true,
		AllowNativePasswords: true,
	}
	dsn := c.FormatDSN()
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open db master(%s): %w", dsn, err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping db: %w", err)
	}
	return &TestDB{
		DB:        db,
		container: container,
	}, nil
}

func (t *TestDB) Close() {
	_ = gnomock.Stop(t.container)
}

func getInitStmts() ([]byte, error) {
	pathName := "../../db/schema.sql"

	b, err := ioutil.ReadFile(pathName)
	if err != nil {
		return nil, fmt.Errorf("cannot open schema.sql: %w", err)
	}

	return b, nil
}
