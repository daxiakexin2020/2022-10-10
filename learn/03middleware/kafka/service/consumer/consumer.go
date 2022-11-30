package consumer

import (
	"fmt"
	"github.com/Shopify/sarama"
	"sync/atomic"
)

func (c *Consumer) ReadMsg(topic string) {

	//todo 主题->分区，依据主题，遍历所有分区数据
	var total int64
	partitionList, err := c.Client.Partitions(topic) // 根据topic取到所有的分区
	if err != nil {
		fmt.Printf("fail to get list of partition:err%v\n", err)
		return
	}
	for partition := range partitionList {
		partitionConsumer, err := c.Client.ConsumePartition(topic, int32(partition), sarama.OffsetOldest)
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
