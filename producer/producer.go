package producer

import (
	"fmt"
	"log"
	"time"

	broker "github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/cmd"
	// To enable rabbitmq plugin uncomment
	"github.com/micro/go-plugins/broker/rabbitmq"
)

var (
	topic = "go.micro.srv"
	brokerPub = rabbitmq.NewBroker()
)

func Publish(item map[string]string) error {
	msg := &broker.Message{
		Header: item,
		Body: []byte(fmt.Sprintf("%s", time.Now().String())),
	}

	if err := brokerPub.Publish(topic, msg); err != nil {
		return err
	}

	return nil
}

func InitProducer() {
	cmd.Init()

	if err := brokerPub.Init(); err != nil {
		log.Fatalf("Broker Init error: %v", err)
	}

	if err := brokerPub.Connect(); err != nil {
		log.Fatalf("Broker Connect error: %v", err)
	}
}
