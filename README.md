# Extends database/sql package

[![Build Status](https://travis-ci.org/tomi77/go-sqlx.svg?branch=master)](https://travis-ci.org/tomi77/go-sqlx)
[![Coverage Status](https://coveralls.io/repos/github/tomi77/go-sqlx/badge.svg?branch=master)](https://coveralls.io/github/tomi77/go-sqlx?branch=master)
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
