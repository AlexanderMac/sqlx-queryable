package queryable

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type Queryable interface {
	Exec(query string, args ...any) (sql.Result, error)
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	Prepare(query string) (*sql.Stmt, error)
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	Query(query string, args ...any) (*sql.Rows, error)
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	QueryRow(query string, args ...any) *sql.Row
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row

	BindNamed(query string, arg any) (string, []any, error)
	DriverName() string
	Get(dest any, query string, args ...any) error
	GetContext(ctx context.Context, dest any, query string, args ...any) error
	MustExec(query string, args ...any) sql.Result
	MustExecContext(ctx context.Context, query string, args ...any) sql.Result
	NamedExec(query string, arg any) (sql.Result, error)
	NamedExecContext(ctx context.Context, query string, arg any) (sql.Result, error)
	NamedQuery(query string, arg any) (*sqlx.Rows, error)
	PrepareNamed(query string) (*sqlx.NamedStmt, error)
	PrepareNamedContext(ctx context.Context, query string) (*sqlx.NamedStmt, error)
	Preparex(query string) (*sqlx.Stmt, error)
	PreparexContext(ctx context.Context, query string) (*sqlx.Stmt, error)
	QueryRowx(query string, args ...any) *sqlx.Row
	QueryRowxContext(ctx context.Context, query string, args ...any) *sqlx.Row
	Queryx(query string, args ...any) (*sqlx.Rows, error)
	QueryxContext(ctx context.Context, query string, args ...any) (*sqlx.Rows, error)
	Rebind(query string) string
	Select(dest any, query string, args ...any) error
	SelectContext(ctx context.Context, dest any, query string, args ...any) error
}
