package disbursementHandler

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	disbursementModel "github.com/paper-indonesia/pg-mcp-server/internal/model/disbursement"
)

func (h *DisbursementHandler) CheckDisbursementTransactionHandler(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {

	referenceID, err := req.RequireString("reference_id")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	response, err := h.disbursementService.CheckDisbursementTransaction(ctx, &disbursementModel.CheckDisbursementTransactionCmd{
		ReferenceID: referenceID,
	})
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	jsonStr, err := response.ToJSON()
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return mcp.NewToolResultText(jsonStr), nil
}
