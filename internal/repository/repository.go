package repository

import (
	"context"

	disbursementModel "github.com/paper-indonesia/pg-mcp-server/internal/model/disbursement"
)

type IBackendPortal interface {
	FindDisbursementTransaction(
		ctx context.Context,
		q *disbursementModel.DisbusementTransactionQuery,
	) (*disbursementModel.DisbursementTransaction, error)
}
