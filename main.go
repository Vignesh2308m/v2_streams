package main

import (
	"database/sql"
	"fmt"
	"time"
)

func main() {

	d := NewDuckDBConn("")

	con, err := d.Connect()
	if err != nil {
		print(err)
	}

	sql.OpenDB(d.conn).Exec("CREATE TABLE test (i INTEGER)")

	a := NewAppender(con, "", "test")

	p := NewProcess(d, "SELECT COUNT(i) FROM test")

	t := time.Now()
	for i := 0; i < 10000; i++ {
		a.Append(int32(i))

		_ = p.Run()

		//		ConsoleWriter(res)

	}
	fmt.Println(t.Sub(time.Now()))

}
