/*
 * This is where all database activity should happen
 * each call to the database unless two are very similar
 * should have it's own function and be called here
 * Data should then be returned
 */
package database

import (
	"database/sql"

	"github.com/cfrank/tny.al/constants"
	_ "github.com/go-sql-driver/mysql"
)

/*
 * Holds the db object as well as a variable telling if
 * Ping() was successful
 */
type DB struct {
	Db    *sql.DB
	Alive bool
}

// Global DB struct
var MyDb DB = DB{}

func OpenDatabase() error {
	db, sqlOpenError := sql.Open("mysql", constants.DB_HOST)

	if sqlOpenError != nil {
		return sqlOpenError
	}

	MyDb.Db = db

	pingError := db.Ping()

	if pingError != nil {
		MyDb.Alive = false
	} else {
		MyDb.Alive = true
	}

	return nil
}

func CloseDatabase() {
	MyDb.Db.Close()
}

/*
 * Check to make sure a userid is available
 *
 * This needs to be done because when a user visits the homepage
 * they are automatically given an id, and they must be unique
 */
func UniqueIdCheck(id string) bool {
	var userid string
	err := MyDb.Db.QueryRow("Select userid from user where userid =?", id).Scan(&userid)
	switch {
	case err == sql.ErrNoRows:
		return true
	case err != nil:
		return false
	default:
		return false
	}
}
