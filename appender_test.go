package main

import (
	"context"
	"database/sql"
	"testing"
)

func TestAppender(t *testing.T) {
	d := NewDuckDBConn("")

	_, err := sql.OpenDB(d.conn).Exec("CREATE TABLE test(i INTEGER)")
	if err != nil {
		print(err)
	}

	con, err := d.Connect()
	if err != nil {
		print(err)
	}

	a := NewAppender(con, "", "test")
	a.Append(int32(1))

	res := sql.OpenDB(d.conn).QueryRowContext(context.Background(), "SELECT * FROM test")
	var b int32
	res.Scan(&b)
	print(b)
	//t.Errorf("Value Mismatch", b, 1)

}
