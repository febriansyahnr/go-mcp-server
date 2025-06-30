package gcloudlogging

import (
	"fmt"
	"os"

	"cloud.google.com/go/logging/logadmin"
)

type QueryOptions struct {
	ProjectID           string
	CredentialsFile     string
	MaxResults          int
	OrderBy             logadmin.EntriesOption
	AdditionalProjectIDs []string
}

type Config struct {
	ProjectID           string `yaml:"project_id" json:"project_id"`
	CredentialsFile     string `yaml:"credentials_file" json:"credentials_file"`
	DefaultMaxResults   int    `yaml:"default_max_results" json:"default_max_results"`
	DefaultTimeWindow   string `yaml:"default_time_window" json:"default_time_window"`
	AdditionalProjectIDs []string `yaml:"additional_project_ids" json:"additional_project_ids"`
}

func (c *Config) Validate() error {
	if c.ProjectID == "" {
		return fmt.Errorf("project_id is required")
	}
	
	if c.CredentialsFile != "" {
		if _, err := os.Stat(c.CredentialsFile); os.IsNotExist(err) {
			return fmt.Errorf("credentials file does not exist: %s", c.CredentialsFile)
		}
	}
	
	if c.DefaultMaxResults <= 0 {
		c.DefaultMaxResults = 100
	}
	
	if c.DefaultTimeWindow == "" {
		c.DefaultTimeWindow = "24h"
	}
	
	return nil
}

func DefaultConfig() *Config {
	return &Config{
		ProjectID:         os.Getenv("GOOGLE_CLOUD_PROJECT"),
		CredentialsFile:   os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"),
		DefaultMaxResults: 100,
		DefaultTimeWindow: "24h",
	}
}

func (c *Config) ToQueryOptions() *QueryOptions {
	return &QueryOptions{
		ProjectID:            c.ProjectID,
		CredentialsFile:      c.CredentialsFile,
		MaxResults:          c.DefaultMaxResults,
		AdditionalProjectIDs: c.AdditionalProjectIDs,
	}
}