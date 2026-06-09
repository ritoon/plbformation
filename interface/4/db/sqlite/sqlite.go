package sqlite

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type SQLite struct {
	DB *gorm.DB
}

func NewSQLite(name string) *SQLite {
	db, err := gorm.Open(sqlite.Open(name), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return &SQLite{DB: db}
}

// Implementation of UserStore interface
// Implementation of ProductStore interface
