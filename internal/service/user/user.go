package user

import (
	"context"

	"github.com/VicenteOstornol/lotesapi/entities"
	"github.com/VicenteOstornol/lotesapi/internal/repository"
	"github.com/VicenteOstornol/lotesapi/pkg/errors"
)

type UserService struct {
	repository *repository.Repository
}

func New(repo *repository.Repository) *UserService {
	return &UserService{
		repository: repo,
	}
}

func (p *UserService) Create(ctx context.Context, user *entities.User) (*entities.User, error) {

	createdUser, err := p.repository.User.Create(ctx, user)
	if err != nil {
		return nil, errors.Wrap(err, "user: UserService.Create p.repository.User.Create error")
	}

	return createdUser, nil
}

func (p *UserService) GetBatchesByUserID(ctx context.Context, userID string) ([]entities.Batch, error) {
	batches, err := p.repository.User.GetBatchesByUserID(ctx, userID)
	if err != nil {
		return nil, errors.Wrap(err, "user: UserService.GetBatchesByUserID p.repository.User.GetBatchesByUserID error")
	}

	return batches, nil
}
