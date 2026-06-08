package database

import (
	"time"
	"github.com/remuser309-gif/Golovan_Serghiy_OOP/internal/domain"
	"github.com/google/uuid"
	"github.com/upper/db/v4"
)

const DevicesTableName = "devices"

type device struct {
	Id               uint64          `db:"id,omitempty"`
	OrganizationId   uint64          `db:"organization_id"`
	RoomId           uint64          `db:"room_id"`
	GUID             string          `db:"guid"`
	InventoryNumber  string          `db:"inventory_number"`
	SerialNumber     string          `db:"serial_number"`
	Characteristics  string          `db:"characteristics"`
	Category         domain.Category `db:"category"`
	Units            string          `db:"units"`
	PowerConsumption float64         `db:"power_consumption"`
	CreatedDate      time.Time       `db:"created_date"`
	UpdatedDate      time.Time       `db:"updated_date"`
	DeletedDate      *time.Time      `db:"deleted_date,omitempty"`
}

type DeviceRepository interface {
	Save(dev domain.Device) (domain.Device, error)
	FindById(id uint64) (domain.Device, error)
	FindByOrganizationId(orgId uint64) ([]domain.Device, error)
	FindByRoomId(roomId uint64) ([]domain.Device, error)
	FindByCategory(cat domain.Category) ([]domain.Device, error)
	Update(dev domain.Device) (domain.Device, error)
	Delete(id uint64) error
}

type deviceRepository struct {
	coll db.Collection
	sess db.Session
}

func NewDeviceRepository(dbSession db.Session) DeviceRepository {
	return deviceRepository{
		coll: dbSession.Collection(DevicesTableName),
		sess: dbSession,
	}
}

func (r deviceRepository) Save(dev domain.Device) (domain.Device, error) {
	d := r.mapDomainToModel(dev)
	d.CreatedDate, d.UpdatedDate = time.Now(), time.Now()
	err := r.coll.InsertReturning(&d)
	if err != nil {
		return domain.Device{}, err
	}
	return r.mapModelToDomain(d), nil
}

func (r deviceRepository) FindById(id uint64) (domain.Device, error) {
	var d device
	err := r.coll.Find(db.Cond{"id": id, "deleted_date": nil}).One(&d)
	if err != nil {
		return domain.Device{}, err
	}
	return r.mapModelToDomain(d), nil
}

func (r deviceRepository) FindByOrganizationId(orgId uint64) ([]domain.Device, error) {
	var devices []device
	err := r.coll.Find(db.Cond{"organization_id": orgId, "deleted_date": nil}).All(&devices)
	if err != nil {
		return nil, err
	}
	return r.mapModelsToDomain(devices), nil
}

func (r deviceRepository) FindByRoomId(roomId uint64) ([]domain.Device, error) {
	var devices []device
	err := r.coll.Find(db.Cond{"room_id": roomId, "deleted_date": nil}).All(&devices)
	if err != nil {
		return nil, err
	}
	return r.mapModelsToDomain(devices), nil
}

func (r deviceRepository) FindByCategory(cat domain.Category) ([]domain.Device, error) {
	var devices []device
	err := r.coll.Find(db.Cond{"category": cat, "deleted_date": nil}).All(&devices)
	if err != nil {
		return nil, err
	}
	return r.mapModelsToDomain(devices), nil
}

func (r deviceRepository) Update(dev domain.Device) (domain.Device, error) {
	d := r.mapDomainToModel(dev)
	d.UpdatedDate = time.Now()
	err := r.coll.Find(db.Cond{"id": d.Id, "deleted_date": nil}).Update(&d)
	if err != nil {
		return domain.Device{}, err
	}
	return r.mapModelToDomain(d), nil
}

func (r deviceRepository) Delete(id uint64) error {
	return r.coll.Find(db.Cond{"id": id, "deleted_date": nil}).Update(map[string]interface{}{"deleted_date": time.Now()})
}

func (r deviceRepository) mapDomainToModel(d domain.Device) device {
	return device{
		Id:               d.Id,
		OrganizationId:   d.OrganizationId,
		RoomId:           d.RoomId,
		GUID:             d.GUID.String(),
		InventoryNumber:  d.InventoryNumber,
		SerialNumber:     d.SerialNumber,
		Characteristics:  d.Characteristics,
		Category:         d.Category,
		Units:            d.Units,
		PowerConsumption: d.PowerConsumption,
		CreatedDate:      d.CreatedDate,
		UpdatedDate:      d.UpdatedDate,
		DeletedDate:      d.DeletedDate,
	}
}

func (r deviceRepository) mapModelToDomain(m device) domain.Device {
	guid, _ := uuid.Parse(m.GUID)
	return domain.Device{
		Id:               m.Id,
		OrganizationId:   m.OrganizationId,
		RoomId:           m.RoomId,
		GUID:             guid,
		InventoryNumber:  m.InventoryNumber,
		SerialNumber:     m.SerialNumber,
		Characteristics:  m.Characteristics,
		Category:         m.Category,
		Units:            m.Units,
		PowerConsumption: m.PowerConsumption,
		CreatedDate:      m.CreatedDate,
		UpdatedDate:      m.UpdatedDate,
		DeletedDate:      m.DeletedDate,
	}
}

func (r deviceRepository) mapModelsToDomain(models []device) []domain.Device {
	result := make([]domain.Device, len(models))
	for i, m := range models {
		result[i] = r.mapModelToDomain(m)
	}
	return result
}
