# pg-mcp-go

**pg-mcp-go** is an enterprise-grade Model Context Protocol (MCP) server built in Go by Paper Indonesia Development Team. It provides MCP tools as HTTP/SSE services with comprehensive enterprise infrastructure integrations including databases, caching, messaging, monitoring, and observability.

## Features

- **MCP Protocol Support**: Full implementation of Model Context Protocol with HTTP and SSE endpoints
- **Enterprise Infrastructure**: Built-in integrations for MySQL, Redis, RabbitMQ, and cloud services
- **Clean Architecture**: Layered architecture with clear separation of concerns
- **Comprehensive Observability**: OpenTelemetry tracing, metrics collection, and structured logging
- **Production Ready**: Circuit breakers, rate limiting, health checks, and error handling
- **Extensible Tool System**: Easy-to-add MCP tools with schema validation

## Project Structure

```
pg-mcp-go/
├── cmd/                        # CLI commands and server setup
│   ├── cmd.go                  # Root command configuration
│   ├── serveHTTP.go           # HTTP server command
│   └── setupMCP.go            # MCP server setup
├── config/                     # Configuration management
│   ├── config.go              # Application configuration
│   ├── secret.go              # Secret management
│   └── ...
├── constant/                   # Application constants
│   ├── constant.go            # General constants
│   ├── error.go               # Error constants
│   └── ...
├── internal/                   # Private application code
│   ├── handlers/              # Request handlers
│   │   ├── extra/             # MCP tool handlers
│   │   │   ├── calculator_handler.go
│   │   │   └── type.go
│   │   └── handlers.go
│   ├── model/                 # Data structures
│   │   ├── common/            # Common models
│   │   ├── disbursement/      # Domain-specific models
│   │   └── model.go
│   ├── repository/            # Data access layer
│   │   ├── backendPortal/     # Backend portal repository
│   │   └── repository.go
│   └── service/               # Business logic layer
│       └── service.go
├── pkg/                       # Reusable packages
│   ├── consulExt/             # Consul service discovery
│   ├── dictionary/            # Data dictionary utilities
│   ├── error/                 # Error handling utilities
│   ├── gcs/                   # Google Cloud Storage
│   ├── httpRequestExt/        # HTTP client extensions
│   ├── logger/                # Structured logging
│   ├── monitor/               # Monitoring utilities
│   ├── mySqlExt/              # MySQL extensions
│   ├── rabbitMqExt/           # RabbitMQ messaging
│   ├── redisExt/              # Redis caching
│   ├── slackExt/              # Slack notifications
│   ├── util/                  # General utilities
│   └── validatorExt/          # Validation extensions
├── tools/                     # MCP tool definitions
│   ├── calculator.go          # Calculator tool schema
│   └── tools.go               # Tool registry
├── mocks/                     # Generated test mocks
├── main.go                    # Application entry point
├── Makefile                   # Build and development commands
├── CLAUDE.md                  # Claude Code instructions
└── README.md                  # This file
```

## Quick Start

### Prerequisites

- Go 1.24.2+
- MySQL database
- Redis server
- RabbitMQ (optional)

### Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd pg-mcp-go
```

2. Install dependencies:
```bash
go mod download
```

3. Configure the application:
```bash
# Copy and edit configuration files
cp .config.yaml.example .config.yaml
cp .secret.yaml.example .secret.yaml
```

4. Start the HTTP server:
```bash
go run main.go serveHTTP --config .config.yaml --secret .secret.yaml
```

### Development Commands

```bash
# Run HTTP server
go run main.go serveHTTP --config .config.yaml --secret .secret.yaml

# Run SSE server
make run-sse

# Run consumer mode
make run-consumer

# Run tests
go test ./...

# Generate mocks
make gen-mocks
```

## Architecture

### Clean Architecture Layers

- **cmd/**: CLI commands and server setup using Cobra
- **internal/**: Private application code with handlers, services, and repositories
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

## MCP Tool Development

### Adding New Tools

1. Define tool schema in `tools/` directory
2. Implement handler in `internal/handlers/extra/`
3. Register tool in MCP server setup

### Example Tool

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

## Configuration

The application uses two configuration files:

- `.config.yaml`: Application settings and service endpoints
- `.secret.yaml`: Sensitive data (database credentials, API keys)

Both files are required for all server commands.

## Testing

The project uses comprehensive testing with mock generation:

```bash
# Run all tests
go test ./...

# Run specific package tests
go test -v ./pkg/util/...
go test -v ./internal/service/...

# Generate mocks after interface changes
make gen-mocks
```

## Dependencies

### Core Dependencies

- **PDK**: `github.com/paper-indonesia/pdk/v2` - Platform Development Kit for standardized integrations
- **MCP Framework**: `github.com/mark3labs/mcp-go` - MCP protocol implementation
- **Cobra**: CLI framework for command-line interface
- **Viper**: Configuration management

### Enterprise Integrations

- MySQL with connection pooling and tracing
- Redis with distributed locking and rate limiting
- RabbitMQ for async messaging
- OpenTelemetry for observability
- New Relic for monitoring
- Consul for service discovery

## Contributing

1. Follow the established architecture patterns
2. Add tests for new functionality
3. Generate mocks after interface changes: `make gen-mocks`
4. Ensure all tests pass: `go test ./...`
5. Follow Go best practices and the existing code style

## License

[Add license information here]

## Support

[Add support information here]

## Contributors

- [Febriansyah](https://github.com/febriansyahnr)