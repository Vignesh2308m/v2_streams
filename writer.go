package main

import (
	"database/sql"
	"fmt"
)

func ConsoleWriter(res *sql.Rows) {

	for res.Next() {
		var b any
		res.Scan(&b)
		fmt.Println(b)
	}

}
