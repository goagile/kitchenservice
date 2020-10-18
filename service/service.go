package service

import (
	"github.com/goagile/kitchenservice/event"
	"github.com/goagile/kitchenservice/ticket"
	"github.com/goagile/kitchenservice/ticket/repo"
)

type Publisher interface {
	Publish(topic string, e event.Event)
}

//
// NewService
//
func NewService(events Publisher, orderRepliesTopic string, ticketRepo repo.TicketRepo) *Service {
	return &Service{
		DomainEvents:      events,
		TicketRepo:        ticketRepo,
		OrderRepliesTopic: orderRepliesTopic,
	}
}

//
// Service
//
type Service struct {
	DomainEvents      Publisher
	OrderRepliesTopic string
	TicketRepo        repo.TicketRepo
}

//
// TicketDetails
//
type TicketDetails struct {
	OrderID int64 `json:"order_id"`
}

//
// CreateTicket
//
func (s *Service) CreateTicket(details TicketDetails) (ticket.TicketID, error) {
	id, err := s.TicketRepo.NextID()
	if err != nil {
		return id, err
	}

	tic := ticket.New(id)
	tic.OrderID = details.OrderID
	if err := s.TicketRepo.Save(tic); err != nil {
		return id, err
	}

	s.DomainEvents.Publish(s.OrderRepliesTopic, &event.TicketCreated{
		TicketID: tic.ID,
		OrderID:  tic.OrderID,
	})

	return id, nil
}

//
// AcceptTicket
//
func (s *Service) AcceptTicket(id ticket.TicketID) error {
	tic, err := s.TicketRepo.Find(id)
	if err != nil {
		return err
	}

	err = tic.Accept()
	if err != nil {
		return err
	}

	err = s.TicketRepo.Save(tic)
	if err != nil {
		return err
	}

	s.DomainEvents.Publish(s.OrderRepliesTopic, &event.TicketAccepted{
		TicketID: tic.ID,
	})

	return nil
}

//
// PrepareTicket
//
func (s *Service) PrepareTicket(id ticket.TicketID) error {
	tic, err := s.TicketRepo.Find(id)
	if err != nil {
		return err
	}

	err = tic.Prepare()
	if err != nil {
		return err
	}

	err = s.TicketRepo.Save(tic)
	if err != nil {
		return err
	}

	s.DomainEvents.Publish(s.OrderRepliesTopic, &event.TicketPrepared{
		TicketID: tic.ID,
	})

	return nil
}

//
// ReadyToPickUp
//
func (s *Service) ReadyToPickUp(id ticket.TicketID) error {
	tic, err := s.TicketRepo.Find(id)
	if err != nil {
		return err
	}

	err = tic.ReadyToPickUp()
	if err != nil {
		return err
	}

	err = s.TicketRepo.Save(tic)
	if err != nil {
		return err
	}

	s.DomainEvents.Publish(s.OrderRepliesTopic, &event.TicketReadyToPickUp{
		TicketID: tic.ID,
	})

	return nil
}

//
// Cancel
//
func (s *Service) Cancel(id ticket.TicketID) error {
	tic, err := s.TicketRepo.Find(id)
	if err != nil {
		return err
	}

	err = tic.Cancel()
	if err != nil {
		return err
	}

	err = s.TicketRepo.Save(tic)
	if err != nil {
		return err
	}

	s.DomainEvents.Publish(s.OrderRepliesTopic, &event.TicketCancelled{
		TicketID: tic.ID,
	})

	return nil
}
