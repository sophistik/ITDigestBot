package postgres

import (
	"database/sql"
)

type Queryable interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	Prepare(query string) (*sql.Stmt, error)
}

func DBOrTx(db *sql.DB, tx *sql.Tx) Queryable {
	var q Queryable

	if db != nil {
		q = db
	} else {
		q = tx
	}

	return q
}
