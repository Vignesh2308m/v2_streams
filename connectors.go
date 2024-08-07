package main

import (
	"context"
	"database/sql/driver"

	"github.com/marcboeker/go-duckdb"
)

type DuckDb struct {
	conn *duckdb.Connector
}

func NewDuckDBConn(db string) *DuckDb {
	conn, err := duckdb.NewConnector(db, nil)
	if err != nil {
		print(err)
	}

	return &DuckDb{
		conn: conn,
	}
}

func (d *DuckDb) Connect() (driver.Conn, error) {
	return d.conn.Connect(context.Background())
}
