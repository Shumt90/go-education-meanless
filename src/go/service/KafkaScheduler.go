package service
import (
	"awesomeProject/src/go/config"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
	"os"
	"time"
)
func ReadUpdateSignal()  {
	for {
		ev := config.AppKafkaConsumer.Poll(0)
		switch e := ev.(type) {
		case *kafka.Message:
			UpdateRootPageData()
		case kafka.PartitionEOF:
			fmt.Printf("%% Reached %v\n", e)
		case kafka.Error:
			fmt.Fprintf(os.Stderr, "%% Error: %v\n", e)
		default:
		}


		time.Sleep(2 * time.Second)
	}
}

func PushUpdateSignal()  {

	delivery_chan := make(chan kafka.Event, 10000)

	topic := config.GetConfigByValue("kafka.topic")

	err := config.AppKafkaProducer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value: []byte("no matter")},
		delivery_chan,
	)

	e := <-delivery_chan
	m := e.(*kafka.Message)

	if m.TopicPartition.Error != nil {
		fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
	} else {
		fmt.Printf("Delivered message to topic %s [%d] at offset %v\n",
			*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
	}

	if err!=nil {
		log.Panicln(err.Error())
	}
}
