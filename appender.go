package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/marcboeker/go-duckdb"
)

type Appender struct {
	appender_conn *duckdb.Appender
	schema_name   map[int8]string
	table_name    string
	counter       int
}

func NewAppender(con *DuckDb, schema, table string) *Appender {
	d, err := con.Connect()
	if err != nil {
		print("Unable to connect to duckdb")
	}
	a, err := duckdb.NewAppenderFromConn(d, schema, table)
	if err != nil {
		print(err)
	}

	res, err := sql.OpenDB(con.conn).QueryContext(context.Background(), fmt.Sprint("DESCRIBE %s", table))

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

	return &Appender{}
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
