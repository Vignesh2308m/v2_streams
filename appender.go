package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/marcboeker/go-duckdb"
)

type Appender struct {
	appender_conn *duckdb.Appender
	column_name   []string
	table_name    string
	counter       int
}

type tbl_desc struct {
	col_name  string
	col_type  string
	col_null  string
	col_key   string
	col_def   string
	col_extra string
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

	res, err := sql.OpenDB(con.conn).QueryContext(context.Background(), fmt.Sprintf("DESCRIBE %s", table))

	if err != nil {
		print("Set schema failed")
	}

	columns := []string{}

	for res.Next() {
		var t tbl_desc
		res.Scan(&t.col_name,
			&t.col_type,
			&t.col_null,
			&t.col_key,
			&t.col_def,
			&t.col_extra)
		columns = append(columns, t.col_name)
	}

	return &Appender{appender_conn: a,
		column_name: columns,
		table_name:  table,
		counter:     0}
}
func (a *Appender) CloseAppender() {
	a.appender_conn.Close()
}

func (a *Appender) Append(inp chan map[string]interface{}) {
	for i := range inp {
		s := []any{}

		for _, j := range a.column_name {
			s = append(s, i[j])
		}

		a.appender_conn.AppendRow(s)

	}
}
