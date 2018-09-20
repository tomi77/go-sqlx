# Extends database/sql package

[![Go Report Card](https://goreportcard.com/badge/github.com/tomi77/go-sqlx)](https://goreportcard.com/report/github.com/tomi77/go-sqlx)
[![GoDoc](https://godoc.org/github.com/tomi77/go-sqlx?status.svg)](https://godoc.org/github.com/tomi77/go-sqlx)

## Installation

~~~sh
go get -u github.com/tomi77/go-sqlx
~~~

## Usage

~~~go
import sqlx "github.com/tomi77/go-sqlx"

type Tab struct {
  ID    uint32
  Value sqlx.NullInt8
}
~~~
