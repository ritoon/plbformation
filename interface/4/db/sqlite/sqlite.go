package sqlite

import (
	"plbformation/interface/4/db"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func NewSQLite(name string) *db.DBStore {
	_, err := gorm.Open(sqlite.Open(name), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return &db.DBStore{
		// User:    &SQLiteUser{db: dbconn},
		// Product: &SQLiteProduct{db: dbconn},
	}
}

// Implementation of UserStore interface
// Implementation of ProductStore interface
