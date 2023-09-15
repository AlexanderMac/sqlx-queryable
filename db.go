package queryable

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type QueryableDb struct {
	db *sqlx.DB
}

func (q *QueryableDb) Rebind(query string) string {
	return q.db.Rebind(query)
}

func (q *QueryableDb) Select(dest any, query string, args ...any) error {
	return q.db.Select(dest, query, args...)
}

func (q *QueryableDb) Get(dest any, query string, args ...any) error {
	return q.db.Get(dest, query, args...)
}

func (q *QueryableDb) Exec(query string, args ...any) (sql.Result, error) {
	return q.db.Exec(query, args...)
}

func (q *QueryableDb) NamedExec(query string, arg any) (sql.Result, error) {
	return q.db.NamedExec(query, arg)
}

func (q *QueryableDb) Queryx(query string, args ...interface{}) (*sqlx.Rows, error) {
	return q.db.Queryx(query, args...)
}
