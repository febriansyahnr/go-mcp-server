# Google Cloud Logging Query Examples

This directory contains examples for querying Google Cloud Logging using Go.

## Setup

1. Install dependencies:
```bash
go get cloud.google.com/go/logging
```

2. Set up authentication:
```bash
# Option 1: Service Account Key
export GOOGLE_APPLICATION_CREDENTIALS="/path/to/service-account-key.json"

# Option 2: Application Default Credentials
gcloud auth application-default login
```

3. Set your project ID:
```bash
export GOOGLE_CLOUD_PROJECT="your-gcp-project-id"
```

## Usage

### Basic Usage

```go
package main

import (
    "context"
    "time"
    "github.com/paper-indonesia/pg-mcp-go/pkg/gcloudlogging"
)

func main() {
    opts := gcloudlogging.QueryOptions{
        ProjectID: "your-project-id",
    }
    
    service, err := gcloudlogging.NewLogQueryService(opts)
    if err != nil {
        panic(err)
    }
    defer service.Close()
    
    since := time.Now().Add(-24 * time.Hour)
    entries, err := service.QueryErrorLogs(context.Background(), since, 10)
    if err != nil {
        panic(err)
    }
    
    for _, entry := range entries {
        fmt.Printf("%s: %s\n", entry.Timestamp, entry.TextPayload)
    }
}
```

### Run Example

```bash
# Update project ID in the example file first
go run cmd/examples/gcloud_logging_example.go
```

## Available Query Methods

### 1. Query Error Logs
```go
entries, err := service.QueryErrorLogs(ctx, since, maxResults)
```

### 2. Query by Resource Type
```go
entries, err := service.QueryLogsByResource(ctx, "k8s_container", since, maxResults)
```

### 3. Query by Text Search
```go
entries, err := service.QueryLogsByTextSearch(ctx, "database error", since, maxResults)
```

### 4. Query by JSON Field
```go
entries, err := service.QueryLogsByJSONField(ctx, "user_id", "12345", since, maxResults)
```

### 5. Query Application Logs
```go
entries, err := service.QueryApplicationLogs(ctx, "app-name", "ERROR", since, maxResults)
```

### 6. Query Database Logs
```go
entries, err := service.QueryDatabaseLogs(ctx, "SELECT", since, maxResults)
```

### 7. Query HTTP Requests
```go
entries, err := service.QueryHTTPRequests(ctx, 500, since, maxResults)
```

### 8. Custom Filter Query
```go
customFilter := `severity>=ERROR AND timestamp>="2024-06-01T00:00:00Z"`
entries, err := service.QueryCustomFilter(ctx, customFilter, maxResults)
```

## Common Filter Examples

### Time-based Filters
```
timestamp>="2024-06-30T00:00:00Z"
timestamp>="2024-06-30T00:00:00Z" AND timestamp<="2024-06-30T23:59:59Z"
```

### Severity Filters
```
severity>=ERROR
severity="WARNING"
```

### Resource Filters
```
resource.type="k8s_container"
resource.type="gce_instance"
resource.labels.container_name="my-app"
```

### Text Search
```
textPayload:"database connection failed"
textPayload:"timeout"
```

### JSON Payload Search
```
jsonPayload.user_id="12345"
jsonPayload.level="error"
jsonPayload.message:"failed to process"
```

### Combined Filters
```
resource.type="k8s_container" AND
severity>=WARNING AND
timestamp>="2024-06-01T00:00:00Z" AND
(textPayload:"error" OR jsonPayload.level="error")
```

## Configuration

You can use the configuration struct for easier setup:

```go
config := &gcloudlogging.Config{
    ProjectID:         "your-project-id",
    CredentialsFile:   "/path/to/credentials.json",
    DefaultMaxResults: 100,
    DefaultTimeWindow: "24h",
}

if err := config.Validate(); err != nil {
    panic(err)
}

service, err := gcloudlogging.NewLogQueryService(config.ToQueryOptions())
```

## Error Handling

All methods return detailed error information:

```go
entries, err := service.QueryErrorLogs(ctx, since, maxResults)
if err != nil {
    log.Printf("Failed to query logs: %v", err)
    return
}
```

## Performance Tips

1. Use specific time ranges to limit results
2. Include resource type filters when possible
3. Set reasonable maxResults limits
4. Use newest first ordering for recent logs
5. Cache the client instance for multiple queries