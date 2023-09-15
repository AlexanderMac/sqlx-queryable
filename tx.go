package queryable

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type QueryableTx struct {
	tx *sqlx.Tx
}

func (q *QueryableTx) Rebind(query string) string {
	return q.tx.Rebind(query)
}

func (q *QueryableTx) Select(dest any, query string, args ...any) error {
	return q.tx.Select(dest, query, args...)
}

func (q *QueryableTx) Get(dest any, query string, args ...any) error {
	return q.tx.Get(dest, query, args...)
}

func (q *QueryableTx) Exec(query string, args ...any) (sql.Result, error) {
	return q.tx.Exec(query, args...)
}

func (q *QueryableTx) NamedExec(query string, arg any) (sql.Result, error) {
	return q.tx.NamedExec(query, arg)
}

func (q *QueryableTx) Queryx(query string, args ...interface{}) (*sqlx.Rows, error) {
	return q.tx.Queryx(query, args...)
}
