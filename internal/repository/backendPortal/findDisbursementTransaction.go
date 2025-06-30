package backendportalRepository

import (
	"context"
	"database/sql"
	"fmt"

	pdkLogger "github.com/paper-indonesia/pdk/v2/logger"
	"github.com/paper-indonesia/pg-mcp-server/constant"
	disbursementModel "github.com/paper-indonesia/pg-mcp-server/internal/model/disbursement"
)

func (r *BackendPortalRepository) FindDisbursementTransaction(ctx context.Context, q *disbursementModel.DisbusementTransactionQuery) (*disbursementModel.DisbursementTransaction, error) {
	ctx, span := tracer.Start(ctx, "repository/disbursement/FindDisbursementTransaction")
	defer span.End()

	sqlQuery := `
		select 
			at.uuid,
			at.reference_id,
			at.merchant_id,
			at.processor_reference,
			at.processor_reference_id,
			at.processor_transaction_id,
			m.name as merchant_name,
			at.debit as amount,
			at.currency,
			d.beneficiary_account_no,
			d.beneficiary_bank_code,
			d.beneficiary_bank_name,
			d.bank_reference_no,
			at.transaction_timestamp,
			at.status 
		from 
			%s at
		left join disbursements d on at.reference_id = d.uuid
		left join merchants m on at.merchant_id = m.uuid
		%s
		and at.type = 'DISBURSEMENT' 
		and at.channel = 'BANK_TRANSFER' 
		limit 1
	`

	conditions := q.Query()
	if conditions != "" {
		conditions = "where " + conditions
	}

	ctx = context.WithValue(ctx, constant.CtxSQLTableNameKey, tableName)

	query := fmt.Sprintf(sqlQuery, tableName, conditions)

	var transaction disbursementModel.DisbursementTransaction
	err := r.dbClient.GetContext(ctx, &transaction, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		r.logger.Error(ctx, "FindDisbursementTransaction - error when getting disbursement transaction by uuid", pdkLogger.Error(err), pdkLogger.String("reference_id", q.ReferenceID))

		return nil, err
	}

	return &transaction, nil
}
