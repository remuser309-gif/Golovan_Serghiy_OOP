package database

import (
	"time"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/upper/db/v4"
)

const EventsTableName = "events"

type event struct {
	Id          uint64     `db:"id,omitempty"`
	DeviceId    uint64     `db:"device_id"`
	RoomId      uint64     `db:"room_id"`
	Action      string     `db:"action"`
	CreatedDate time.Time  `db:"created_date"`
	UpdatedDate time.Time  `db:"updated_date"`
	DeletedDate *time.Time `db:"deleted_date,omitempty"`
}

type EventRepository interface {
	Save(e domain.Event) (domain.Event, error)
	FindById(id uint64) (domain.Event, error)
	FindByDeviceId(deviceId uint64) ([]domain.Event, error)
	FindByRoomId(roomId uint64) ([]domain.Event, error)
	FindByDeviceIdAndDateRange(deviceId uint64, from, to time.Time) ([]domain.Event, error)
	FindByAction(action string) ([]domain.Event, error)
	Update(e domain.Event) (domain.Event, error)
	Delete(id uint64) error
}

type eventRepository struct {
	coll db.Collection
	sess db.Session
}

func NewEventRepository(dbSession db.Session) EventRepository {
	return eventRepository{
		coll: dbSession.Collection(EventsTableName),
		sess: dbSession,
	}
}

func (r eventRepository) Save(e domain.Event) (domain.Event, error) {
	ev := r.mapDomainToModel(e)
	ev.CreatedDate, ev.UpdatedDate = time.Now(), time.Now()
	err := r.coll.InsertReturning(&ev)
	if err != nil {
		return domain.Event{}, err
	}
	return r.mapModelToDomain(ev), nil
}

func (r eventRepository) FindById(id uint64) (domain.Event, error) {
	var ev event
	err := r.coll.Find(db.Cond{"id": id, "deleted_date": nil}).One(&ev)
	if err != nil {
		return domain.Event{}, err
	}
	return r.mapModelToDomain(ev), nil
}

func (r eventRepository) FindByDeviceId(deviceId uint64) ([]domain.Event, error) {
	var events []event
	err := r.coll.Find(db.Cond{"device_id": deviceId, "deleted_date": nil}).All(&events)
	if err != nil {
		return nil, err
	}
	return r.mapModelsToDomain(events), nil
}

func (r eventRepository) FindByRoomId(roomId uint64) ([]domain.Event, error) {
	var events []event
	err := r.coll.Find(db.Cond{"room_id": roomId, "deleted_date": nil}).All(&events)
	if err != nil {
		return nil, err
	}
	return r.mapModelsToDomain(events), nil
}

func (r eventRepository) FindByDeviceIdAndDateRange(deviceId uint64, from, to time.Time) ([]domain.Event, error) {
	var events []event
	err := r.coll.Find(
		db.Cond{"device_id": deviceId, "deleted_date": nil},
		db.Cond{"created_date >=": from},
		db.Cond{"created_date <=": to},
	).All(&events)
	if err != nil {
		return nil, err
	}
	return r.mapModelsToDomain(events), nil
}

func (r eventRepository) FindByAction(action string) ([]domain.Event, error) {
	var events []event
	err := r.coll.Find(db.Cond{"action": action, "deleted_date": nil}).All(&events)
	if err != nil {
		return nil, err
	}
	return r.mapModelsToDomain(events), nil
}

func (r eventRepository) Update(e domain.Event) (domain.Event, error) {
	ev := r.mapDomainToModel(e)
	ev.UpdatedDate = time.Now()
	err := r.coll.Find(db.Cond{"id": ev.Id, "deleted_date": nil}).Update(&ev)
	if err != nil {
		return domain.Event{}, err
	}
	return r.mapModelToDomain(ev), nil
}

func (r eventRepository) Delete(id uint64) error {
	return r.coll.Find(db.Cond{"id": id, "deleted_date": nil}).Update(map[string]interface{}{"deleted_date": time.Now()})
}

func (r eventRepository) mapDomainToModel(d domain.Event) event {
	return event{
		Id:          d.Id,
		DeviceId:    d.DeviceId,
		RoomId:      d.RoomId,
		Action:      d.Action,
		CreatedDate: d.CreatedDate,
		UpdatedDate: d.UpdatedDate,
		DeletedDate: d.DeletedDate,
	}
}

func (r eventRepository) mapModelToDomain(m event) domain.Event {
	return domain.Event{
		Id:          m.Id,
		DeviceId:    m.DeviceId,
		RoomId:      m.RoomId,
		Action:      m.Action,
		CreatedDate: m.CreatedDate,
		UpdatedDate: m.UpdatedDate,
		DeletedDate: m.DeletedDate,
	}
}

func (r eventRepository) mapModelsToDomain(models []event) []domain.Event {
	result := make([]domain.Event, len(models))
	for i, m := range models {
		result[i] = r.mapModelToDomain(m)
	}
	return result
}
