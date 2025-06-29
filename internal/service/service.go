package service

import (
	"context"

	disbursementModel "github.com/paper-indonesia/pg-mcp-server/internal/model/disbursement"
)

type IDisbursementService interface {
	CheckDisbursementTransaction(ctx context.Context, cmd *disbursementModel.CheckDisbursementTransactionCmd) (*disbursementModel.DisbursementTransactionChecked, error)
}
