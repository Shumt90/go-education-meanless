package config

import "github.com/confluentinc/confluent-kafka-go/kafka"

var AppKafkaProducer  *kafka.Producer
var AppKafkaConsumer  *kafka.Consumer
func KafkaInit()  {

	var err error

	AppKafkaProducer, err = kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": GetConfigByValue("kafka.host"),
		"client.id": GetConfigByValue("kafka.client.id"),
		"acks": "all"})

	if err!=nil {
		panic(err.Error())
	}

	AppKafkaConsumer, err = kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":    GetConfigByValue("kafka.host"),
		"group.id":             "app",
		"auto.offset.reset":    "smallest"})

	if err!=nil {
		panic(err.Error())
	}

	err = AppKafkaConsumer.SubscribeTopics([]string{GetConfigByValue("kafka.topic")}, nil)

	if err!=nil {
		panic(err.Error())
	}
}
