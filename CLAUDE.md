# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

**pg-mcp-go** is an enterprise-grade Model Context Protocol (MCP) server built in Go by Paper Indonesia. It provides MCP tools as HTTP/SSE services with comprehensive enterprise infrastructure integrations including databases, caching, messaging, monitoring, and observability.

## Development Commands

### Core Development
```bash
# Run HTTP server
go run main.go serveHTTP --config .config.yaml --secret .secret.yaml

# Run SSE server  
make run-sse

# Run consumer mode
make run-consumer

# Run all tests
go test ./...

# Run specific package tests
go test -v ./pkg/util/...
go test -v ./internal/service/...
```

### Mock Generation
```bash
# Generate all mocks
make gen-mocks

# Generate specific mocks
make gen-mock-repo      # Repository mocks
make gen-mock-service   # Service mocks  
make gen-pkg-mock       # Package mocks
```

## Architecture

### Clean Architecture Layers
- **cmd/**: CLI commands and server setup using Cobra
- **internal/**: Private application code
  - **handlers/**: HTTP and MCP request handlers
  - **service/**: Business logic layer
  - **repository/**: Data access layer
  - **model/**: Data structures and domain objects
- **pkg/**: Reusable packages with enterprise integrations
- **tools/**: MCP tool definitions and schemas

### Enterprise Infrastructure
- **Database**: MySQL with master/slave configuration via PDK
- **Caching**: Redis with rate limiting and distributed locking
- **Messaging**: RabbitMQ for async communication
- **Monitoring**: New Relic, OpenTelemetry, Statsd via PDK
- **Service Discovery**: Consul integration
- **Secrets**: Vault for production secret management
- **Cloud Storage**: Google Cloud Storage integration

### Configuration
- **Config**: `.config.yaml` for application settings
- **Secrets**: `.secret.yaml` for sensitive data (database credentials, API keys)
- Both files required for all server commands

## MCP Tool Development

### Tool Structure
```go
// Define tool schema in tools/
var NewTool = mcp.NewTool("tool_name",
    mcp.WithDescription("Tool description"),
    mcp.WithString("param1", mcp.Required(), mcp.Description("Parameter description")),
)

// Implement handler in internal/handlers/extra/
func (h *Handler) HandleToolName(ctx context.Context, params map[string]interface{}) (*mcp.CallToolResult, error) {
    // Implementation
}
```

### Adding New Tools
1. Define tool schema in `tools/` directory
2. Implement handler in `internal/handlers/extra/`
3. Register tool in MCP server setup (`cmd/setupMCP.go`)

## Development Dependencies

### Paper Indonesia PDK
- `github.com/paper-indonesia/pdk/v2` - Platform Development Kit
- Provides standardized logger, MySQL, Redis, New Relic integrations
- Follow PDK patterns for consistency across Paper Indonesia services

### Core MCP Framework
- `github.com/mark3labs/mcp-go` - MCP protocol implementation
- Used for tool definitions and server setup

## Testing

### Mock Generation
- Uses Mockery for interface mocking
- Mocks generated in `mocks/` directory
- Regenerate mocks after interface changes: `make gen-mocks`

### Test Structure
- Unit tests alongside source files (`*_test.go`)
- Comprehensive coverage for utilities, handlers, and services
- Mock-based testing for external dependencies

## Key Patterns

### Error Handling
- Custom error types in `pkg/error/`
- Structured error responses with proper HTTP status codes
- Error logging with context and tracing

### Logging
- PDK-based structured logging with Zap
- Context-aware logging throughout the application
- Environment-specific log levels

### Database Access
- Repository pattern for data access
- PDK MySQL extensions with connection pooling
- Master/slave configuration support
- OpenTelemetry tracing for database operations

### Configuration Management
- Viper-based configuration loading
- Environment-specific configs
- Separate secret management for sensitive data

## Environment Setup

### Prerequisites
- Go 1.24.2+
- MySQL database
- Redis server
- RabbitMQ (optional)

### Local Development
1. Configure `.config.yaml` with local service endpoints
2. Set up `.secret.yaml` with database credentials
3. Ensure MySQL and Redis are running
4. Run: `go run main.go serveHTTP --config .config.yaml --secret .secret.yaml`

## Service Architecture

### Dependency Injection
- Constructor-based dependency injection pattern
- Service layer depends on repository interfaces
- Handlers depend on service interfaces
- Facilitates testing with mock implementations

### Observability
- Comprehensive tracing with OpenTelemetry
- Metrics collection via Statsd and New Relic
- Structured logging with correlation IDs
- Health checks and service monitoring

This MCP server follows enterprise patterns with robust infrastructure integrations, making it suitable for production deployment with comprehensive monitoring and observability.