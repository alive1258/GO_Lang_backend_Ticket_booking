package booking

import (
	"goticket/internal/booking/dto"
	"goticket/internal/event"
	"uuid"
)

type service struct {
	bookingRepo Repository
	eventRepo   event.Repository
}

func NewService(bookingRepo Repository, eventRepo event.Repository) *service {
	return &service{
		bookingRepo: bookingRepo,
		eventRepo:   eventRepo,
	}
}

func generateBookingCode() string {
	return "GT-" + uuid.New().String()
}

func (s *service) CreateBooking(userId uint, req dto.CreateRequest) (*dto.Response, error) {
	booking, err := s.bookingRepo.CreateWithTicketUpdate(userId, req.EventID, req.Quantity)
	if err != nil {
		return nil, err
	}

	return booking.ToResponse(), nil
}