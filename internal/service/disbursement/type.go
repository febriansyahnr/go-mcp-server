package disbursementService

import (
	pdkLogger "github.com/paper-indonesia/pdk/v2/logger"
	"github.com/paper-indonesia/pg-mcp-server/config"
	"github.com/paper-indonesia/pg-mcp-server/internal/repository"
	"github.com/paper-indonesia/pg-mcp-server/internal/service"
	"go.opentelemetry.io/otel"
)

var tracer = otel.Tracer("DisbursementService")

type DisbursementService struct {
	backendPortal repository.IBackendPortal
	snapCore      repository.ISnapCore
	logger        pdkLogger.ILogger
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

func WithSnapCoreRepository(snapCore repository.ISnapCore) dependencyFunc {
	return func(repo *DisbursementService) {
		repo.snapCore = snapCore
	}
}

func WithLogger(logger pdkLogger.ILogger) dependencyFunc {
	return func(repo *DisbursementService) {
		repo.logger = logger
	}
}

var _ service.IDisbursementService = (*DisbursementService)(nil)
