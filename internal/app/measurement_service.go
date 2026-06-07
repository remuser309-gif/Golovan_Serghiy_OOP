package app

import (
	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/database"
	"log"
	"time"
)

type MeasurementService interface {
	Save(m domain.Measurement) (domain.Measurement, error)
	Find(id uint64) (domain.Measurement, error)
	FindByDevice(deviceId uint64) ([]domain.Measurement, error)
	FindByRoom(roomId uint64) ([]domain.Measurement, error)
	FindByDeviceAndDateRange(deviceId uint64, from, to time.Time) ([]domain.Measurement, error)
	Update(m domain.Measurement) (domain.Measurement, error)
	Delete(id uint64) error
}

type measurementService struct {
	measRepo database.MeasurementRepository
}

func NewMeasurementService(mr database.MeasurementRepository) MeasurementService {
	return measurementService{
		measRepo: mr,
	}
}

func (s measurementService) Save(m domain.Measurement) (domain.Measurement, error) {
	m, err := s.measRepo.Save(m)
	if err != nil {
		log.Printf("MeasurementService: %s", err)
		return domain.Measurement{}, err
	}
	return m, nil
}

func (s measurementService) Find(id uint64) (domain.Measurement, error) {
	m, err := s.measRepo.FindById(id)
	if err != nil {
		log.Printf("MeasurementService: %s", err)
		return domain.Measurement{}, err
	}
	return m, nil
}

func (s measurementService) FindByDevice(deviceId uint64) ([]domain.Measurement, error) {
	measurements, err := s.measRepo.FindByDeviceId(deviceId)
	if err != nil {
		log.Printf("MeasurementService: %s", err)
		return nil, err
	}
	return measurements, nil
}

func (s measurementService) FindByRoom(roomId uint64) ([]domain.Measurement, error) {
	measurements, err := s.measRepo.FindByRoomId(roomId)
	if err != nil {
		log.Printf("MeasurementService: %s", err)
		return nil, err
	}
	return measurements, nil
}

func (s measurementService) FindByDeviceAndDateRange(deviceId uint64, from, to time.Time) ([]domain.Measurement, error) {
	measurements, err := s.measRepo.FindByDeviceIdAndDateRange(deviceId, from, to)
	if err != nil {
		log.Printf("MeasurementService: %s", err)
		return nil, err
	}
	return measurements, nil
}

func (s measurementService) Update(m domain.Measurement) (domain.Measurement, error) {
	m, err := s.measRepo.Update(m)
	if err != nil {
		log.Printf("MeasurementService: %s", err)
		return domain.Measurement{}, err
	}
	return m, nil
}

func (s measurementService) Delete(id uint64) error {
	err := s.measRepo.Delete(id)
	if err != nil {
		log.Printf("MeasurementService: %s", err)
		return err
	}
	return nil
}
