package repository

import (
	"context"

	"github.com/VicenteOstornol/lotesapi/entities"
)

type UserRepository interface {
	Create(context.Context, *entities.User) (*entities.User, error)
	GetBatchesByUserID(context.Context, string) ([]entities.Batch, error)
}

type BatchRepository interface {
	Create(context.Context, *entities.Batch) (*entities.Batch, error)
	GetByID(context.Context, string) (*entities.Batch, error)
}
