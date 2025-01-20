package db

import (
	"testing"

	_ "github.com/lib/pq"
)

var (
	dsn = "postgresql://postgres:admin@localhost:5432/gin_http_register?sslmode=disable"
)

func TestNew(t *testing.T) {

	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "connect to localHost",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(dsn)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != nil && got.db.Close() != nil {
				t.Error("Failed to close db")
				return
			}

		})
	}
}
