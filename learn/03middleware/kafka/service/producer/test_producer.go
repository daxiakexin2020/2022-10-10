package producer

import (
	"github.com/Shopify/sarama"
	"log"
)

func Test(topic string, msg string) {
	product, err := NewProducer(
		[]string{"127.0.0.1:9092"},
		WithProducerRequiredAcks(sarama.WaitForAll),
		WithProducerPartitioner(sarama.NewRandomPartitioner),
		WithProducerSuccess(true),
	)
	if err != nil {
		log.Fatalf("生产者连接错误->>>%v", err)
	}

	defer product.Client.Close()

	product.SendMsg(topic, msg)
}
