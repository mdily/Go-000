package sql

import "database/sql"

type Rows struct {
	rows []interface{}
}

func (rs *Rows) Next() bool {
	return true
}

func Query() (*Rows, error) {
	return nil, sql.ErrNoRows
}
