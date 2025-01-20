package db

import "github.com/jmoiron/sqlx"

type Client struct {
	db *sqlx.DB
}
