package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"

	"github.com/goagile/kitchenservice/ticket/repo"

	"github.com/Shopify/sarama"
	"github.com/goagile/kitchenservice/command"
	"github.com/goagile/kitchenservice/event"
	"github.com/goagile/kitchenservice/service"
	"github.com/goagile/kitchenservice/ticket/repo/pg"
)

type Config struct {
	KafkaBroker          string
	KitchenCommandsTopic string
	OrderRepliesTopic    string
	DBHost               string
	DBPort               string
	DBName               string
	DBUser               string
	DBPassword           string
}

func GetConfig(path string) *Config {
	var conf *Config
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("GetConfig %v %v", path, err)
	}
	d := json.NewDecoder(file)
	err = d.Decode(&conf)
	if err != nil {
		log.Fatalf("GetConfig %v %v", path, err)
	}
	return conf
}

var (
	Conf             *Config
	KafkaBrokers     []string
	KafkaBrokersConf *sarama.Config
	TicketRepo       repo.TicketRepo
	DomainEvents     service.Publisher
	Kitchen          *service.Service
)

func init() {
	var confPath string
	flag.StringVar(&confPath, "c", "conf.json", "configuration file")
	flag.Parse()
	Conf = GetConfig(confPath)

	KafkaBrokers = []string{Conf.KafkaBroker}
	KafkaBrokersConf = sarama.NewConfig()
	KafkaBrokersConf.Consumer.Return.Errors = true
	KafkaBrokersConf.Producer.Return.Errors = true
	KafkaBrokersConf.Producer.Return.Successes = true

	pg.DB = pg.Connected(os.Getenv("KITCHEN_PG"))
	TicketRepo = pg.NewTicketRepo()
	DomainEvents = newKafkaPublisher()
	Kitchen = service.NewService(DomainEvents, Conf.OrderRepliesTopic, TicketRepo)
}

type KafkaMessageEnvelope struct {
	Name string          `json:"name"`
	Data json.RawMessage `json:"data"`
}

func main() {
	consumer, err := sarama.NewConsumer(KafkaBrokers, KafkaBrokersConf)
	if err != nil {
		log.Fatalf("NewConsumer: %v\n", err)
	}
	defer consumer.Close()
	listen(consumer, Conf.KitchenCommandsTopic, 0)
}

func listen(consumer sarama.Consumer, topic string, partition int32) {
	master, err := consumer.ConsumePartition(topic, partition, sarama.OffsetNewest)
	if err != nil {
		log.Fatalf("ConsumePartition: %v\n", err)
	}
	defer master.Close()

	log.Printf("Listen Kafka Topic %v ... \n", topic)

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

			case "create_ticket":
				var c command.CreateTicket
				err = json.Unmarshal(envelope.Data, &c)
				if err != nil {
					log.Printf("Unmarshal w Data: %v\n", err)
					continue
				}
				id, err := Kitchen.CreateTicket(service.TicketDetails{
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
				err = Kitchen.AcceptTicket(c.TicketID)
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
				err = Kitchen.Cancel(c.TicketID)
				if err != nil {
					log.Printf("AcceptTicket: %v\n", err)
					continue
				}
				log.Printf("Cancelled Ticket with id:%v\n", c.TicketID)
			}
		}
	}
}

type KafkaPublisher struct {
	Producer sarama.SyncProducer
}

func newKafkaPublisher() service.Publisher {
	producer, err := sarama.NewSyncProducer(KafkaBrokers, KafkaBrokersConf)
	if err != nil {
		log.Fatalf("NewSyncProducer: %v\n", err)
	}
	// defer producer.Close()
	return &KafkaPublisher{
		Producer: producer,
	}
}

type KafkaProducerMessageEnvelope struct {
	Name string      `json:"name"`
	Data event.Event `json:"data"`
}

//
// Publish
//
func (k *KafkaPublisher) Publish(topic string, e event.Event) {
	m := &KafkaProducerMessageEnvelope{
		Name: e.Name(),
		Data: e,
	}
	b, err := json.Marshal(m)
	if err != nil {
		log.Printf("Marshal %v\n", err)
	}
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(b),
	}
	_, _, err = k.Producer.SendMessage(msg)
	if err != nil {
		log.Printf("SendMessage %v\n", err)
	}
}
