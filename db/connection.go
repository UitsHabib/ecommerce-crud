package db

import (
	"fmt"

	"github.com/UitsHabib/ecommerce-crud/config"
	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

func Connect(cnf *config.DB) (*sqlx.DB, error) {
	// establish a connection to the PostgreSQL database
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cnf.Host, cnf.Port, cnf.User, cnf.Password, cnf.DBName)
	return sqlx.Connect("postgres", connectionString)
}
