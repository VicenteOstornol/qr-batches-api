package user

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/VicenteOstornol/lotesapi/entities"
	"github.com/VicenteOstornol/lotesapi/pkg/errors"
	"github.com/jmoiron/sqlx"
)

const usersSchema = "users"

type UserRepository struct {
	postgres *sqlx.DB
}

func New(postgres *sqlx.DB) *UserRepository {
	return &UserRepository{
		postgres: postgres,
	}
}

func (p *UserRepository) Create(ctx context.Context, user *entities.User) (*entities.User, error) {
	query, args, err := sq.
		Insert(usersSchema).
		Columns("name").
		Values(
			user.Name,
		).
		Suffix("RETURNING *").
		PlaceholderFormat(sq.Dollar).
		ToSql()

	if err != nil {
		return nil, errors.Wrap(err, "users repository: Create ToSql error")
	}

	var createdUser entities.User
	err = p.postgres.QueryRowx(query, args...).StructScan(&createdUser)
	if err != nil {
		return nil, errors.Wrap(err, "users repository: Create postgres.Exec error")
	}

	return &createdUser, nil
}

func (p *UserRepository) GetBatchesByUserID(ctx context.Context, userID string) ([]entities.Batch, error) {
	query, args, err := sq.
		Select("id", "name", "user_id").
		From("batches").
		Where(sq.Eq{"user_id": userID}).
		PlaceholderFormat(sq.Dollar).
		ToSql()
	if err != nil {
		return nil, errors.Wrap(err, "users repository: GetBatchesByUserID ToSql error")
	}

	var batches []entities.Batch
	err = p.postgres.Select(&batches, query, args...)
	if err != nil {
		return nil, errors.Wrap(err, "users repository: GetBatchesByUserID postgres.Select error")
	}

	return batches, nil
}
