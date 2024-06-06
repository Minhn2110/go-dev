package main

import (
	"github.com/IBM/sarama"
	"log"
)

func main() {
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, nil)
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if err := consumer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	partitionConsumer, err := consumer.ConsumePartition("my_topic", 0, sarama.OffsetOldest)
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		if err := partitionConsumer.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	for {
		select {
		case msg := <-partitionConsumer.Messages():
			log.Printf("Received message from topic:%s partition:%d offset:%d message:%s\n", msg.Topic, msg.Partition, msg.Offset, string(msg.Value))
		}
	}
}
