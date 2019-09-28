package migrations

import (
	"database/sql"

	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upProducts, downProducts)
}

func upProducts(tx *sql.Tx) error {
	_, err := tx.Exec(`
	CREATE TABLE products (
		id uuid NOT NULL PRIMARY KEY,
		name text,
		price int
	);
	`)
	if err != nil {
		return err
	}
	return nil
}

func downProducts(tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE products;")
	if err != nil {
		return err
	}
	return nil
}
