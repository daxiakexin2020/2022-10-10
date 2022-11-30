package producer

import (
	"github.com/Shopify/sarama"
)

type Producer struct {
	Client sarama.SyncProducer
	Addr   []string       `json:"addr"`
	Config *sarama.Config `json:"config"`
}

type Poptions func(p *Producer)

func WithProducerRequiredAcks(requiredAcks sarama.RequiredAcks) Poptions {
	return func(p *Producer) {
		p.Config.Producer.RequiredAcks = requiredAcks
	}
}

func WithProducerPartitioner(partitioner sarama.PartitionerConstructor) Poptions {
	return func(p *Producer) {
		p.Config.Producer.Partitioner = partitioner
	}
}

func WithProducerSuccess(success bool) Poptions {
	return func(p *Producer) {
		p.Config.Producer.Return.Successes = success
	}
}

func (p *Producer) apply(opts []Poptions) {
	for _, opt := range opts {
		opt(p)
	}
}

func NewProducer(addr []string, opts ...Poptions) (*Producer, error) {
	p := &Producer{Config: sarama.NewConfig()}
	p.apply(opts)
	client, err := sarama.NewSyncProducer(addr, p.Config)
	if err != nil {
		return nil, err
	}
	p.Client = client
	return p, nil
}
