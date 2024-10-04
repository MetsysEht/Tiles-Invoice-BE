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
	_, err := tx.Exec(`CREATE TABLE jaquar_products (
    	id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    	series VARCHAR(5) NOT NULL,
    	color_code VARCHAR(5) NOT NULL,
    	code_number VARCHAR(15) NOT NULL,
		description VARCHAR(100) NOT NULL,
  		nrp BIGINT NOT NULL,
		mrp BIGINT NOT NULL,
    	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
		);`)
	if err != nil {
		return err
	}
	return nil
}

func downUsers(_ context.Context, tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE jaquar_products;")
	if err != nil {
		return err
	}
	return nil
}
