package main

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"fmt"

	"github.com/marcboeker/go-duckdb"
)

type Appender struct {
	appender_conn *duckdb.Appender
	schema_name   map[int8]string
	table_name    string
	counter       int
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

func (a *Appender) SetSchema(d *duckdb.Connector) {
	res, err := sql.OpenDB(d).QueryContext(context.Background(), fmt.Sprint("DESCRIBE %s", a.table_name))
	if err != nil {
		print("Set schema failed")
	}
	i := 0

	for res.Next() {
		var s *string
		res.Scan(&s)
		a.table_name[i] = s
		i++
	}
}
func (a *Appender) CloseAppender() {
	a.appender_conn.Close()
}

func (a *Appender) Append(inp chan map[string]interface{}) {
	for i := range inp {
		a.appender_conn.AppendRow(
			i["key"],
			i["val"])
		a.appender_conn.Flush()
	}
}
