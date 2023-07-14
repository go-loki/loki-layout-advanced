package consumer

import (
	"context"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/go-hasaki/hasaki-layout-advanced/pkg/xkafka"
)

type Client interface {
	Start(ctx context.Context) error
	Stop() error
}

type HandlerFunc func(message *kafka.Message)

type Config struct {
	Topics           []string                     `json:"topics"`
	GroupId          string                       `json:"group.id"`
	BootstrapServers string                       `json:"bootstrap.servers"`
	SecurityProtocol string                       `json:"security.protocol"`
	SaslMechanism    string                       `json:"sasl.mechanism"`
	SaslUsername     string                       `json:"sasl.username"`
	SaslPassword     string                       `json:"sasl.password"`
	SslCaLocation    string                       `json:"ssl.ca.location"`
	ConfigMap        map[string]kafka.ConfigValue `json:"config.map"`
}

func NewConsumer(cfg Config, opts ...Option) Client {
	hlog.Info("init kafka consumer, it may take a few seconds to init the connection")
	// common arguments
	var kafkaConf = &kafka.ConfigMap{
		"api.version.request":       "true",
		"auto.offset.reset":         "latest",
		"heartbeat.interval.ms":     3000,
		"session.timeout.ms":        30000,
		"max.poll.interval.ms":      120000,
		"fetch.max.bytes":           1024000,
		"max.partition.fetch.bytes": 256000,
	}
	if cfg.ConfigMap != nil {
		for k, v := range cfg.ConfigMap {
			_ = kafkaConf.SetKey(k, v)
		}
	}
	_ = kafkaConf.SetKey("bootstrap.servers", cfg.BootstrapServers)
	_ = kafkaConf.SetKey("group.id", cfg.GroupId)
	switch cfg.SecurityProtocol {
	case "plaintext":
		_ = kafkaConf.SetKey("security.protocol", cfg.SecurityProtocol)
	case "sasl_ssl":
		_ = kafkaConf.SetKey("security.protocol", cfg.SecurityProtocol)
		_ = kafkaConf.SetKey("ssl.ca.location", cfg.SslCaLocation)
		_ = kafkaConf.SetKey("sasl.username", cfg.SaslUsername)
		_ = kafkaConf.SetKey("sasl.password", cfg.SaslPassword)
		_ = kafkaConf.SetKey("sasl.mechanism", cfg.SaslMechanism)
	case "sasl_plaintext":
		_ = kafkaConf.SetKey("security.protocol", cfg.SecurityProtocol)
		_ = kafkaConf.SetKey("sasl.username", cfg.SaslUsername)
		_ = kafkaConf.SetKey("sasl.password", cfg.SaslPassword)
		_ = kafkaConf.SetKey("sasl.mechanism", cfg.SaslMechanism)
	default:
		panic(kafka.NewError(kafka.ErrUnknownProtocol, "unknown protocol", true))
	}
	consumer, err := kafka.NewConsumer(kafkaConf)
	if err != nil {
		panic(err)
	}
	hlog.Info("init kafka consumer success")
	cli := &client{consumer: consumer}
	cli.opts = Options{
		Topics: cfg.Topics,
	}
	for _, o := range opts {
		o(&cli.opts)
	}
	cli.notifyClose = make(chan error)
	return cli
}

type client struct {
	opts        Options
	consumer    *kafka.Consumer
	notifyClose chan error
}

func (cli *client) Start(ctx context.Context) error {
	return xkafka.RunWithContext(ctx, cli.start)
}

func (cli *client) start() error {
	err := cli.consumer.SubscribeTopics(cli.opts.Topics, cli.opts.RebalanceCb)
	if err != nil {
		return err
	}
	for {
		ev := cli.consumer.Poll(-1)
		switch e := ev.(type) {
		case *kafka.Message:
			if e.TopicPartition.Error != nil {
				hlog.Errorf("Consumer error: %+v (%+v)\n", e.TopicPartition.Error, e)
			}
			cli.opts.Handler(e)
		case kafka.Error:
			hlog.Errorf("Consumer error: %+v\n", e)
		default:
			// Ignore other event types
		}
		select {
		case err = <-cli.notifyClose:
			return err
		default:
			// Ignore
		}
	}
}

func (cli *client) Stop() error {
	close(cli.notifyClose)
	return cli.consumer.Close()
}
