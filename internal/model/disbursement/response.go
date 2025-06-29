package disbursementModel

import "encoding/json"

type DisbursementTransactionChecked struct {
	DisbursementTransaction *DisbursementTransaction `json:"disbursement_transaction,omitempty"`
}

func (d *DisbursementTransactionChecked) ToJSON() (string, error) {
	jsonByte, err := json.Marshal(d)
	if err != nil {
		return "", err
	}

	return string(jsonByte), nil
}
