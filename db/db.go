package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func New(dsn string) (*Client, error) {
	var c Client
	var err error
	if c.db, err = sqlx.Connect("postgres", dsn); err != nil {
		return nil, err
	}
	return &c, nil
}

