package repository

import (
	"context"

	bankTransferModel "github.com/paper-indonesia/pg-mcp-server/internal/model/bankTransfer"
	disbursementModel "github.com/paper-indonesia/pg-mcp-server/internal/model/disbursement"
)

type IBackendPortal interface {
	FindDisbursementTransaction(
		ctx context.Context,
		q *disbursementModel.DisbusementTransactionQuery,
	) (*disbursementModel.DisbursementTransaction, error)
}

type ISnapCore interface {
	FindBankTransfer(
		ctx context.Context,
		q *bankTransferModel.BankTransferQuery,
	) (*bankTransferModel.BankTransfer, error)
}
