package config

import (
	"database/sql"
	"fmt"
	"golang-crud/helpers"

	"github.com/rs/zerolog/log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbName   = "test"
)

func DatabaseConnection() *sql.DB {
	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

	db, err := sql.Open("postgres", sqlInfo)
	helpers.PanicIfError(err)

	err = db.Ping()
	helpers.PanicIfError(err)

	log.Info().Msg("Connected to database!")

	return db
}
