package database

import (
	"time"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/upper/db/v4"
)

const OrganizationsTableName = "organizations"

type organization struct {
	Id          uint64     `db:"id,omitempty"`
	UserId      uint64     `db:"user_id"`
	Name        string     `db:"name"`
	Description string     `db:"description"`
	City        string     `db:"city"`
	Address     string     `db:"address"`
	Lat         float64    `db:"lat"`
	Lon         float64    `db:"lon"`
	CreatedDate time.Time  `db:"created_date"`
	UpdatedDate time.Time  `db:"updated_date"`
	DeletedDate *time.Time `db:"deleted_date,omitempty"`
}

type OrganizationRepository interface {
	Save(org domain.Organization) (domain.Organization, error)
	FindById(id uint64) (domain.Organization, error)
	FindByUserId(userId uint64) ([]domain.Organization, error)
	Update(org domain.Organization) (domain.Organization, error)
	Delete(id uint64) error
}

type organizationRepository struct {
	coll db.Collection
	sess db.Session
}

func NewOrganizationRepository(dbSession db.Session) OrganizationRepository {
	return organizationRepository{
		coll: dbSession.Collection(OrganizationsTableName),
		sess: dbSession,
	}
}

func (r organizationRepository) Save(org domain.Organization) (domain.Organization, error) {
	o := r.mapDomainToModel(org)
	o.CreatedDate, o.UpdatedDate = time.Now(), time.Now()
	err := r.coll.InsertReturning(&o)
	if err != nil {
		return domain.Organization{}, err
	}
	return r.mapModelToDomain(o), nil
}

func (r organizationRepository) FindById(id uint64) (domain.Organization, error) {
	var o organization
	err := r.coll.Find(db.Cond{"id": id, "deleted_date": nil}).One(&o)
	if err != nil {
		return domain.Organization{}, err
	}
	return r.mapModelToDomain(o), nil
}

func (r organizationRepository) FindByUserId(userId uint64) ([]domain.Organization, error) {
	var orgs []organization
	err := r.coll.Find(db.Cond{"user_id": userId, "deleted_date": nil}).All(&orgs)
	if err != nil {
		return nil, err
	}
	return r.mapModelsToDomain(orgs), nil
}

func (r organizationRepository) Update(org domain.Organization) (domain.Organization, error) {
	o := r.mapDomainToModel(org)
	o.UpdatedDate = time.Now()
	err := r.coll.Find(db.Cond{"id": o.Id, "deleted_date": nil}).Update(&o)
	if err != nil {
		return domain.Organization{}, err
	}
	return r.mapModelToDomain(o), nil
}

func (r organizationRepository) Delete(id uint64) error {
	return r.coll.Find(db.Cond{"id": id, "deleted_date": nil}).Update(map[string]interface{}{"deleted_date": time.Now()})
}

func (r organizationRepository) mapDomainToModel(d domain.Organization) organization {
	return organization{
		Id:          d.Id,
		UserId:      d.UserId,
		Name:        d.Name,
		Description: d.Description,
		City:        d.City,
		Address:     d.Address,
		Lat:         d.Lat,
		Lon:         d.Lon,
		CreatedDate: d.CreatedDate,
		UpdatedDate: d.UpdatedDate,
		DeletedDate: d.DeletedDate,
	}
}

func (r organizationRepository) mapModelToDomain(m organization) domain.Organization {
	return domain.Organization{
		Id:          m.Id,
		UserId:      m.UserId,
		Name:        m.Name,
		Description: m.Description,
		City:        m.City,
		Address:     m.Address,
		Lat:         m.Lat,
		Lon:         m.Lon,
		CreatedDate: m.CreatedDate,
		UpdatedDate: m.UpdatedDate,
		DeletedDate: m.DeletedDate,
	}
}

func (r organizationRepository) mapModelsToDomain(models []organization) []domain.Organization {
	result := make([]domain.Organization, len(models))
	for i, m := range models {
		result[i] = r.mapModelToDomain(m)
	}
	return result
}
