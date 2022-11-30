package service

import (
	"fmt"
	"github.com/Shopify/sarama"
	"sync"
	"sync/atomic"
)

func DemoProduct() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认  ,follow从leader中同步数据，成功以后，发送ACK至leader，leader发送消息至生产者
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出⼀个partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回
	// sasl认证
	//config.Net.SASL.Enable = true
	//config.Net.SASL.User = "admin"
	//config.Net.SASL.Password = "admin"
	// 连接kafka
	client, err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config)
	if err != nil {
		fmt.Println("producer close, err:", err)
		return
	}
	defer client.Close()

	// 构造⼀个消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = "test"
	msg.Value = sarama.StringEncoder("hello kafka")
	// 发送消息
	var wg sync.WaitGroup
	limit := 100
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < limit; i++ {
			pid, offset, err := client.SendMessage(msg)
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

func DemoCunsumer() {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	//config.Net.SASL.Enable = true
	//config.Net.SASL.User = "admin"
	//config.Net.SASL.Password = "admin"
	// consumer
	consumer, err := sarama.NewConsumer([]string{"127.0.0.1:9092"}, config)
	if err != nil {
		fmt.Printf("consumer_test create consumer error %s\n", err.Error())
		return
	}

	//todo 主题->分区，依据主题，遍历所有分区数据
	var total int64
	partitionList, err := consumer.Partitions("test") // 根据topic取到所有的分区
	if err != nil {
		fmt.Printf("fail to get list of partition:err%v\n", err)
		return
	}
	defer consumer.Close()
	for partition := range partitionList {
		partitionConsumer, err := consumer.ConsumePartition("test", int32(partition), sarama.OffsetOldest)
		if err != nil {
			fmt.Printf("try create partition_consumer error %s\n", err.Error())
			return
		}
		defer partitionConsumer.Close()
		for {
			select {
			case msg := <-partitionConsumer.Messages():
				atomic.AddInt64(&total, 1)
				fmt.Printf("msg offset: %d, partition: %d, timestamp: %s, value: %s, 总数：%d\n",
					msg.Offset, msg.Partition, msg.Timestamp.String(), string(msg.Value), total)
			case err := <-partitionConsumer.Errors():
				fmt.Printf("err :%s\n", err.Error())
			}
		}
	}
}
