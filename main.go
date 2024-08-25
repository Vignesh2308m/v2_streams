package main

import (
	"context"
	"database/sql"
)

type ABS struct {
	a string
	b string
}

func main() {

	d := NewDuckDBConn("")

	_, err := d.Connect()
	if err != nil {
		print(err)
	}

	_, err = sql.OpenDB(d.conn).Exec("CREATE TABLE test (i INTEGER, j VARCHAR)")
	if err != nil {
		print("Table creation failed")
	}

	_, err = sql.OpenDB(d.conn).Exec("INSERT INTO test VALUES (1, 'a')")
	if err != nil {
		print("Insert failed")
	}

	res, err := sql.OpenDB(d.conn).QueryContext(context.Background(), "SELECT * FROM test")

	if err != nil {
		print("Query Error")
	}

	var abs string
	res.Next()
	res.Scan(&abs)
	print(abs)

}
