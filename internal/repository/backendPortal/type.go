package backendportalRepository

import (
	pdkLogger "github.com/paper-indonesia/pdk/v2/logger"
	"github.com/paper-indonesia/pg-mcp-server/config"
	"github.com/paper-indonesia/pg-mcp-server/internal/repository"
	"github.com/paper-indonesia/pg-mcp-server/pkg/mySqlExt"
	"go.opentelemetry.io/otel"
)

var tracer = otel.Tracer("BackendPortalRepository")
var tableName = "account_transactions"

type BackendPortalRepository struct {
	dbClient mySqlExt.IMySqlExt
	logger   pdkLogger.ILogger
}

type dependencyFunc func(*BackendPortalRepository)

func New(conf *config.Config, deps ...dependencyFunc) *BackendPortalRepository {
	repo := &BackendPortalRepository{}
	for _, dep := range deps {
		dep(repo)
	}
	return repo
}

func WithDBClient(dbClient mySqlExt.IMySqlExt) dependencyFunc {
	return func(repo *BackendPortalRepository) {
		repo.dbClient = dbClient
	}
}

func WithLogger(logger pdkLogger.ILogger) dependencyFunc {
	return func(repo *BackendPortalRepository) {
		repo.logger = logger
	}
}

var _ repository.IBackendPortal = (*BackendPortalRepository)(nil)
