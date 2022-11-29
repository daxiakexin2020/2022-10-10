package service

import "github.com/Shopify/sarama"

type Product struct {
	Client *sarama.SyncProducer
	Addr   []string       `json:"addr"`
	Config *sarama.Config `json:"config"`
}

type Options func(p *Product)

func WithRequiredAcks(requiredAcks sarama.RequiredAcks) Options {
	return func(p *Product) {
		p.Config.Producer.RequiredAcks = requiredAcks
	}
}

func (p *Product) WithPartitioner(partitioner sarama.Partitioner) Options {
	return func(p *Product) {
		p.Config.Producer.Partitioner = partitioner
	}
}

func NewOption(opts ...Options) *Product {
	p := &Product{}
	for _, opt := range opts {
		opt(p)
	}
	client, err := sarama.NewSyncProducer([]string{}, p.Config)
	return p
}
