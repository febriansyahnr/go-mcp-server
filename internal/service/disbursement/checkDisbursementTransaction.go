package disbursementService

import (
	"context"

	"github.com/paper-indonesia/pg-mcp-server/constant"
	disbursementModel "github.com/paper-indonesia/pg-mcp-server/internal/model/disbursement"
)

func (s *DisbursementService) CheckDisbursementTransaction(ctx context.Context, cmd *disbursementModel.CheckDisbursementTransactionCmd) (*disbursementModel.DisbursementTransactionChecked, error) {
	ctx, span := tracer.Start(ctx, "disbursementService/CheckDisbursementTransaction")
	defer span.End()

	response := &disbursementModel.DisbursementTransactionChecked{}

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

	return response, nil
}
