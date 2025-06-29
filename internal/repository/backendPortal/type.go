package backendportalrepository

import (
	pdkLogger "github.com/paper-indonesia/pdk/v2/logger"
	"github.com/paper-indonesia/pg-mcp-server/config"
	"github.com/paper-indonesia/pg-mcp-server/pkg/mySqlExt"
)

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
