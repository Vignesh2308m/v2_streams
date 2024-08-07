package main

import (
	"context"
	"database/sql"
)

type Process struct {
	conn  *DuckDb
	query string
}

func NewProcess(db *DuckDb, query string) *Process {
	return &Process{
		conn:  db,
		query: query,
	}
}

func (p *Process) Run() *sql.Rows {
	res, err := sql.OpenDB(p.conn.conn).QueryContext(context.Background(), p.query)
	if err != nil {
		print(err)
	}

	return res
}
