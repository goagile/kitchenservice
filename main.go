package main

import (
	"encoding/json"
	"log"

	"github.com/Shopify/sarama"
	"github.com/goagile/kitchenservice/command"
	"github.com/goagile/kitchenservice/service"
)

var (
	KafkaBrokers     = []string{"127.0.0.1:9092"}
	KafkaBrokersConf *sarama.Config
)

func init() {
	KafkaBrokersConf = sarama.NewConfig()
	KafkaBrokersConf.Consumer.Return.Errors = true
}

type KafkaMessageEnvelope struct {
	Name string          `json:"name"`
	Data json.RawMessage `json:"data"`
}

func main() {
	//
	// KafkaConsume
	//
	consumer, err := sarama.NewConsumer(KafkaBrokers, KafkaBrokersConf)
	if err != nil {
		log.Fatalf("NewConsumer: %v\n", err)
	}
	defer consumer.Close()

	listen(consumer, "test", 0)
}

func listen(consumer sarama.Consumer, topic string, partition int32) {
	master, err := consumer.ConsumePartition(topic, partition, sarama.OffsetNewest)
	if err != nil {
		log.Fatalf("ConsumePartition: %v\n", err)
	}
	defer master.Close()

	log.Println("Listen Kafka ...")

	for {
		select {

		case err := <-master.Errors():
			log.Fatalf("Consumer:%v\n", err)

		case m := <-master.Messages():
			var envelope KafkaMessageEnvelope
			err := json.Unmarshal(m.Value, &envelope)
			if err != nil {
				log.Printf("KafkaMessageEnvelope Unmarshal:%v\n", err)
				continue
			}
			switch envelope.Name {

			default:
				log.Printf("Unknown message name:%v\n", envelope.Name)
				continue

			case "create_ticket":
				var c command.CreateTicket
				err = json.Unmarshal(envelope.Data, &c)
				if err != nil {
					log.Printf("Unmarshal w Data: %v\n", err)
					continue
				}
				id, err := service.CreateTicket(service.TicketDetails{
					OrderID: c.OrderID,
				})
				if err != nil {
					log.Printf("CreateTicket: %v\n", err)
					continue
				}
				log.Printf("Created Ticket with id:%v\n", id)

			case "accept_ticket":
				var c command.AcceptTicket
				err = json.Unmarshal(envelope.Data, &c)
				if err != nil {
					log.Printf("Unmarshal w Data: %v\n", err)
					continue
				}
				err = service.AcceptTicket(c.TicketID)
				if err != nil {
					log.Printf("AcceptTicket: %v\n", err)
					continue
				}
				log.Printf("Accepted Ticket with id:%v\n", c.TicketID)

			case "cancel_ticket":
				var c command.CancelTicket
				err = json.Unmarshal(envelope.Data, &c)
				if err != nil {
					log.Printf("Unmarshal w Data: %v\n", err)
					continue
				}
				err = service.Cancel(c.TicketID)
				if err != nil {
					log.Printf("AcceptTicket: %v\n", err)
					continue
				}
				log.Printf("Cancelled Ticket with id:%v\n", c.TicketID)
			}
		}
	}
}
