package gcloudlogging

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/logging/logadmin"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

type LogQueryService struct {
	projectID string
	client    *logadmin.Client
}


type LogEntry struct {
	Timestamp   time.Time
	Severity    string
	Resource    map[string]any
	TextPayload string
	JSONPayload map[string]any
	LogName     string
	InsertID    string
	Labels      map[string]string
}

func NewLogQueryService(opts QueryOptions) (*LogQueryService, error) {
	ctx := context.Background()
	
	var client *logadmin.Client
	var err error
	
	if opts.CredentialsFile != "" {
		client, err = logadmin.NewClient(ctx, opts.ProjectID, 
			option.WithCredentialsFile(opts.CredentialsFile))
	} else {
		client, err = logadmin.NewClient(ctx, opts.ProjectID)
	}
	
	if err != nil {
		return nil, fmt.Errorf("failed to create logging client: %v", err)
	}
	
	return &LogQueryService{
		projectID: opts.ProjectID,
		client:    client,
	}, nil
}

func (s *LogQueryService) Close() error {
	return s.client.Close()
}

func (s *LogQueryService) QueryErrorLogs(ctx context.Context, since time.Time, maxResults int) ([]*LogEntry, error) {
	filter := fmt.Sprintf(`
		severity>=ERROR AND
		timestamp>="%s"
	`, since.Format(time.RFC3339))
	
	return s.executeQuery(ctx, filter, maxResults)
}

func (s *LogQueryService) QueryLogsByResource(ctx context.Context, resourceType string, since time.Time, maxResults int) ([]*LogEntry, error) {
	filter := fmt.Sprintf(`
		resource.type="%s" AND
		timestamp>="%s"
	`, resourceType, since.Format(time.RFC3339))
	
	return s.executeQuery(ctx, filter, maxResults)
}

func (s *LogQueryService) QueryLogsByTextSearch(ctx context.Context, searchText string, since time.Time, maxResults int) ([]*LogEntry, error) {
	filter := fmt.Sprintf(`
		textPayload:"%s" AND
		timestamp>="%s"
	`, searchText, since.Format(time.RFC3339))
	
	return s.executeQuery(ctx, filter, maxResults)
}

func (s *LogQueryService) QueryLogsByJSONField(ctx context.Context, fieldPath, value string, since time.Time, maxResults int) ([]*LogEntry, error) {
	filter := fmt.Sprintf(`
		jsonPayload.%s="%s" AND
		timestamp>="%s"
	`, fieldPath, value, since.Format(time.RFC3339))
	
	return s.executeQuery(ctx, filter, maxResults)
}

func (s *LogQueryService) QueryCustomFilter(ctx context.Context, filter string, maxResults int) ([]*LogEntry, error) {
	return s.executeQuery(ctx, filter, maxResults)
}

func (s *LogQueryService) executeQuery(ctx context.Context, filter string, maxResults int) ([]*LogEntry, error) {
	iter := s.client.Entries(ctx, 
		logadmin.Filter(filter),
		logadmin.NewestFirst(),
	)
	
	var entries []*LogEntry
	count := 0
	
	for {
		if maxResults > 0 && count >= maxResults {
			break
		}
		
		entry, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("failed to get next entry: %v", err)
		}
		
		logEntry := &LogEntry{
			Timestamp: entry.Timestamp,
			Severity:  entry.Severity.String(),
			LogName:   entry.LogName,
			InsertID:  entry.InsertID,
			Labels:    entry.Labels,
		}
		
		if entry.Resource != nil {
			logEntry.Resource = map[string]any{
				"type":   entry.Resource.Type,
				"labels": entry.Resource.Labels,
			}
		}
		
		switch payload := entry.Payload.(type) {
		case string:
			logEntry.TextPayload = payload
		case map[string]any:
			logEntry.JSONPayload = payload
		}
		
		entries = append(entries, logEntry)
		count++
	}
	
	return entries, nil
}

func (s *LogQueryService) QueryApplicationLogs(ctx context.Context, appName string, severity string, since time.Time, maxResults int) ([]*LogEntry, error) {
	severityFilter := ""
	if severity != "" {
		severityFilter = fmt.Sprintf(" AND severity>=%s", severity)
	}
	
	filter := fmt.Sprintf(`
		(resource.type="k8s_container" OR resource.type="gce_instance") AND
		(resource.labels.container_name:"%s" OR jsonPayload.app="%s")%s AND
		timestamp>="%s"
	`, appName, appName, severityFilter, since.Format(time.RFC3339))
	
	return s.executeQuery(ctx, filter, maxResults)
}

func (s *LogQueryService) QueryDatabaseLogs(ctx context.Context, operation string, since time.Time, maxResults int) ([]*LogEntry, error) {
	filter := fmt.Sprintf(`
		(textPayload:"database" OR jsonPayload.component="database") AND
		(textPayload:"%s" OR jsonPayload.operation="%s") AND
		timestamp>="%s"
	`, operation, operation, since.Format(time.RFC3339))
	
	return s.executeQuery(ctx, filter, maxResults)
}

func (s *LogQueryService) QueryHTTPRequests(ctx context.Context, statusCode int, since time.Time, maxResults int) ([]*LogEntry, error) {
	filter := fmt.Sprintf(`
		(httpRequest.status=%d OR jsonPayload.status_code=%d) AND
		timestamp>="%s"
	`, statusCode, statusCode, since.Format(time.RFC3339))
	
	return s.executeQuery(ctx, filter, maxResults)
}