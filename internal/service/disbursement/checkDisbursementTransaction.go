package disbursementService

import (
	"context"

	pdkLogger "github.com/paper-indonesia/pdk/v2/logger"
	"github.com/paper-indonesia/pg-mcp-server/constant"
	bankTransferModel "github.com/paper-indonesia/pg-mcp-server/internal/model/bankTransfer"
	disbursementModel "github.com/paper-indonesia/pg-mcp-server/internal/model/disbursement"
)

func (s *DisbursementService) CheckDisbursementTransaction(ctx context.Context, cmd *disbursementModel.CheckDisbursementTransactionCmd) (*disbursementModel.DisbursementTransactionChecked, error) {
	ctx, span := tracer.Start(ctx, "disbursementService/CheckDisbursementTransaction")
	defer span.End()

	response := &disbursementModel.DisbursementTransactionChecked{}

	// First get the transaction as it's required for bank transfer lookup
	transaction, err := s.backendPortal.
		FindDisbursementTransaction(ctx, &disbursementModel.DisbusementTransactionQuery{
			ReferenceID: cmd.ReferenceID,
			Prefix: struct{ AccountTransactions string }{
				AccountTransactions: "at",
			},
		})
	if err != nil {
		return response, err
	}

	if transaction == nil {
		return response, constant.ErrRecordNotFound
	}

	response.DisbursementTransaction = transaction

	bankTransfer, bankTransferErr := s.snapCore.
		FindBankTransfer(ctx, &bankTransferModel.BankTransferQuery{
			ExternalID: transaction.UUID,
		})

	if bankTransferErr != nil {
		s.logger.Warn(ctx, "CheckDisbursementTransaction - error when getting bank transfer by external id", pdkLogger.Error(bankTransferErr), pdkLogger.String("external_id", transaction.UUID))
		return response, nil
	}

	response.Banktransfer = bankTransfer

	return response, nil
}
