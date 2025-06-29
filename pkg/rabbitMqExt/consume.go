package rabbitMqExt

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/paper-indonesia/pdk/v2/amqp"
	pdkConst "github.com/paper-indonesia/pdk/v2/constant"
	pdkLogger "github.com/paper-indonesia/pdk/v2/logger"
	"github.com/paper-indonesia/pg-mcp-server/constant"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

var otelTracer = otel.Tracer("Consumer")

func (r *rabbitMQExt) Consume(signal context.Context, document string, process func(context.Context, []byte) error) error {
	rabbitMqConfig := setRabbitMqConfig(document)
	if rabbitMqConfig.Exchange == "" {
		return errors.New("rabbitMqConfig.Exchange is empty")
	}

	ch, err := r.getChannel()
	if err != nil {
		return err
	}

	messages, err := ch.ConsumeWithContext(
		signal,                   // Signal to stop the message consumption process
		rabbitMqConfig.QueueName, // queue
		"",                       // consumer tag
		false,                    // auto-acknowledge
		false,                    // exclusive
		false,                    // no-local
		false,                    // no-wait
		nil,                      // args
	)
	if err != nil {
		r.logger.Error(signal, "error rabbitmq when consume", pdkLogger.Error(err))
		return err
	}

	defer ch.Close()
	defer ch.CloseConsume(messages)

	for {
		select {
		case <-signal.Done():
			return nil

		case msg := <-messages.Delivery:
			if msg.Body == nil {
				continue
			}

			r.processMessage(document, rabbitMqConfig, msg, process)
		}
	}
}

func (r *rabbitMQExt) processMessage(document string, rabbitMqConfig RabbitMqExchangeConfig, msg amqp.Delivery, process func(context.Context, []byte) error) {
	ctx, span := otelTracer.Start(context.Background(), "Consumer "+rabbitMqConfig.QueueName, trace.WithSpanKind(trace.SpanKindConsumer))
	defer span.End()

	txn := r.nrApp.GetApp().StartTransaction("Consumer " + rabbitMqConfig.QueueName)
	defer txn.End()

	txn.AddAttribute(newrelic.AttributeMessageSystem, "rabbitmq")
	txn.AddAttribute(newrelic.AttributeMessageCorrelationID, msg.CorrelationId)
	txn.AddAttribute(newrelic.AttributeMessageHeaders, msg.Headers)
	txn.AddAttribute(newrelic.AttributeMessageExchangeType, rabbitMqConfig.Exchange)
	txn.AddAttribute(newrelic.AttributeMessageQueueName, rabbitMqConfig.QueueName)
	txn.AddAttribute(newrelic.AttributeMessageRoutingKey, rabbitMqConfig.RoutingKey)

	ctx = newrelic.NewContext(ctx, txn)

	attrs := []attribute.KeyValue{
		attribute.String("exchange", rabbitMqConfig.Exchange),
		attribute.String("queue", rabbitMqConfig.QueueName),
		attribute.String("routing", rabbitMqConfig.RoutingKey),
	}

	// Get traceID
	traceID, ok := msg.Headers[HeaderTraceId].(string)
	if !ok {
		traceID = uuid.NewString()
	}
	// Set the transaction in the context
	ctx = context.WithValue(ctx, constant.CtxRabbitMQStartTime, time.Now().UTC())
	ctx = context.WithValue(ctx, pdkConst.CtxTraceIdKey, traceID)

	r.logger.Info(ctx, "consume message",
		pdkLogger.String("queueName", rabbitMqConfig.QueueName),
		pdkLogger.String("publishedAt", msg.Timestamp.Format(time.RFC3339)),
	)

	// Increment retry count and republish the message
	retryCount, ok := msg.Headers["x-retry-count"].(int)
	if !ok {
		retryCount = 0
	}

	if retryCount > constant.MaxRetryMechanism {
		return
	}

	// Get ReplyTo when exists
	if msg.ReplyTo != "" {
		ctx = context.WithValue(ctx, constant.CtxRabbitMqReplyTo, msg.ReplyTo)
	}

	if err := process(ctx, msg.Body); err != nil {
		r.logger.Error(ctx, "error rabbitmq when processing message", pdkLogger.Error(err))

		span.RecordError(err)
		span.SetStatus(codes.Error, "An error occurred")

		if err = r.Publish(ctx, msg.Body, document, msg); err != nil {
			r.logger.Error(ctx, "error rabbitmq when retry publishing message", pdkLogger.Error(err))
		}

		attrs = append(attrs, attribute.Int64("delivery_tag", int64(msg.DeliveryTag)))
		attrs = append(attrs, attribute.Bool("ack", false))
		span.SetAttributes(attrs...)

		msg.Nack(false, true)
		return
	}

	attrs = append(attrs, attribute.Bool("ack", true))
	span.SetAttributes(attrs...)

	msg.Ack(false)
}
