package consumer

import "log"

func Test(topic string) {

	consumer, err := NewConsumer(
		[]string{"127.0.0.1:9092"},
		WithConsumerErrors(true),
	)
	if err != nil {
		log.Fatalf("消费者连接错误->>>%v", err)
	}

	defer consumer.Client.Close()

	consumer.ReadMsg(topic)
}
