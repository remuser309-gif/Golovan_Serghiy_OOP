package app

import (
	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/database"
	"log"
)

type DeviceService interface {
	Save(dev domain.Device) (domain.Device, error)
	Find(id uint64) (domain.Device, error)
	FindByOrg(orgId uint64) ([]domain.Device, error)
	FindByRoom(roomId uint64) ([]domain.Device, error)
	FindByCategory(cat domain.Category) ([]domain.Device, error)
	Update(dev domain.Device) (domain.Device, error)
	Delete(id uint64) error
}

type deviceService struct {
	devRepo database.DeviceRepository
}

func NewDeviceService(dr database.DeviceRepository) DeviceService {
	return deviceService{
		devRepo: dr,
	}
}

func (s deviceService) Save(dev domain.Device) (domain.Device, error) {
	dev, err := s.devRepo.Save(dev)
	if err != nil {
		log.Printf("DeviceService: %s", err)
		return domain.Device{}, err
	}
	return dev, nil
}

func (s deviceService) Find(id uint64) (domain.Device, error) {
	dev, err := s.devRepo.FindById(id)
	if err != nil {
		log.Printf("DeviceService: %s", err)
		return domain.Device{}, err
	}
	return dev, nil
}

func (s deviceService) FindByOrg(orgId uint64) ([]domain.Device, error) {
	devices, err := s.devRepo.FindByOrganizationId(orgId)
	if err != nil {
		log.Printf("DeviceService: %s", err)
		return nil, err
	}
	return devices, nil
}

func (s deviceService) FindByRoom(roomId uint64) ([]domain.Device, error) {
	devices, err := s.devRepo.FindByRoomId(roomId)
	if err != nil {
		log.Printf("DeviceService: %s", err)
		return nil, err
	}
	return devices, nil
}

func (s deviceService) FindByCategory(cat domain.Category) ([]domain.Device, error) {
	devices, err := s.devRepo.FindByCategory(cat)
	if err != nil {
		log.Printf("DeviceService: %s", err)
		return nil, err
	}
	return devices, nil
}

func (s deviceService) Update(dev domain.Device) (domain.Device, error) {
	dev, err := s.devRepo.Update(dev)
	if err != nil {
		log.Printf("DeviceService: %s", err)
		return domain.Device{}, err
	}
	return dev, nil
}

func (s deviceService) Delete(id uint64) error {
	err := s.devRepo.Delete(id)
	if err != nil {
		log.Printf("DeviceService: %s", err)
		return err
	}
	return nil
}
