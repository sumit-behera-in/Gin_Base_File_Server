package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func New(dsn string) (*Client, error) {
	var c Client
	var err error

	// connect to db
	if c.db, err = sqlx.Connect("postgres", dsn); err != nil {
		return nil, err
	}

	// create all the tables we need
	if err := c.init(); err != nil {
		return nil, err
	}

	return &c, nil
}

// initialize all tables
func (c *Client) init() error {
	if err := c.initClusters(); err != nil {
		return err
	}

	return nil
}

func (c *Client) Close() error {
	return c.db.Close()
}
