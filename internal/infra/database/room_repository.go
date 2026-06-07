package database

import (
	"time"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/upper/db/v4"
)

const RoomsTableName = "rooms"

type room struct {
	Id             uint64     `db:"id,omitempty"`
	OrganizationId uint64     `db:"organization_id"`
	Name           string     `db:"name"`
	Description    string     `db:"description"`
	CreatedDate    time.Time  `db:"created_date"`
	UpdatedDate    time.Time  `db:"updated_date"`
	DeletedDate    *time.Time `db:"deleted_date,omitempty"`
}

type RoomRepository interface {
	Save(room domain.Room) (domain.Room, error)
	FindById(id uint64) (domain.Room, error)
	FindByOrganizationId(orgId uint64) ([]domain.Room, error)
	Update(room domain.Room) (domain.Room, error)
	Delete(id uint64) error
}

type roomRepository struct {
	coll db.Collection
	sess db.Session
}

func NewRoomRepository(dbSession db.Session) RoomRepository {
	return roomRepository{
		coll: dbSession.Collection(RoomsTableName),
		sess: dbSession,
	}
}

func (r roomRepository) Save(room domain.Room) (domain.Room, error) {
	rm := r.mapDomainToModel(room)
	rm.CreatedDate, rm.UpdatedDate = time.Now(), time.Now()
	err := r.coll.InsertReturning(&rm)
	if err != nil {
		return domain.Room{}, err
	}
	return r.mapModelToDomain(rm), nil
}

func (r roomRepository) FindById(id uint64) (domain.Room, error) {
	var rm room
	err := r.coll.Find(db.Cond{"id": id, "deleted_date": nil}).One(&rm)
	if err != nil {
		return domain.Room{}, err
	}
	return r.mapModelToDomain(rm), nil
}

func (r roomRepository) FindByOrganizationId(orgId uint64) ([]domain.Room, error) {
	var rooms []room
	err := r.coll.Find(db.Cond{"organization_id": orgId, "deleted_date": nil}).All(&rooms)
	if err != nil {
		return nil, err
	}
	return r.mapModelsToDomain(rooms), nil
}

func (r roomRepository) Update(room domain.Room) (domain.Room, error) {
	rm := r.mapDomainToModel(room)
	rm.UpdatedDate = time.Now()
	err := r.coll.Find(db.Cond{"id": rm.Id, "deleted_date": nil}).Update(&rm)
	if err != nil {
		return domain.Room{}, err
	}
	return r.mapModelToDomain(rm), nil
}

func (r roomRepository) Delete(id uint64) error {
	return r.coll.Find(db.Cond{"id": id, "deleted_date": nil}).Update(map[string]interface{}{"deleted_date": time.Now()})
}

func (r roomRepository) mapDomainToModel(d domain.Room) room {
	return room{
		Id:             d.Id,
		OrganizationId: d.OrganizationId,
		Name:           d.Name,
		Description:    d.Description,
		CreatedDate:    d.CreatedDate,
		UpdatedDate:    d.UpdatedDate,
		DeletedDate:    d.DeletedDate,
	}
}

func (r roomRepository) mapModelToDomain(m room) domain.Room {
	return domain.Room{
		Id:             m.Id,
		OrganizationId: m.OrganizationId,
		Name:           m.Name,
		Description:    m.Description,
		CreatedDate:    m.CreatedDate,
		UpdatedDate:    m.UpdatedDate,
		DeletedDate:    m.DeletedDate,
	}
}

func (r roomRepository) mapModelsToDomain(models []room) []domain.Room {
	result := make([]domain.Room, len(models))
	for i, m := range models {
		result[i] = r.mapModelToDomain(m)
	}
	return result
}
