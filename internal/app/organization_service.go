package app

import (
	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/database"
	"log"
)

type OrganizationService interface {
	Save(org domain.Organization) (domain.Organization, error)
	Find(id uint64) (domain.Organization, error)
	FindByUser(userId uint64) ([]domain.Organization, error)
	Update(org domain.Organization) (domain.Organization, error)
	Delete(id uint64) error
}

type organizationService struct {
	orgRepo database.OrganizationRepository
}

func NewOrganizationService(or database.OrganizationRepository) OrganizationService {
	return organizationService{
		orgRepo: or,
	}
}

func (s organizationService) Save(org domain.Organization) (domain.Organization, error) {
	org, err := s.orgRepo.Save(org)
	if err != nil {
		log.Printf("OrganizationService: %s", err)
		return domain.Organization{}, err
	}
	return org, nil
}

func (s organizationService) Find(id uint64) (domain.Organization, error) {
	org, err := s.orgRepo.FindById(id)
	if err != nil {
		log.Printf("OrganizationService: %s", err)
		return domain.Organization{}, err
	}
	return org, nil
}

func (s organizationService) FindByUser(userId uint64) ([]domain.Organization, error) {
	orgs, err := s.orgRepo.FindByUserId(userId)
	if err != nil {
		log.Printf("OrganizationService: %s", err)
		return nil, err
	}
	return orgs, nil
}

func (s organizationService) Update(org domain.Organization) (domain.Organization, error) {
	org, err := s.orgRepo.Update(org)
	if err != nil {
		log.Printf("OrganizationService: %s", err)
		return domain.Organization{}, err
	}
	return org, nil
}

func (s organizationService) Delete(id uint64) error {
	err := s.orgRepo.Delete(id)
	if err != nil {
		log.Printf("OrganizationService: %s", err)
		return err
	}
	return nil
}
