package migrations

import (
	"context"
	"database/sql"

	"github.com/VicenteOstornol/lotesapi/pkg/errors"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upAddBatchSchema, downAddBatchSchema)
}

func upAddBatchSchema(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, `
		CREATE TABLE IF NOT EXISTS batches (
			id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
			name TEXT NOT NULL,
			user_id UUID NOT NULL,
			amount_qrs INT DEFAULT 0,
			CONSTRAINT fk_user_id FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
		);
	`)
	if err != nil {
		return errors.Wrap(err, "migrations: upAddBatchSchema tx.ExecContext error")
	}
	return nil
}

func downAddBatchSchema(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	_, err := tx.ExecContext(ctx, `
		DROP TABLE IF EXISTS batches;
	`)

	if err != nil {
		return errors.Wrap(err, "migrations: downAddBatchSchema tx.ExecContext error")
	}

	return nil
}
