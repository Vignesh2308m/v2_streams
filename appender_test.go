package main

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"testing"
)

func TestAppender(t *testing.T) {
	d := NewDuckDBConn("")
	defer d.conn.Close()

	_, err := sql.OpenDB(d.conn).Exec("CREATE TABLE test(i INTEGER)")
	if err != nil {
		print(err)
	}

	a := NewAppender(d, "", "test")
	defer a.appender_conn.Close()

	print(a.column_name[0])
	//t.Errorf("Value Mismatch", b, 1)
}

func TestAppendRow(t *testing.T) {
	d := NewDuckDBConn("")

	_, err := sql.OpenDB(d.conn).Exec("CREATE TABLE test(i INTEGER, j INTEGER)")
	if err != nil {
		print(err)
	}

	a := NewAppender(d, "", "test")

	a.appender_conn.AppendRow([]driver.Value{1, 2}...)
	a.appender_conn.Flush()

	res := sql.OpenDB(d.conn).QueryRowContext(context.Background(), "SELECT * FROM test")
	var c, b int
	res.Scan(&c, &b)
	print(c)
	print(b)
}
