package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/paper-indonesia/pg-mcp-server/pkg/gcloudlogging"
)

func main() {
	ctx := context.Background()
	
	opts := gcloudlogging.QueryOptions{
		ProjectID: "your-gcp-project-id",
	}
	
	logService, err := gcloudlogging.NewLogQueryService(opts)
	if err != nil {
		log.Fatalf("Failed to create log service: %v", err)
	}
	defer logService.Close()
	
	since := time.Now().Add(-24 * time.Hour)
	maxResults := 10
	
	fmt.Println("=== Example 1: Query Error Logs ===")
	if err := queryErrorLogs(ctx, logService, since, maxResults); err != nil {
		log.Printf("Error querying error logs: %v", err)
	}
	
	fmt.Println("\n=== Example 2: Query Kubernetes Container Logs ===")
	if err := queryKubernetesLogs(ctx, logService, since, maxResults); err != nil {
		log.Printf("Error querying k8s logs: %v", err)
	}
	
	fmt.Println("\n=== Example 3: Query Application Logs ===")
	if err := queryApplicationLogs(ctx, logService, since, maxResults); err != nil {
		log.Printf("Error querying app logs: %v", err)
	}
	
	fmt.Println("\n=== Example 4: Query Database Logs ===")
	if err := queryDatabaseLogs(ctx, logService, since, maxResults); err != nil {
		log.Printf("Error querying db logs: %v", err)
	}
	
	fmt.Println("\n=== Example 5: Query HTTP 500 Errors ===")
	if err := queryHTTPErrors(ctx, logService, since, maxResults); err != nil {
		log.Printf("Error querying HTTP errors: %v", err)
	}
	
	fmt.Println("\n=== Example 6: Custom Filter Query ===")
	if err := queryCustomFilter(ctx, logService, maxResults); err != nil {
		log.Printf("Error with custom query: %v", err)
	}
}

func queryErrorLogs(ctx context.Context, service *gcloudlogging.LogQueryService, since time.Time, maxResults int) error {
	entries, err := service.QueryErrorLogs(ctx, since, maxResults)
	if err != nil {
		return err
	}
	
	fmt.Printf("Found %d error logs:\n", len(entries))
	for i, entry := range entries {
		fmt.Printf("%d. [%s] %s: %s\n", 
			i+1, 
			entry.Timestamp.Format(time.RFC3339), 
			entry.Severity,
			truncateString(getPayloadText(entry), 100))
	}
	
	return nil
}

func queryKubernetesLogs(ctx context.Context, service *gcloudlogging.LogQueryService, since time.Time, maxResults int) error {
	entries, err := service.QueryLogsByResource(ctx, "k8s_container", since, maxResults)
	if err != nil {
		return err
	}
	
	fmt.Printf("Found %d Kubernetes container logs:\n", len(entries))
	for i, entry := range entries {
		containerName := "unknown"
		if entry.Resource != nil {
			if labels, ok := entry.Resource["labels"].(map[string]any); ok {
				if name, exists := labels["container_name"]; exists {
					containerName = fmt.Sprintf("%v", name)
				}
			}
		}
		
		fmt.Printf("%d. [%s] Container: %s, %s\n", 
			i+1, 
			entry.Timestamp.Format(time.RFC3339), 
			containerName,
			truncateString(getPayloadText(entry), 80))
	}
	
	return nil
}

func queryApplicationLogs(ctx context.Context, service *gcloudlogging.LogQueryService, since time.Time, maxResults int) error {
	appName := "pg-mcp-go"
	severity := "WARNING"
	
	entries, err := service.QueryApplicationLogs(ctx, appName, severity, since, maxResults)
	if err != nil {
		return err
	}
	
	fmt.Printf("Found %d application logs for %s (severity >= %s):\n", len(entries), appName, severity)
	for i, entry := range entries {
		fmt.Printf("%d. [%s] %s: %s\n", 
			i+1, 
			entry.Timestamp.Format(time.RFC3339), 
			entry.Severity,
			truncateString(getPayloadText(entry), 100))
	}
	
	return nil
}

func queryDatabaseLogs(ctx context.Context, service *gcloudlogging.LogQueryService, since time.Time, maxResults int) error {
	operation := "SELECT"
	
	entries, err := service.QueryDatabaseLogs(ctx, operation, since, maxResults)
	if err != nil {
		return err
	}
	
	fmt.Printf("Found %d database logs for operation '%s':\n", len(entries), operation)
	for i, entry := range entries {
		fmt.Printf("%d. [%s] %s\n", 
			i+1, 
			entry.Timestamp.Format(time.RFC3339),
			truncateString(getPayloadText(entry), 120))
	}
	
	return nil
}

func queryHTTPErrors(ctx context.Context, service *gcloudlogging.LogQueryService, since time.Time, maxResults int) error {
	statusCode := 500
	
	entries, err := service.QueryHTTPRequests(ctx, statusCode, since, maxResults)
	if err != nil {
		return err
	}
	
	fmt.Printf("Found %d HTTP %d errors:\n", len(entries), statusCode)
	for i, entry := range entries {
		method := "unknown"
		url := "unknown"
		
		if entry.JSONPayload != nil {
			if m, ok := entry.JSONPayload["method"]; ok {
				method = fmt.Sprintf("%v", m)
			}
			if u, ok := entry.JSONPayload["url"]; ok {
				url = fmt.Sprintf("%v", u)
			}
		}
		
		fmt.Printf("%d. [%s] %s %s: %s\n", 
			i+1, 
			entry.Timestamp.Format(time.RFC3339), 
			method,
			url,
			truncateString(getPayloadText(entry), 80))
	}
	
	return nil
}

func queryCustomFilter(ctx context.Context, service *gcloudlogging.LogQueryService, maxResults int) error {
	customFilter := `
		(severity>=ERROR OR jsonPayload.level="error") AND
		(textPayload:"timeout" OR jsonPayload.error:"timeout") AND
		timestamp>="2024-06-01T00:00:00Z"
	`
	
	entries, err := service.QueryCustomFilter(ctx, customFilter, maxResults)
	if err != nil {
		return err
	}
	
	fmt.Printf("Found %d logs matching custom filter:\n", len(entries))
	for i, entry := range entries {
		fmt.Printf("%d. [%s] %s: %s\n", 
			i+1, 
			entry.Timestamp.Format(time.RFC3339), 
			entry.Severity,
			truncateString(getPayloadText(entry), 100))
	}
	
	return nil
}

func getPayloadText(entry *gcloudlogging.LogEntry) string {
	if entry.TextPayload != "" {
		return entry.TextPayload
	}
	
	if entry.JSONPayload != nil {
		if msg, ok := entry.JSONPayload["message"]; ok {
			return fmt.Sprintf("%v", msg)
		}
		if msg, ok := entry.JSONPayload["msg"]; ok {
			return fmt.Sprintf("%v", msg)
		}
		return fmt.Sprintf("%v", entry.JSONPayload)
	}
	
	return "No payload"
}

func truncateString(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen] + "..."
}