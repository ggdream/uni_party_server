package mq

import (
	"fmt"
	"github.com/Shopify/sarama"
)

func NewKafka()  {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true

	msg := &sarama.ProducerMessage{
		Topic: "test",
		Value: sarama.StringEncoder("hello"),
	}

	client, err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config)
	if err != nil {
		panic(err)
	}
	defer func(client sarama.SyncProducer) {
		err := client.Close()
		if err != nil {
			panic(err)
		}
	}(client)

	pid, offset, err := client.SendMessage(msg)
	if err != nil {
		panic(err)
	}
	fmt.Println(pid, offset)
}
