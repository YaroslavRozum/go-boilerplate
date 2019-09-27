package migrations

import (
	"database/sql"
	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upProducts, downProducts)
}

func upProducts(tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	return nil
}

func downProducts(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	return nil
}
