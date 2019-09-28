package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upUsers, downUsers)
}

func upUsers(tx *sql.Tx) error {
	_, err := tx.Exec(`
	CREATE TABLE users (
		id uuid NOT NULL PRIMARY KEY,
		username text,
		name text,
		email text,
		surname text,
		password text
	);
	`)
	if err != nil {
		return err
	}
	return nil
}

func downUsers(tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE users;")
	if err != nil {
		return err
	}
	return nil
}
