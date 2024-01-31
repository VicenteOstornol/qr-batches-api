package batch

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/VicenteOstornol/lotesapi/entities"
	"github.com/VicenteOstornol/lotesapi/pkg/errors"
	"github.com/jmoiron/sqlx"
)

const batchSchema = "batches"

type BatchRepository struct {
	postgres *sqlx.DB
}

func New(postgres *sqlx.DB) *BatchRepository {
	return &BatchRepository{
		postgres: postgres,
	}
}

func (p *BatchRepository) Create(ctx context.Context, batch *entities.Batch) (*entities.Batch, error) {

	query, args, err := sq.
		Insert(batchSchema).
		Columns("name", "user_id", "amount_qrs").
		Values(
			batch.Name,
			batch.UserID,
			batch.AmountQrs,
		).
		Suffix("RETURNING *").
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, errors.Wrap(err, "batch repository: Create ToSql error")
	}

	var createdBatch entities.Batch
	err = p.postgres.QueryRowx(query, args...).StructScan(&createdBatch)
	if err != nil {
		return nil, errors.Wrap(err, "batch repository: Create postgres.Exec error")
	}

	return &createdBatch, nil
}

func (p *BatchRepository) GetByID(ctx context.Context, batchID string) (*entities.Batch, error) {
	query, args, err := sq.
		Select("*").
		From(batchSchema).
		Where(sq.Eq{"id": batchID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, errors.Wrap(err, "batch repository: GetByID ToSql error")
	}

	var batch entities.Batch
	err = p.postgres.Get(&batch, query, args...)
	if err != nil {
		return nil, errors.Wrap(err, "batch repository: GetByID postgres.Get error")
	}

	return &batch, nil
}
