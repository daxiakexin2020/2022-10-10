package consumer

import (
	"github.com/Shopify/sarama"
)

type Consumer struct {
	Client sarama.Consumer
	Addr   []string       `json:"addr"`
	Config *sarama.Config `json:"config"`
}

type Coptions func(p *Consumer)

func WithConsumerErrors(isReturnErrors bool) Coptions {
	return func(c *Consumer) {
		c.Config.Consumer.Return.Errors = isReturnErrors
	}
}

func (c *Consumer) apply(opts []Coptions) {
	for _, opt := range opts {
		opt(c)
	}
}

func NewConsumer(addr []string, opts ...Coptions) (*Consumer, error) {
	c := &Consumer{Config: sarama.NewConfig()}
	c.apply(opts)
	client, err := sarama.NewConsumer(addr, c.Config)
	if err != nil {
		return nil, err
	}
	c.Client = client
	return c, nil
}
