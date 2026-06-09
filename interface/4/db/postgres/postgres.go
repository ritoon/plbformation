package postgres

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"plbformation/interface/4/db"
)

func NewPostgres(name string) *db.DBStore {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered in NewPostgres: %v", r)
		}
	}()

	dbconn, err := connectPostgres()
	if err != nil {
		panic(err)
	}

	var store db.DBStore
	store.User = &PostgresUser{db: dbconn}
	store.Product = &PostgresProduct{db: dbconn}

	return &store
}

func connectPostgres() (*gorm.DB, error) {
	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db, nil
}
