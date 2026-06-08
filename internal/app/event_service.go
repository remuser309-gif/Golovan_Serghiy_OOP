package app

import (
	"github.com/remuser309-gif/Golovan_Serghiy_OOP/internal/domain"
	"github.com/remuser309-gif/Golovan_Serghiy_OOP/internal/infra/database"
	"log"
	"time"
)

type EventService interface {
	Save(e domain.Event) (domain.Event, error)
	Find(id uint64) (domain.Event, error)
	FindByDevice(deviceId uint64) ([]domain.Event, error)
	FindByRoom(roomId uint64) ([]domain.Event, error)
	FindByDeviceAndDateRange(deviceId uint64, from, to time.Time) ([]domain.Event, error)
	FindByAction(action string) ([]domain.Event, error)
	Update(e domain.Event) (domain.Event, error)
	Delete(id uint64) error
}

type eventService struct {
	eventRepo database.EventRepository
}

func NewEventService(er database.EventRepository) EventService {
	return eventService{
		eventRepo: er,
	}
}

func (s eventService) Save(e domain.Event) (domain.Event, error) {
	e, err := s.eventRepo.Save(e)
	if err != nil {
		log.Printf("EventService: %s", err)
		return domain.Event{}, err
	}
	return e, nil
}

func (s eventService) Find(id uint64) (domain.Event, error) {
	e, err := s.eventRepo.FindById(id)
	if err != nil {
		log.Printf("EventService: %s", err)
		return domain.Event{}, err
	}
	return e, nil
}

func (s eventService) FindByDevice(deviceId uint64) ([]domain.Event, error) {
	events, err := s.eventRepo.FindByDeviceId(deviceId)
	if err != nil {
		log.Printf("EventService: %s", err)
		return nil, err
	}
	return events, nil
}

func (s eventService) FindByRoom(roomId uint64) ([]domain.Event, error) {
	events, err := s.eventRepo.FindByRoomId(roomId)
	if err != nil {
		log.Printf("EventService: %s", err)
		return nil, err
	}
	return events, nil
}

func (s eventService) FindByDeviceAndDateRange(deviceId uint64, from, to time.Time) ([]domain.Event, error) {
	events, err := s.eventRepo.FindByDeviceIdAndDateRange(deviceId, from, to)
	if err != nil {
		log.Printf("EventService: %s", err)
		return nil, err
	}
	return events, nil
}

func (s eventService) FindByAction(action string) ([]domain.Event, error) {
	events, err := s.eventRepo.FindByAction(action)
	if err != nil {
		log.Printf("EventService: %s", err)
		return nil, err
	}
	return events, nil
}

func (s eventService) Update(e domain.Event) (domain.Event, error) {
	e, err := s.eventRepo.Update(e)
	if err != nil {
		log.Printf("EventService: %s", err)
		return domain.Event{}, err
	}
	return e, nil
}

func (s eventService) Delete(id uint64) error {
	err := s.eventRepo.Delete(id)
	if err != nil {
		log.Printf("EventService: %s", err)
		return err
	}
	return nil
}
