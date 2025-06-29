package rabbitMqExt

import (
	"context"
	"errors"

	"github.com/paper-indonesia/pdk/v2/amqp"
	pdkLogger "github.com/paper-indonesia/pdk/v2/logger"
)

// Setup iterate to all config and declare exchange, queue, and bind them.
func (r *rabbitMQExt) Setup(ctx context.Context) error {
	ch, err := r.getChannel()
	if err != nil {
		return err
	}
	defer ch.Close()

	for _, conf := range rabbitMqExchangeConfig {
		args := amqp.Table{}
		if conf.DLQExchange != "" {
			args = amqp.Table{
				"x-queue-type":              "quorum",
				"x-dead-letter-exchange":    conf.DLQExchange,
				"x-dead-letter-routing-key": conf.RoutingKey,
			}
		}

		if conf.Type == "x-delayed-message" {
			args["x-delayed-type"] = "direct"
		}

		if err := r.declare(ctx, ch, &conf, args); err != nil {
			return err
		}
		if conf.DLQExchange != "" {
			if err := r.declareDlq(ctx, ch, &conf, args); err != nil {
				return err
			}
		}
	}
	return nil
}

func (r *rabbitMQExt) declare(ctx context.Context, ch *amqp.Channel, conf *RabbitMqExchangeConfig, args amqp.Table) error {
	if conf.Exchange == "" {
		return errors.New("rabbitMqConfig.Exchange is empty")
	}

	err := ch.ExchangeDeclare(
		conf.Exchange, // exchange name
		conf.Type,     // type
		true,          // durable
		false,         // auto-deleted
		false,         // internal
		false,         // no-wait
		args,          // arguments
	)
	if err != nil {
		r.logger.Error(ctx, "error rabbitmq when exchange declare", pdkLogger.Error(err))
		return err
	}

	_, err = ch.QueueDeclare(
		conf.QueueName, // Queue name
		true,           // Durable (the queue will not survive server restarts)
		false,          // Delete when unused
		false,          // Exclusive (queue only accessible by the connection that declares it)
		false,          // No-wait
		args,           // Arguments
	)
	if err != nil {
		r.logger.Error(ctx, "error rabbitmq when queue declare", pdkLogger.Error(err))
		return err
	}

	err = ch.QueueBind(
		conf.QueueName,  // Queue name
		conf.RoutingKey, // Routing key
		conf.Exchange,   // Exchange
		false,
		nil,
	)
	if err != nil {
		r.logger.Error(ctx, "error rabbitmq when queue bind", pdkLogger.Error(err))
		return err
	}

	return nil
}

func (r *rabbitMQExt) declareDlq(ctx context.Context, ch *amqp.Channel, conf *RabbitMqExchangeConfig, args amqp.Table) error {
	err := ch.ExchangeDeclare(
		conf.DLQExchange, // name
		conf.Type,        // type
		true,             // durable
		false,            // auto-deleted
		false,            // internal
		false,            // no-wait
		args,             // arguments
	)
	if err != nil {
		r.logger.Error(ctx, "error rabbitmq when DLQ exchange declare", pdkLogger.Error(err))
		return err
	}

	_, err = ch.QueueDeclare(
		conf.DLQQueueName, // Queue name
		true,              // Durable (the queue will not survive server restarts)
		false,             // Delete when unused
		false,             // Exclusive (queue only accessible by the connection that declares it)
		false,             // No-wait
		nil,               // Arguments
	)
	if err != nil {
		r.logger.Error(ctx, "error rabbitmq when DLQ queue declare", pdkLogger.Error(err))
		return err
	}

	err = ch.QueueBind(
		conf.DLQQueueName, // Queue name
		conf.RoutingKey,   // Routing key
		conf.DLQExchange,  // Exchange
		false,
		nil,
	)
	if err != nil {
		r.logger.Error(ctx, "error rabbitmq when DLQ queue bind", pdkLogger.Error(err))
		return err
	}

	return nil
}
