package db

import (
	"database/sql"
	"os"

	"github.com/glugate/uno/pkg/uno/db/query"
	"github.com/glugate/uno/pkg/uno/log"
)

// DB
type DB struct {
	StdDB   *sql.DB
	Builder *query.Builder
}

// NewDB
func NewDB() *DB {
	driver, _ := os.LookupEnv("DB_DRIVER")
	dsn, _ := os.LookupEnv("DB_DSN")

	log.Default().Debug("Connecting to DB: %s, %s", driver, dsn)

	// Create default stantard sql connection
	stdDB := Open(driver, dsn)

	// Create adapter with existing db sql conn
	builder := query.NewBuilder()

	// Ping
	err := stdDB.Ping()
	if err != nil {
		panic(err)
	}

	Migrate(stdDB)
	return &DB{
		StdDB:   stdDB,
		Builder: builder,
	}
}

// Open
func Open(driver string, dsn string) *sql.DB {
	log.Default().Debug("%s: connecting to the database...", dsn)
	conn, err := sql.Open(driver, dsn)
	if err != nil {
		panic(err)
	}
	return conn
}
