// Author(s): Carl Saldanha
// Database Connector and Helper

package model

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type DatabaseManager struct {
	database *sql.DB
}

func NewDatabaseManager(fileName string) (*DatabaseManager, error) {
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		return nil, err
	}
	return &DatabaseManager{
		database: db,
	}, nil
}

func (d *DatabaseManager) runQuery() error {
	// insert
	stmt, err := d.database.Prepare("INSERT INTO userinfo(username, departname, created) values(?,?,?)")
	if err != nil {
		return err
	}

	res, err := stmt.Exec("astaxie", "FF", "DD")

	_, err = res.LastInsertId()
	if err != nil {
		return err
	}
	return nil
}

func (d *DatabaseManager) close() {
	d.database.Close()
}
