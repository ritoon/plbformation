package postgres

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	DB      *gorm.DB
	User    *PostgresUser
	Product *PostgresProduct
}

func NewPostgres(name string) *Postgres {
	dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return &Postgres{DB: db}
}
