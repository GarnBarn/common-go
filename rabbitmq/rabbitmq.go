package rabbitmq

import (
	"fmt"
	"strconv"

	"github.com/sirupsen/logrus"
	"github.com/wagslane/go-rabbitmq"
)

type ConsumerConfig struct {
	ConnectionString   string
	MaxRetry           int
	DeadLetterExchange string
	FailoverExchange   string
	ConsumeQueue       string
}

type Processor interface {
	Process(d rabbitmq.Delivery) error
}

type rabbitMQ struct {
	conn      *rabbitmq.Conn
	publisher *rabbitmq.Publisher
}

type RabbitMQ interface {
	Consume(processor Processor, consumerConfig ConsumerConfig) (*rabbitmq.Consumer, error)
	GetPublisher() *rabbitmq.Publisher
	CloseConnection()
}

func NewRabbitMQ(connectionString string) (RabbitMQ, error) {
	conn, err := rabbitmq.NewConn(
		connectionString,
		rabbitmq.WithConnectionOptionsLogging,
	)
	if err != nil {
		logrus.Error("RabbitMQ Error: ", err)
		return &rabbitMQ{}, err
	}

	publisher, err := rabbitmq.NewPublisher(
		conn,
		rabbitmq.WithPublisherOptionsLogging,
	)
	if err != nil {
		logrus.Error("RabbitMQ Error: ", err)
		return &rabbitMQ{}, err
	}

	return &rabbitMQ{
		conn:      conn,
		publisher: publisher,
	}, nil
}

func (r *rabbitMQ) CloseConnection() {
	r.publisher.Close()
	r.conn.Close()
}

func (r *rabbitMQ) GetPublisher() *rabbitmq.Publisher {
	return r.publisher
}

func (r *rabbitMQ) Consume(processor Processor, consumerConfig ConsumerConfig) (*rabbitmq.Consumer, error) {
	return rabbitmq.NewConsumer(
		r.conn,
		func(d rabbitmq.Delivery) (action rabbitmq.Action) {
			logrus.Info("Start Processing the message")
			defer logrus.Info("End Processing the message")
			err := processor.Process(d)

			if err != nil {
				value, ok := d.Headers["x-retry"]
				retryCount := 0
				if ok {
					convertResult, err := strconv.Atoi(fmt.Sprint(value))
					if err != nil {
						convertResult = 0
					}
					retryCount = convertResult
				}

				retryCount++

				rabbitMqHeaderTable := rabbitmq.Table{
					"x-retry": retryCount,
				}

				if retryCount >= consumerConfig.MaxRetry {
					logrus.Warn("Maximum retry exceeded, Publishing the message to dead lettering exchange")
					r.publisher.Publish(d.Body, []string{d.RoutingKey},
						rabbitmq.WithPublishOptionsExchange(consumerConfig.DeadLetterExchange),
						rabbitmq.WithPublishOptionsContentType(d.ContentType),
						rabbitmq.WithPublishOptionsHeaders(rabbitMqHeaderTable),
					)
					return rabbitmq.Ack
				}

				logrus.Warn("Publishing message back to exchange, ")
				r.publisher.Publish(d.Body, []string{d.RoutingKey},
					rabbitmq.WithPublishOptionsExchange(consumerConfig.FailoverExchange),
					rabbitmq.WithPublishOptionsContentType(d.ContentType),
					rabbitmq.WithPublishOptionsHeaders(rabbitMqHeaderTable),
				)
			}

			return rabbitmq.Ack
		},
		consumerConfig.ConsumeQueue,
		rabbitmq.WithConsumerOptionsQueueDurable,
	)
}
