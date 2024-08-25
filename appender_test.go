package main

import (
	"database/sql"
	"testing"
)

func TestAppender(t *testing.T) {
	d := NewDuckDBConn("")

	_, err := sql.OpenDB(d.conn).Exec("CREATE TABLE test(i INTEGER)")
	if err != nil {
		print(err)
	}

	a := NewAppender(d, "", "test")

	print(a.column_name[0])
	//t.Errorf("Value Mismatch", b, 1)

}
