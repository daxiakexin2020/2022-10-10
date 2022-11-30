package producer

import (
	"fmt"
	"github.com/Shopify/sarama"
	"sync"
)

func (p *Producer) SendMsg(topic string, smsg string) {

	// 构造⼀个消息
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(smsg),
	}
	// 发送消息
	var wg sync.WaitGroup
	limit := 100
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < limit; i++ {
			pid, offset, err := p.Client.SendMessage(msg)
			if err != nil {
				fmt.Println("send message failed,", err)
				return
			}
			fmt.Printf("pid:%v offset:%v\n", pid, offset)
		}
	}()
	wg.Wait()
	fmt.Printf("发送%d条消息ok\n", limit)
}
