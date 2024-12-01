package storage

import (
	"database/sql"
)

func RemoveFromTableUsingId(db *sql.DB, id []byte) {
	sql := "DELETE FROM notes WHERE id = ($1)"
	db.Exec(sql, id)
}
