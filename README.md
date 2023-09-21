# sqlx-queryable

[![Build Status](https://github.com/AlexanderMac/sqlx-queryable/actions/workflows/ci.yml/badge.svg)](https://github.com/AlexanderMac/sqlx-queryable/actions/workflows/ci.yml)
[![MIT license](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)
[![GoDoc](https://pkg.go.dev/badge/github.com/AlexanderMac/sqlx-queryable)](https://pkg.go.dev/github.com/AlexanderMac/sqlx-queryable)

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
