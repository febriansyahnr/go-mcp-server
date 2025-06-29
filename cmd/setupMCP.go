package cmd

import (
	"log"

	"github.com/mark3labs/mcp-go/server"
	"github.com/paper-indonesia/pg-mcp-server/config"
	disbursementHandler "github.com/paper-indonesia/pg-mcp-server/internal/handlers/disbursement"
	extraHandler "github.com/paper-indonesia/pg-mcp-server/internal/handlers/extra"
	"github.com/paper-indonesia/pg-mcp-server/tools"
)

type MCPServer struct {
	mcpServer           *server.MCPServer
	conf                *config.Config
	extraHandler        *extraHandler.ExtraHandler
	disbursementHandler *disbursementHandler.DisbursementHandler
}

type depFunc func(*MCPServer)

func NewMCPServer(conf *config.Config, deps ...depFunc) *MCPServer {
	s := server.NewMCPServer(
		conf.ServiceName,
		conf.ServiceVersion,
		server.WithToolCapabilities(true),
		server.WithRecovery(),
	)
	mcps := &MCPServer{
		conf:      conf,
		mcpServer: s,
	}
	for _, dep := range deps {
		dep(mcps)
	}
	return mcps
}

// setup mcp tools, register mcp tools
func (s *MCPServer) setupMCP() {
	s.mcpServer.AddTool(tools.CalculatorTool, s.extraHandler.CalculatorHandler)
	s.mcpServer.AddTool(tools.CheckDisbursementTransactionTool, s.disbursementHandler.CheckDisbursementTransactionHandler)
}

func WithExtraHandler(handler *extraHandler.ExtraHandler) depFunc {
	return func(s *MCPServer) {
		s.extraHandler = handler
	}
}

func WithDisbursementHandler(handler *disbursementHandler.DisbursementHandler) depFunc {
	return func(s *MCPServer) {
		s.disbursementHandler = handler
	}
}

func (s *MCPServer) StartSSE() {
	s.setupMCP()

	sseServer := server.NewSSEServer(s.mcpServer, server.WithBaseURL("/sse"))
	if err := sseServer.Start(":8080"); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
