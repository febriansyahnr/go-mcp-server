package rabbitMqExt

import (
	"context"
	"fmt"
	"io"
	"log"
	"sync"

	pdkLogger "github.com/paper-indonesia/pdk/v2/logger"
	pdkNewRelic "github.com/paper-indonesia/pdk/v2/newRelicExt"
	"github.com/paper-indonesia/pg-mcp-server/config"

	"github.com/paper-indonesia/pdk/v2/amqp"
)

type RabbitMQExt interface {
	HealthCheck(ctx context.Context) (active bool, err error)
	Publish(ctx context.Context, message interface{}, document string, failedMsg ...amqp.Delivery) error
	Consume(signal context.Context, document string, process func(context.Context, []byte) error) error

	PublishAndWaitReply(ctx context.Context, document string, message interface{}) (event *amqp.Event, close io.Closer, err error)
	PublishToReplyQueue(ctx context.Context, replyToAddress string, msg amqp.Publishing) error

	Close() error
}

type rabbitMQExt struct {
	url        string
	connection *amqp.Connection
	logger     pdkLogger.ILogger
	ctx        context.Context
	once       *once

	nrApp pdkNewRelic.INewRelicExt
}

type optionFunc func(*rabbitMQExt)

func WithContext(ctx context.Context) optionFunc {
	return func(r *rabbitMQExt) {
		r.ctx = ctx
	}
}

func New(
	config config.RabbitMQConfig,
	secret config.RabbitMQSecret,
	logger pdkLogger.ILogger,
	nrApp pdkNewRelic.INewRelicExt,
	options ...optionFunc,
) (RabbitMQExt, error) {
	url := fmt.Sprintf("amqp://%s:%s@%s:%s",
		secret.Username,
		secret.Password,
		config.Host,
		config.Port)

	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	rabbit := &rabbitMQExt{
		connection: conn,
		url:        url,
		logger:     logger,
		nrApp:      nrApp,
		once: &once{
			directReplyTo: new(sync.Map),
		},
	}

	for _, option := range options {
		option(rabbit)
	}
	if rabbit.ctx == nil {
		rabbit.ctx = context.Background()
	}

	if err := rabbit.Setup(rabbit.ctx); err != nil {
		return nil, err
	}

	return rabbit, nil
}

func (r *rabbitMQExt) getChannel() (*amqp.Channel, error) {
	if r.connection == nil || r.connection.IsClosed() {
		conn, err := amqp.Dial(r.url)
		if err != nil {
			return nil, err
		}
		r.connection = conn
	}

	ch, err := r.connection.Channel()
	if err != nil {
		return nil, err
	}

	return ch, nil
}

func (r *rabbitMQExt) Close() error {
	defer log.Println("RabbitMQ health check stopped")

	if channelHc != nil {
		if err := channelHc.Close(); err != nil {
			log.Println("Error close health check channel:", err.Error())
		}
	}
	return r.connection.Close()
}

func (r *rabbitMQExt) incrementRetryCountAndPrepareMessage(msg amqp.Delivery) amqp.Publishing {
	retryCount, ok := msg.Headers["x-delivery-count"].(int)
	if !ok {
		retryCount = 0
	}
	retryCount++

	headers := amqp.Table{
		"x-delivery-count": retryCount,
	}

	publishing := amqp.Publishing{
		ContentType: "text/plain",
		Body:        msg.Body,
		Headers:     headers,
	}

	return publishing
}
