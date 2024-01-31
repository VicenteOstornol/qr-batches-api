package service

import (
	"context"

	"github.com/VicenteOstornol/lotesapi/entities"
	"github.com/go-pdf/fpdf"
)

type UserService interface {
	Create(context.Context, *entities.User) (*entities.User, error)
	GetBatchesByUserID(context.Context, string) ([]entities.Batch, error)
}

type BatchService interface {
	Create(context.Context, *entities.Batch) (*entities.Batch, error)
	Get(ctx context.Context, batchID string) (*entities.Batch, error)
	DownloadPDF(ctx context.Context, batchID string) (*fpdf.Fpdf, error)
}
