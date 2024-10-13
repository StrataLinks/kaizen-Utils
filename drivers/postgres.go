package drivers

import (
	"database/sql"

	_ "github.com/jackc/pgx/v4/stdlib"
)

// InitDatabase initializes and returns a database connection.
// It opens a connection to the database using the provided driver and dataSourceName.
// The function verifies the connection with a Ping.
func ProvidePostgreSQlDatabaseClient(driver, dataSourceName string) (*sql.DB, error) {
	return providerPostgreSQlDatabaseClient(driver, dataSourceName)
}

func providerPostgreSQlDatabaseClient(driver, dataSourceName string) (*sql.DB, error) {
	// Open the database connection with the provided driver and dataSourceName.
	// sql.Open doesn't actually establish any connections to the database,
	// it just prepares the database connection value. The actual connection to the
	// database will be established lazily, upon the first need for an interaction
	// with the database.
	db, err := sql.Open(driver, dataSourceName)
	if err != nil {
		// If there is an error opening the connection, return the error immediately.
		return nil, err
	}

	// Check the connection with Ping.
	// Ping verifies a connection to the database is still alive, establishing a connection if necessary.
	if err := db.Ping(); err != nil {
		// If Ping returns an error, close the database connection to avoid leaving open connections.
		db.Close() // It's safe to call Close on a nil pointer, but here db is not nil since sql.Open was successful.
		return nil, err
	}

	// Return the database connection handle.
	return db, nil
}
