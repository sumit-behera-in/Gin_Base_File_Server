package db

import (
	"github.com/jmoiron/sqlx"
)

type Client struct {
	db *sqlx.DB
}

type UserInfo struct {
	Uid       int64
	Email     string
	UserName  string
	ClusterId int64
}

type ClusterInfo struct {
	ClusterName  string  `db:"name"`
	ClusterURL   string  `db:"url"`
	UserCount    int16   `db:"user_count"`
	TotalStorage float64 `db:"total_storage"` // in GB
	TotalUsed    float64 `db:"total_used"`    // in GB
}
