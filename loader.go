package loader

import (
	"database/sql"
	"os"
)

func LoadFromDatabase(db *sql.DB) (bool, error) {
	dbName := "people"
	name := dbName + ".sql"
	sqlFile, err := os.ReadFile(name)
	if err != nil {
		return false, err
	}
	_, err = db.Query(string(sqlFile))
	if err != nil {
		return false, err
	}
	return true, nil
}
