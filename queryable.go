package queryable

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type Queryable interface {
	Rebind(query string) string
	Select(dest any, query string, args ...any) error
	Get(dest any, query string, args ...any) error
	Exec(query string, args ...any) (sql.Result, error)
	NamedExec(query string, arg any) (sql.Result, error)
	Queryx(query string, args ...interface{}) (*sqlx.Rows, error)
}
