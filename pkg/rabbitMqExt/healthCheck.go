package rabbitMqExt

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/paper-indonesia/pdk/v2/amqp"

	pdkLogger "github.com/paper-indonesia/pdk/v2/logger"
)

var (
	onceHc     = sync.Once{}
	channelHc  *amqp.Channel
	errInitHc  error
	consumerHc *amqp.Event

	healthCheckQueueName, healthCheckRoutingKey string
)

const (
	replyToQueueName     = "amq.rabbitmq.reply-to"
	queueHealthCheckName = "q.snap-core.health-check"
)

func (r *rabbitMQExt) HealthCheck(ctx context.Context) (bool, error) {
	onceHc.Do(func() {
		if channelHc, errInitHc = r.getChannel(); errInitHc != nil {
			return
		}

		if errInitHc = r.RunConsumerHealthCheck(context.Background(), channelHc, queueHealthCheckName); errInitHc != nil {
			return
		}

		consumerHc, errInitHc = channelHc.Consume(replyToQueueName, "", true, true, false, false, nil)
	})
	if errInitHc != nil {
		return false, errInitHc
	}

	requestBody := `{"message":"PING"}`

	// Publish health check message
	err := channelHc.Publish(
		"",                   // Exchange
		queueHealthCheckName, // Queue Name
		false,                // Mandatory
		false,                // Immediate
		amqp.Publishing{
			ContentType:   "application/json",
			Body:          []byte(requestBody),
			ReplyTo:       replyToQueueName,
			CorrelationId: uuid.NewString(),
		},
	)
	if err != nil {
		r.logger.Error(ctx, "publish health check message", pdkLogger.Error(err))
		return false, err
	}

	select {
	case msg := <-consumerHc.Delivery:
		return string(msg.Body) == requestBody, nil

	case <-time.After(6 * time.Second):
		return false, errors.New("health check timeout")
	}
}

func (r *rabbitMQExt) RunConsumerHealthCheck(ctx context.Context, ch *amqp.Channel, queueName string) error {
	if _, err := channelHc.QueueDeclare(queueName, true, false, false, false, nil); err != nil {
		return err
	}
	msgs, err := ch.ConsumeWithContext(ctx, queueName, "", true, false, false, false, nil)
	if err != nil {
		return err
	}
	go func() {
		for {
			select {
			case <-ctx.Done():
				return

			case msg := <-msgs.Delivery:
				if msg.Body == nil {
					continue
				}
				err := ch.Publish(
					"",          // Exchange
					msg.ReplyTo, // Queue Name
					false,       // Mandatory
					false,       // Immediate
					amqp.Publishing{
						Body:          msg.Body,
						ContentType:   msg.ContentType,
						CorrelationId: msg.CorrelationId,
					},
				)
				if err != nil {
					r.logger.Error(ctx, "Reply health check message", pdkLogger.Error(err))
				}
			}
		}
	}()
	return nil
}
