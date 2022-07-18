package pgclient

import "github.com/jackc/pgx/v4"

type ClientInterface interface {
	Connect(connectionString string) error
	Close()
	Ping() error
	Exec(sql string, arguments ...interface{}) error
	QueryRow(sql string, args ...interface{}) pgx.Row
	CopyFromRows(table pgx.Identifier, columns []string, rows [][]interface{}) (int64, error)
}
