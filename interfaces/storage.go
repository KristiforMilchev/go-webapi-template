package interfaces

import "database/sql"

type Storage interface {
	Open(connection string) bool
	Close() bool
	Single(sql string, params []interface{}) *sql.Row
	Where(sql string, params []interface{}) *sql.Rows
}
