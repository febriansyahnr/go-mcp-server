package snapcoreRepository

import (
	pdkLogger "github.com/paper-indonesia/pdk/v2/logger"
	"github.com/paper-indonesia/pg-mcp-server/config"
	"github.com/paper-indonesia/pg-mcp-server/internal/repository"
	"github.com/paper-indonesia/pg-mcp-server/pkg/mySqlExt"
	"go.opentelemetry.io/otel"
)

var tracer = otel.Tracer("BankTransferRepository")
var tableName = "bank_transfer"

type BankTransferRepository struct {
	dbClient mySqlExt.IMySqlExt
	logger   pdkLogger.ILogger
}

type dependencyFunc func(*BankTransferRepository)

func New(conf *config.Config, deps ...dependencyFunc) *BankTransferRepository {
	repo := &BankTransferRepository{}
	for _, dep := range deps {
		dep(repo)
	}
	return repo
}

func WithDBClient(dbClient mySqlExt.IMySqlExt) dependencyFunc {
	return func(repo *BankTransferRepository) {
		repo.dbClient = dbClient
	}
}

func WithLogger(logger pdkLogger.ILogger) dependencyFunc {
	return func(repo *BankTransferRepository) {
		repo.logger = logger
	}
}

var _ repository.ISnapCore = (*BankTransferRepository)(nil)
