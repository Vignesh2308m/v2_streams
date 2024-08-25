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

	sql.OpenDB(d.conn).Exec("CREATE TABLE test (i INTEGER)")

	res, err := sql.OpenDB(d.conn).QueryContext(context.Background(), "DESCRIBE test")

	if err != nil {
		print("Query Error")
	}

	var abs ABS
	res.Next()
	res.Scan(&abs.a, &abs.b)
	print(abs.a)
	print(abs.b)

}
