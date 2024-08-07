package main

import (
	"database/sql/driver"

	"github.com/marcboeker/go-duckdb"
)

type Appender struct {
	appender_conn *duckdb.Appender
	schema_name   string
	table_name    string
}

func NewAppender(con driver.Conn, schema, table string) *Appender {
	a, err := duckdb.NewAppenderFromConn(con, schema, table)
	if err != nil {
		print(err)
	}

	return &Appender{
		appender_conn: a,
		schema_name:   schema,
		table_name:    table,
	}
}

func (a *Appender) CloseAppender() {
	a.appender_conn.Close()
}

func (a *Appender) Append(val int32) {
	a.appender_conn.AppendRow(val)
	if int32(val)%100 == 0 {
		a.appender_conn.Flush()
	}
}
