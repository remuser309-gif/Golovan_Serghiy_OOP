package app

import (
	"github.com/remuser309-gif/Golovan_Serghiy_OOP/internal/domain"
	"github.com/remuser309-gif/Golovan_Serghiy_OOP/internal/infra/database"
	"log"
)

type RoomService interface {
	Save(room domain.Room) (domain.Room, error)
	Find(id uint64) (domain.Room, error)
	FindByOrg(orgId uint64) ([]domain.Room, error)
	Update(room domain.Room) (domain.Room, error)
	Delete(id uint64) error
}

type roomService struct {
	roomRepo database.RoomRepository
}

func NewRoomService(rr database.RoomRepository) RoomService {
	return roomService{
		roomRepo: rr,
	}
}

func (s roomService) Save(room domain.Room) (domain.Room, error) {
	room, err := s.roomRepo.Save(room)
	if err != nil {
		log.Printf("RoomService: %s", err)
		return domain.Room{}, err
	}
	return room, nil
}

func (s roomService) Find(id uint64) (domain.Room, error) {
	room, err := s.roomRepo.FindById(id)
	if err != nil {
		log.Printf("RoomService: %s", err)
		return domain.Room{}, err
	}
	return room, nil
}

func (s roomService) FindByOrg(orgId uint64) ([]domain.Room, error) {
	rooms, err := s.roomRepo.FindByOrganizationId(orgId)
	if err != nil {
		log.Printf("RoomService: %s", err)
		return nil, err
	}
	return rooms, nil
}

func (s roomService) Update(room domain.Room) (domain.Room, error) {
	room, err := s.roomRepo.Update(room)
	if err != nil {
		log.Printf("RoomService: %s", err)
		return domain.Room{}, err
	}
	return room, nil
}

func (s roomService) Delete(id uint64) error {
	err := s.roomRepo.Delete(id)
	if err != nil {
		log.Printf("RoomService: %s", err)
		return err
	}
	return nil
}
