package postgres

import (
	"database/sql"
	"errors"

	"github.com/doug-martin/goqu/v9"
)

type GoquDB interface {
	From(cols ...interface{}) *goqu.SelectDataset
	Insert(table interface{}) *goqu.InsertDataset
	Update(table interface{}) *goqu.UpdateDataset
	Delete(table interface{}) *goqu.DeleteDataset
	ScanStruct(i interface{}, query string, args ...interface{}) (bool, error)
	ScanStructs(i interface{}, query string, args ...interface{}) error
}

func GoquFromQ(q Queryable) (GoquDB, error) {
	switch t := q.(type) {
	case *sql.Tx:
		return goqu.NewTx("postgres", t), nil
	case *sql.DB:
		return goqu.New("postgres", t), nil
	default:
		return nil, errors.New("unsupported type of queryable")
	}
}
