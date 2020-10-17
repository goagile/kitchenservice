package kafka

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Shopify/sarama"
)

var (
	Brokers         = []string{"localhost:9092"}
	Partition int32 = 0
	Offset          = sarama.OffsetNewest
	Conf            = sarama.NewConfig()
)

//
// listener
//
type listener struct{}

//
// Listen
//
func (kl *listener) Listen(topic string) error {
	cons, err := sarama.NewConsumer(Brokers, Conf)
	if err != nil {
		return fmt.Errorf("NewConsumer: %v\n", err)
	}
	defer cons.Close()

	master, err := cons.ConsumePartition(topic, Partition, Offset)
	if err != nil {
		return fmt.Errorf("ConsumePartition: %v\n", err)
	}
	defer master.Close()

	for {
		select {
		case err := <-master.Errors():
			log.Printf("Consume err: %v\n", err)
			continue
		case m := <-master.Messages():
			// log.Printf("raw m:%+v\n", m)
			c, err := MessageDecoder.Decode(m.Value)
			if err != nil {
				log.Printf("Decode:%v\n", err)
				continue
			}
			log.Printf("c:%+v\n", c)
		}
	}
}

//
// Message
//
type Message interface {
	Name() string
}

var EventListener = &listener{}

//
// Listen
//
func Listen(topic string) error {
	return EventListener.Listen(topic)
}

var MessageDecoder Decoder

type Decoder interface {
	Decode([]byte) (Message, error)
}

//
// CommandMap
//
type CommandMap map[string]Message

//
// NewJSONDecoder
//
func NewJSONDecoder(cmds ...Message) *decoder {
	m := make(CommandMap, len(cmds))
	for _, c := range cmds {
		m[c.Name()] = c
	}
	return &decoder{m: m}
}

type decoder struct {
	m CommandMap
}

type rawCommand struct {
	Name string          `json:"name"`
	Data json.RawMessage `json:"data"`
}

//
// Decode
//
func (d *decoder) Decode(data []byte) (Message, error) {
	var r rawCommand
	if err := json.Unmarshal(data, &r); err != nil {
		return nil, fmt.Errorf("Unmarshal w Name: %v\n", err)
	}

	c, ok := d.m[r.Name]
	if !ok {
		return nil, fmt.Errorf("Not Found Command: %v", r.Name)
	}

	if err := json.Unmarshal(r.Data, &c); err != nil {
		return nil, fmt.Errorf("Unmarshal w Data: %v\n", err)
	}

	return c, nil
}
