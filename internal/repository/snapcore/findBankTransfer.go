package snapcoreRepository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/paper-indonesia/pg-mcp-server/constant"
	bankTransferModel "github.com/paper-indonesia/pg-mcp-server/internal/model/bankTransfer"
)

func (r *BankTransferRepository) FindBankTransfer(ctx context.Context, q *bankTransferModel.BankTransferQuery) (*bankTransferModel.BankTransfer, error) {
	ctx, span := tracer.Start(ctx, "repository/snapcore/FindBankTransfer")
	defer span.End()

	sqlQuery := `
		select 
			uuid,
			partner_reference_no,
			bank_acquirer,
			bank_reference_no,
			beneficiary_id,
			external_id,
			transfer_type,
			status
		from %s
		%s
		limit 1
	`

	conditions := q.Query()

	if conditions != "" {
		conditions = "where " + conditions
	}

	query := fmt.Sprintf(sqlQuery, tableName, conditions)
	ctx = context.WithValue(ctx, constant.CtxSQLTableNameKey, tableName)

	var transaction bankTransferModel.BankTransfer
	err := r.dbClient.GetContext(ctx, &transaction, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &transaction, nil
}
