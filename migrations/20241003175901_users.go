package migrations

import (
	"context"
	"database/sql"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upUsers, downUsers)
}

func upUsers(_ context.Context, tx *sql.Tx) error {
	_, err := tx.Exec(`CREATE TABLE users (
    	id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    	name VARCHAR(100) NOT NULL,
    	email VARCHAR(100) NOT NULL UNIQUE,
    	password VARCHAR(255) NOT NULL,
    	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
		);`)
	if err != nil {
		return err
	}
	return nil
}

func downUsers(_ context.Context, tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE users;")
	if err != nil {
		return err
	}
	return nil
}
