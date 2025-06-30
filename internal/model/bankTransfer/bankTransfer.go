package bankTransferModel

import (
	"fmt"
	"strings"

	"github.com/paper-indonesia/pg-mcp-server/pkg/util"
)

type BankTransfer struct {
	UUID               string `json:"uuid" db:"uuid"`
	PartnerReferenceNo string `json:"partner_reference_no" db:"partner_reference_no"`
	BankAcquirer       string `json:"bank_acquirer" db:"bank_acquirer"`
	BankReferenceNo    string `json:"bank_reference_no" db:"bank_reference_no"`
	BeneficiaryID      string `json:"beneficiary_id" db:"beneficiary_id"`
	ExternalID         string `json:"external_id" db:"external_id"`
	TransferType       string `json:"transfer_type" db:"transfer_type"`
	Status             string `json:"status" db:"status"`
}

type BankTransferQuery struct {
	UUID       string
	ExternalID string
}

func (q *BankTransferQuery) Query() string {
	conditions := []string{}
	if q.UUID != "" {
		sanitizedUUID := util.SanitizeInput(q.UUID)
		conditions = append(conditions, fmt.Sprintf("uuid = '%s'", sanitizedUUID))
	}
	if q.ExternalID != "" {
		sanitizedExternalID := util.SanitizeInput(q.ExternalID)
		conditions = append(conditions, fmt.Sprintf("external_id = '%s'", sanitizedExternalID))
	}
	return strings.Join(conditions, " AND ")
}

func (q *BankTransferQuery) SetUUID(uuid string) {
	q.UUID = uuid
}

func (q *BankTransferQuery) SetExternalID(externalID string) {
	if q.UUID != "" {
		q.UUID = ""
	}
	q.ExternalID = externalID
}
