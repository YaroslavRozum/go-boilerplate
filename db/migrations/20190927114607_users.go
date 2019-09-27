package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upUsers, downUsers)
}

func upUsers(tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	return nil
}

func downUsers(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}
