<p align="center">
  <h1 align="center">sqlx-queryable</h1>
  <p align="center">Provides an interface for identical sqlx.DB and sqlx.Tx methods.</p>
  <p align="center">
    <a href="https://github.com/alexandermac/sqlx-queryable/actions/workflows/ci.yml?query=branch%3Amaster"><img src="https://github.com/alexandermac/sqlx-queryable/actions/workflows/ci.yml/badge.svg" alt="Build Status"></a>
    <a href="https://goreportcard.com/report/github.com/alexandermac/sqlx-queryable"><img src="https://goreportcard.com/badge/github.com/alexandermac/sqlx-queryable" alt="Go Report Card"></a>
    <a href="https://pkg.go.dev/github.com/alexandermac/sqlx-queryable"><img src="https://pkg.go.dev/badge/github.com/alexandermac/sqlx-queryable.svg" alt="Go Docs"></a>
    <a href="LICENSE"><img src="https://img.shields.io/github/license/alexandermac/sqlx-queryable.svg" alt="License"></a>
    <a href="https://img.shields.io/github/v/tag/alexandermac/gosu"><img src="https://img.shields.io/github/v/tag/alexandermac/gosu" alt="GitHub tag"></a>
  </p>
</p>

sqlx-queryable provides the `Queryable` interface which wraps identical `sqlx.DB` and `sqlx.Tx` methods. The interface can be used when a function performs a database query directly or under a transaction.
Requires Golang v1.18 or greater.

### Install
```sh
go get github.com/alexandermac/sqlx-queryable
```

### Usage
```go
import queryable "github.com/alexandermac/sqlx-queryable"

// `q` can be sqlx.DB or sqlx.Tx
func getRecords(q Queryable) ([]Record, error) {
	var records []Record
	err := q.Select(&records, "SELECT * FROM records")
	if err != nil {
		panic(err)
	}

	return records, nil
}
```

### License
Licensed under the MIT license.

### Author
Alexander Mac
