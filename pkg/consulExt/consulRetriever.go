package consulExt

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/hashicorp/consul/api"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/thomaspoignant/go-feature-flag/retriever"
)

type ConsulRetriever struct {
	configPath string
	client     *api.Client
	logger     *log.Logger
	status     retriever.Status
	mu         sync.RWMutex
}

type IConsulRetriever interface {
	Init(ctx context.Context, logger *log.Logger) error
	Retrieve(ctx context.Context) ([]byte, error)
	Shutdown(ctx context.Context) error
	Status() retriever.Status
}

var (
	httpClient *http.Client
	once       sync.Once
)

func getHttpClient() *http.Client {
	once.Do(func() {
		httpClient = &http.Client{
			Transport: newrelic.NewRoundTripper(http.DefaultTransport),
		}
	})
	return httpClient
}

func New(addr, configPath string, token string) (IConsulRetriever, error) {
	httpClient := getHttpClient()

	config := &api.Config{
		Address:    addr,
		Scheme:     "http",
		HttpClient: httpClient,
		WaitTime:   10 * time.Second,
		Token:      token,
	}

	client, err := api.NewClient(config)
	if err != nil {
		return nil, err
	}

	return &ConsulRetriever{
		client:     client,
		configPath: configPath,
		status:     retriever.RetrieverNotReady,
	}, nil
}

func (c *ConsulRetriever) Init(
	ctx context.Context,
	logger *log.Logger) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.logger = logger

	_, err := c.retrieve(ctx)
	if err != nil {
		c.status = retriever.RetrieverError
		return err
	}

	c.status = retriever.RetrieverReady
	return nil
}

func (c *ConsulRetriever) Retrieve(ctx context.Context) ([]byte, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.retrieve(ctx)
}

func (c *ConsulRetriever) Shutdown(ctx context.Context) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.status = retriever.RetrieverNotReady
	return nil
}

func (c *ConsulRetriever) Status() retriever.Status {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.status
}

func (c *ConsulRetriever) retrieve(ctx context.Context) ([]byte, error) {
	kv := c.client.KV()
	pair, _, err := kv.Get(c.configPath, nil)
	if err != nil {
		fmt.Printf("Error getting value from consul: %v\n", err)
		return nil, err
	}
	if pair == nil {
		fmt.Printf("No value found for key %s\n", c.configPath)
		return nil, errors.New("no value found for key " + c.configPath)
	}
	return pair.Value, nil
}
