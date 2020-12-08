package datastore

import (
	"database/sql"
	"fmt"

	// pq is the postgres driver for the sql package
	_ "github.com/lib/pq"
)

// ConnectionDetails represents the pieces needed to open a connection
// to Postgres, including the host, port, user, password and dbname
type ConnectionDetails struct {
	host     string
	port     int
	user     string
	password string
	dbname   string
}

// config is a given instance of the connection details needed for Postgres
var (
	config ConnectionDetails
	db     *sql.DB
	err    error
)

// ConnectPostgres establishes the connection between Deskmate and
// a local Postgres database. The connection details for Deskmate
// should be loaded from environment variables. If those variables
// aren't present, Deskmate should ask for those details at the command
// line. From there, the connection details are pulled together to open a
// connection with Postgres, and then ping the database to ensure the
// connection is active.
func ConnectPostgres() {
	config = ConnectionDetails{
		host:     "db",
		port:     5432,
		user:     "docker",
		password: "docker",
		dbname:   "postgres",
	}

	// Open a connection using connection details
	connection := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.host, config.port, config.user, config.password, config.dbname)
	db, err = sql.Open("postgres", connection)
	if err != nil {
		fmt.Println("Error opening connection to Postgres database:", err.Error())
		// TODO: Add logging for an error on database connection
	}

	// Ping database to ensure connection
	err = db.Ping()
	if err != nil {

		fmt.Println("Error pinging Postgres database:", err.Error())
		// TODO: Add logging for a failed ping, likely due to a connection issue
	}
	checkTable()
}
