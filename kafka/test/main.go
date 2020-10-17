package main

import (
	"log"

	"github.com/goagile/kitchenservice/event"

	"github.com/goagile/kitchenservice/command"
	"github.com/goagile/kitchenservice/kafka"
)

func main() {

	kafka.MessageDecoder = kafka.NewJSONDecoder(
		&command.CreateTicket{},
		&command.AcceptTicket{},
		&event.TicketCancelled{},
	)

	if err := kafka.Listen("test"); err != nil {
		log.Fatal(err)
	}
}
