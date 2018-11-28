package pgx

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

var conn *sql.DB

func GetConnection() *sql.DB {
	if conn != nil {
		return conn
	}

	url := os.Getenv("POSTGRESQL_URL")
	if url == "" {
		url = "postgres://postgres@127.0.0.1/sqlx?sslmode=disable"
	}
	conn, err := sql.Open("postgres", url)
	if err != nil {
		panic(err)
	}

	return conn
}

func Exec(query string, args ...interface{}) (sql.Result, error) {
	return GetConnection().Exec(query, args...)
}

func QueryRow(query string, args ...interface{}) *sql.Row {
	return GetConnection().QueryRow(query, args...)
}

func init() {
	conn := GetConnection()

	conn.Exec(`DROP TABLE IF EXISTS test_nullduration`)
	conn.Exec(`DROP TABLE IF EXISTS test_duration`)
	conn.Exec(`CREATE TABLE test_nullduration(id serial PRIMARY KEY, field interval)`)
	conn.Exec(`CREATE TABLE test_duration(id serial PRIMARY KEY, field interval NOT NULL)`)
}
