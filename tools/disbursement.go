package tools

import "github.com/mark3labs/mcp-go/mcp"

var CheckDisbursementTransactionTool = mcp.NewTool(
	"check_disbursement_transaction",
	mcp.WithDescription("Check disbursement transaction"),
	mcp.WithString("reference_id",
		mcp.Required(),
		mcp.Description("Reference ID of disbursement transaction"),
	),
)
