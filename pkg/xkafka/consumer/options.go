package consumer

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Options struct {
	Topics      []string
	Handler     HandlerFunc
	RebalanceCb kafka.RebalanceCb
	IsDebug     bool
}

type Option func(opts *Options)

func Handler(handler HandlerFunc) Option {
	return func(opts *Options) {
		opts.Handler = func(msg *kafka.Message) {
			handler(msg)
		}
	}
}

func RebalanceCb(rebalanceCb kafka.RebalanceCb) Option {
	return func(opts *Options) {
		opts.RebalanceCb = rebalanceCb
	}
}

func IsDebug(isDebug bool) Option {
	return func(opts *Options) {
		opts.IsDebug = isDebug
	}
}
