package extraHandler

import (
	"context"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
)

func (h *ExtraHandler) CalculatorHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// Using helper functions for type-safe argument access
	op, err := request.RequireString("operation")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	x, err := request.RequireFloat("x")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	y, err := request.RequireFloat("y")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	var result float64
	switch op {
	case "add":
		result = x + y
	case "subtract":
		result = x - y
	case "multiply":
		result = x * y
	case "divide":
		if y == 0 {
			return mcp.NewToolResultError("cannot divide by zero"), nil
		}
		result = x / y
	}

	return mcp.NewToolResultText(fmt.Sprintf("%.2f", result)), nil
}
