package disbursementService

import (
	"github.com/paper-indonesia/pg-mcp-server/config"
	"github.com/paper-indonesia/pg-mcp-server/internal/repository"
	"github.com/paper-indonesia/pg-mcp-server/internal/service"
	"go.opentelemetry.io/otel"
)

var tracer = otel.Tracer("DisbursementService")

type DisbursementService struct {
	backendPortal repository.IBackendPortal
}

type dependencyFunc func(*DisbursementService)

func New(conf *config.Config, deps ...dependencyFunc) *DisbursementService {
	repo := &DisbursementService{}
	for _, dep := range deps {
		dep(repo)
	}
	return repo
}

func WithBackendPortalRepository(backendPortal repository.IBackendPortal) dependencyFunc {
	return func(repo *DisbursementService) {
		repo.backendPortal = backendPortal
	}
}

var _ service.IDisbursementService = (*DisbursementService)(nil)
