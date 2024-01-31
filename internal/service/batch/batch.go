package batch

import (
	"context"

	"github.com/VicenteOstornol/lotesapi/entities"
	"github.com/VicenteOstornol/lotesapi/internal/repository"
	"github.com/VicenteOstornol/lotesapi/pkg/errors"
	"github.com/VicenteOstornol/lotesapi/pkg/pdf"
	"github.com/go-pdf/fpdf"
)

type BatchService struct {
	repository *repository.Repository
}

func New(repo *repository.Repository) *BatchService {
	return &BatchService{
		repository: repo,
	}
}

func (p *BatchService) Create(ctx context.Context, batch *entities.Batch) (*entities.Batch, error) {

	createdBatch, err := p.repository.Batch.Create(ctx, batch)
	if err != nil {
		return nil, errors.Wrap(err, "batch: BatchService.Create p.repository.Batch.Create error")
	}

	return createdBatch, nil
}

func (p *BatchService) DownloadPDF(ctx context.Context, batchID string) (*fpdf.Fpdf, error) {
	batch, err := p.repository.Batch.GetByID(ctx, batchID)
	if err != nil {
		return nil, errors.Wrap(err, "batch: BatchService.Download p.repository.Batch.GetByID error")
	}

	pdfFile, err := pdf.GeneratePDF(batch)
	if err != nil {
		return nil, errors.Wrap(err, "batch: BatchService.Download pdf.GeneratePDF error")
	}

	return pdfFile, nil
}

func (p *BatchService) Get(ctx context.Context, batchID string) (*entities.Batch, error) {
	batch, err := p.repository.Batch.GetByID(ctx, batchID)
	if err != nil {
		return nil, errors.Wrap(err, "batch: BatchService.Get p.repository.Batch.GetByID error")
	}

	return batch, nil
}
