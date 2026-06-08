package database

import (
	"time"
	"github.com/remuser309-gif/Golovan_Serghiy_OOP/internal/domain"
	"github.com/upper/db/v4"
)

const MeasurementsTableName = "measurements"

type measurement struct {
	Id          uint64     `db:"id,omitempty"`
	DeviceId    uint64     `db:"device_id"`
	RoomId      uint64     `db:"room_id"`
	Value       float64    `db:"value"`
	CreatedDate time.Time  `db:"created_date"`
	UpdatedDate time.Time  `db:"updated_date"`
	DeletedDate *time.Time `db:"deleted_date,omitempty"`
}

type MeasurementRepository interface {
	Save(m domain.Measurement) (domain.Measurement, error)
	FindById(id uint64) (domain.Measurement, error)
	FindByDeviceId(deviceId uint64) ([]domain.Measurement, error)
	FindByRoomId(roomId uint64) ([]domain.Measurement, error)
	FindByDeviceIdAndDateRange(deviceId uint64, from, to time.Time) ([]domain.Measurement, error)
	Update(m domain.Measurement) (domain.Measurement, error)
	Delete(id uint64) error
}

type measurementRepository struct {
	coll db.Collection
	sess db.Session
}

func NewMeasurementRepository(dbSession db.Session) MeasurementRepository {
	return measurementRepository{
		coll: dbSession.Collection(MeasurementsTableName),
		sess: dbSession,
	}
}

func (r measurementRepository) Save(m domain.Measurement) (domain.Measurement, error) {
	mm := r.mapDomainToModel(m)
	mm.CreatedDate, mm.UpdatedDate = time.Now(), time.Now()
	err := r.coll.InsertReturning(&mm)
	if err != nil {
		return domain.Measurement{}, err
	}
	return r.mapModelToDomain(mm), nil
}

func (r measurementRepository) FindById(id uint64) (domain.Measurement, error) {
	var mm measurement
	err := r.coll.Find(db.Cond{"id": id, "deleted_date": nil}).One(&mm)
	if err != nil {
		return domain.Measurement{}, err
	}
	return r.mapModelToDomain(mm), nil
}

func (r measurementRepository) FindByDeviceId(deviceId uint64) ([]domain.Measurement, error) {
	var measurements []measurement
	err := r.coll.Find(db.Cond{"device_id": deviceId, "deleted_date": nil}).All(&measurements)
	if err != nil {
		return nil, err
	}
	return r.mapModelsToDomain(measurements), nil
}

func (r measurementRepository) FindByRoomId(roomId uint64) ([]domain.Measurement, error) {
	var measurements []measurement
	err := r.coll.Find(db.Cond{"room_id": roomId, "deleted_date": nil}).All(&measurements)
	if err != nil {
		return nil, err
	}
	return r.mapModelsToDomain(measurements), nil
}

func (r measurementRepository) FindByDeviceIdAndDateRange(deviceId uint64, from, to time.Time) ([]domain.Measurement, error) {
	var measurements []measurement
	err := r.coll.Find(
		db.Cond{"device_id": deviceId, "deleted_date": nil},
		db.Cond{"created_date >=": from},
		db.Cond{"created_date <=": to},
	).All(&measurements)
	if err != nil {
		return nil, err
	}
	return r.mapModelsToDomain(measurements), nil
}

func (r measurementRepository) Update(m domain.Measurement) (domain.Measurement, error) {
	mm := r.mapDomainToModel(m)
	mm.UpdatedDate = time.Now()
	err := r.coll.Find(db.Cond{"id": mm.Id, "deleted_date": nil}).Update(&mm)
	if err != nil {
		return domain.Measurement{}, err
	}
	return r.mapModelToDomain(mm), nil
}

func (r measurementRepository) Delete(id uint64) error {
	return r.coll.Find(db.Cond{"id": id, "deleted_date": nil}).Update(map[string]interface{}{"deleted_date": time.Now()})
}

func (r measurementRepository) mapDomainToModel(d domain.Measurement) measurement {
	return measurement{
		Id:          d.Id,
		DeviceId:    d.DeviceId,
		RoomId:      d.RoomId,
		Value:       d.Value,
		CreatedDate: d.CreatedDate,
		UpdatedDate: d.UpdatedDate,
		DeletedDate: d.DeletedDate,
	}
}

func (r measurementRepository) mapModelToDomain(m measurement) domain.Measurement {
	return domain.Measurement{
		Id:          m.Id,
		DeviceId:    m.DeviceId,
		RoomId:      m.RoomId,
		Value:       m.Value,
		CreatedDate: m.CreatedDate,
		UpdatedDate: m.UpdatedDate,
		DeletedDate: m.DeletedDate,
	}
}

func (r measurementRepository) mapModelsToDomain(models []measurement) []domain.Measurement {
	result := make([]domain.Measurement, len(models))
	for i, m := range models {
		result[i] = r.mapModelToDomain(m)
	}
	return result
}
