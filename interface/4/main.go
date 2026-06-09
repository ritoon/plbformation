package main

import (
	"log"
	"os"

	"plbformation/interface/4/db"
	"plbformation/interface/4/db/postgres"
	"plbformation/interface/4/db/sqlite"
	"plbformation/interface/4/model"
)

const (
	EnvProd = "prod"
	EnvDev  = "dev"
)

func init() {
	// Initialiser la base de données en fonction de l'environnement
	// si l'environnement est prod, initialiser une base de données SQLite
	// sinon initialiser une base de données en mémoire
}

func main() {
	// Créer un utilisateur et un produit
	// Enregistrer l'utilisateur et le produit dans la base de données
	var db *db.DBStore

	env := os.Getenv("ENV")

	if env == EnvProd {
		// Initialiser la base de données PostgreSQL
		db = postgres.NewPostgres("dsn infos")
	} else {
		// Initialiser la base de données SQLite
		db = sqlite.NewSQLite("test.db")
	}

	var u model.User
	u.Name = "John"
	u.Age = 30

	if err := db.User.Create(&u); err != nil {
		log.Fatalf("failed to create user: %v", err)
	}

}
