package disbursementHandler

import (
	"github.com/paper-indonesia/pg-mcp-server/config"
	"github.com/paper-indonesia/pg-mcp-server/internal/service"
)

type DisbursementHandler struct {
	disbursementService service.IDisbursementService
}

type dependencyFunc func(*DisbursementHandler)

func New(conf *config.Config, deps ...dependencyFunc) *DisbursementHandler {
	repo := &DisbursementHandler{}
	for _, dep := range deps {
		dep(repo)
	}
	return repo
}

func WithDisbursementService(disbursementService service.IDisbursementService) dependencyFunc {
	return func(h *DisbursementHandler) {
		h.disbursementService = disbursementService
	}
}
