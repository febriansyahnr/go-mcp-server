package disbursementModel

import (
	"fmt"
	"strings"

	"github.com/paper-indonesia/pg-mcp-server/pkg/util"
)

type DisbursementTransaction struct {
	UUID                   string `json:"uuid" db:"uuid"`
	ReferenceID            string `json:"reference_id" db:"reference_id"`
	MerchantID             string `json:"merchant_id" db:"merchant_id"`
	ProcessorReference     string `json:"processor_reference" db:"processor_reference"`
	ProcessorReferenceID   string `json:"processor_reference_id" db:"processor_reference_id"`
	ProcessorTransactionID string `json:"processor_transaction_id" db:"processor_transaction_id"`
	MerchantName           string `json:"merchant_name" db:"merchant_name"`
	Amount                 string `json:"amount" db:"amount"`
	Currency               string `json:"currency" db:"currency"`
	BeneficiaryAccountNo   string `json:"beneficiary_account_no" db:"beneficiary_account_no"`
	BeneficiaryBankCode    string `json:"beneficiary_bank_code" db:"beneficiary_bank_code"`
	BeneficiaryBankName    string `json:"beneficiary_bank_name" db:"beneficiary_bank_name"`
	BankReferenceNo        string `json:"bank_reference_no" db:"bank_reference_no"`
	TransactionTimestamp   string `json:"transaction_timestamp" db:"transaction_timestamp"`
	Status                 string `json:"status" db:"status"`
}

type DisbusementTransactionQuery struct {
	ReferenceID string
	Prefix      struct {
		AccountTransactions string
	}
}

func (d *DisbusementTransactionQuery) Query() string {
	conditions := []string{}
	if d.ReferenceID != "" {
		sanitizedReferenceID := util.SanitizeInput(d.ReferenceID)
		field := fmt.Sprintf("%s.reference_id", d.Prefix.AccountTransactions)
		conditions = append(conditions, fmt.Sprintf("%s = '%s'", field, sanitizedReferenceID))
	}

	return strings.Join(conditions, " AND ")
}
