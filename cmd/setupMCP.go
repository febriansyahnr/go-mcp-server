package cmd

import (
	"log"

	"github.com/mark3labs/mcp-go/server"
	"github.com/paper-indonesia/pg-mcp-server/config"
	extraHandler "github.com/paper-indonesia/pg-mcp-server/internal/handlers/extra"
	"github.com/paper-indonesia/pg-mcp-server/tools"
)

type MCPServer struct {
	mcpServer    *server.MCPServer
	conf         *config.Config
	extraHandler *extraHandler.ExtraHandler
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

func (s *MCPServer) setupMCP() {
	s.mcpServer.AddTool(tools.CalculatorTool, s.extraHandler.CalculatorHandler)
}

func WithExtraHandler(handler *extraHandler.ExtraHandler) depFunc {
	return func(s *MCPServer) {
		s.extraHandler = handler
	}
}

func (s *MCPServer) StartSSE() {
	s.setupMCP()

	sseServer := server.NewSSEServer(s.mcpServer, server.WithBaseURL("/sse"))
	if err := sseServer.Start(":8080"); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
