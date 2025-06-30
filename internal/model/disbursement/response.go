package disbursementModel

import (
	"encoding/json"

	bankTransferModel "github.com/paper-indonesia/pg-mcp-server/internal/model/bankTransfer"
)

type DisbursementTransactionChecked struct {
	DisbursementTransaction *DisbursementTransaction        `json:"disbursement_transaction,omitempty"`
	Banktransfer            *bankTransferModel.BankTransfer `json:"bank_transfer,omitempty"`
}

func (d *DisbursementTransactionChecked) ToJSON() (string, error) {
	jsonByte, err := json.Marshal(d)
	if err != nil {
		return "", err
	}

	return string(jsonByte), nil
}
