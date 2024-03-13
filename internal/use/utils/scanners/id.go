package scanners

import "database/sql"

func Id(row *sql.Row) (id int, err error) {
	err = row.Scan(&id)
	return
}
