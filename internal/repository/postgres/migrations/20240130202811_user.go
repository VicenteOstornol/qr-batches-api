package migrations

import (
	"context"
	"database/sql"

	"github.com/VicenteOstornol/lotesapi/pkg/errors"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigrationContext(upAddUserSchema, downAddUserSchema)
}

func upAddUserSchema(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	_, err := tx.ExecContext(ctx, `
		CREATE TABLE IF NOT EXISTS users (
			id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
			name TEXT NOT NULL
		);
	`)

	if err != nil {
		return errors.Wrap(err, "migrations: upAddUserSchema tx.ExecContext error")
	}

	return nil
}

func downAddUserSchema(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	_, err := tx.ExecContext(ctx, `
		DROP TABLE IF EXISTS users;
	`)

	if err != nil {
		return errors.Wrap(err, "migrations: downAddUserSchema tx.ExecContext error")
	}

	return nil
}
