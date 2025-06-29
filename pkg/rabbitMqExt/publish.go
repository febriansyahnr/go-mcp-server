package rabbitMqExt

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/google/uuid"
	"github.com/paper-indonesia/pdk/v2/amqp"
	pdkConst "github.com/paper-indonesia/pdk/v2/constant"
	pdkLogger "github.com/paper-indonesia/pdk/v2/logger"
	"github.com/paper-indonesia/pg-mcp-server/constant"
	"github.com/paper-indonesia/pg-mcp-server/pkg/util"
)

func (r *rabbitMQExt) Publish(
	ctx context.Context,
	message interface{},
	document string,
	failedMsg ...amqp.Delivery,
) error {
	rabbitMqConfig := setRabbitMqConfig(document)
	if rabbitMqConfig.Exchange == "" {
		return errors.New("rabbitMqConfig.Exchange is empty")
	}

	ch, err := r.getChannel()
	if err != nil {
		return err
	}
	defer ch.Close()

	var body []byte
	if _, ok := message.([]byte); !ok {
		jsonData, err := json.Marshal(message)
		if err != nil {
			r.logger.Error(ctx, "error rabbitmq when marshaling message", pdkLogger.Error(err))
			return err
		}

		body = []byte(jsonData)
	} else {
		body = message.([]byte)
	}

	// Here you use the function to prepare the message for retry
	publishing := amqp.Publishing{
		ContentType: "text/plain",
		Body:        body,
	}

	// when feature with scheduled queue need implement more, we
	// need implement 1 functions more
	exchangesWithDelay := []string{
		DOCUMENT_TYPE_VA_STATUS_SCHEDULED,
		DOCUMENT_TYPE_QRIS_STATUS_SCHEDULED,
		DOCUMENT_TYPE_CHECK_TRANSFER_STATUS,
		DOCUMENT_TYPE_CHECK_QRIS_ISSUING_STATUS,
	}

	if util.InArray(document, exchangesWithDelay) {
		header, err := r.setHeaderScheduledQueue(ctx)
		if err != nil || header == nil {
			r.logger.Error(ctx, "error rabbitmq when set header for scheduler message", pdkLogger.Error(err))
			return err
		}

		publishing.Headers = *header
	}

	traceID, ok := ctx.Value(pdkConst.CtxTraceIdKey).(string)
	if !ok {
		traceID = uuid.NewString()
	}

	if publishing.Headers == nil {
		publishing.Headers = make(map[string]interface{})
	}

	publishing.Headers["trace-id"] = traceID

	if len(failedMsg) > 0 {
		publishing = r.incrementRetryCountAndPrepareMessage(failedMsg[0])
	}

	r.logger.Info(ctx, "publishing message", pdkLogger.Any("header", publishing.Headers))

	err = ch.PublishWithContext(
		ctx,
		rabbitMqConfig.Exchange,   // Exchange (empty string means the default exchange)
		rabbitMqConfig.RoutingKey, // Routing key (queue name is used as the routing key here)
		false,                     // Mandatory
		false,                     // Immediate
		publishing,
	)
	if err != nil {
		r.logger.Error(ctx, "error rabbitmq when publish", pdkLogger.Error(err))
		return err
	}

	return nil
}

func (r *rabbitMQExt) PublishAndWaitReply(ctx context.Context, document string, message interface{}) (event *amqp.Event, close io.Closer, err error) {

	ch, err := r.getChannel()
	if err != nil {
		return nil, nil, fmt.Errorf("Get Channel: %w", err)
	}
	defer func() {
		if err != nil {
			_ = ch.Close()
		}
	}()

	cfg := setRabbitMqConfig(document)

	if _, ok := r.once.directReplyTo.Load(document); !ok {

		_, err = ch.QueueDeclare(cfg.QueueName, true, false, false, false, amqp.Table{
			"x-queue-type": amqp.QueueTypeClassic,
		})
		if err != nil {
			return nil, nil, fmt.Errorf("Queue Declare: %w", err)
		}
		r.once.directReplyTo.Store(document, cfg)
	}

	buf := pool.Get().(*bytes.Buffer)
	defer func() {
		buf.Reset()
		pool.Put(buf)
	}()

	contentType := "text/plain"
	switch data := message.(type) {
	case []byte:
		_, _ = buf.Write(data)

	default:
		_ = json.NewEncoder(buf).Encode(data)
		contentType = constant.MIMEApplicationJSON
	}

	_ = ch.Qos(1, 0, false)

	msgs, err := ch.Consume(ReplyToQueueName, "", true, true, false, false, nil)
	if err != nil {
		return nil, nil, fmt.Errorf("Wait Reply: %w", err)
	}

	err = ch.PublishWithContext(
		ctx,
		"",            // Default Exchange
		cfg.QueueName, // Queue Name
		false,         // Mandatory
		false,         // Immediate
		amqp.Publishing{
			ContentType:   contentType,
			Body:          buf.Bytes(),
			ReplyTo:       ReplyToQueueName,
			CorrelationId: uuid.NewString(),
		},
	)
	if err != nil {
		return nil, nil, fmt.Errorf("Publish Message: %w", err)
	}
	return msgs, ch, nil
}

func (r *rabbitMQExt) PublishToReplyQueue(ctx context.Context, replyToAddress string, payload amqp.Publishing) error {
	ch, err := r.getChannel()
	if err != nil {
		return fmt.Errorf("Get Channel: %w", err)
	}
	defer ch.Close()

	return ch.PublishWithContext(ctx, "", replyToAddress, false, false, payload)
}

func (r *rabbitMQExt) setHeaderScheduledQueue(ctx context.Context) (*amqp.Table, error) {
	delayTime, ok := ctx.Value(CTX_SCHEDULED_DELAY).(time.Duration)
	if !ok || delayTime == 0 {
		return nil, errors.New("delay time is empty")
	}

	return &amqp.Table{
		"x-delay": delayTime.Milliseconds(),
	}, nil
}
