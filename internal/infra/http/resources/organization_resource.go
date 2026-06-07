package resources

import "github.com/BohdanBoriak/boilerplate-go-back/internal/domain"

type OrganizationDto struct {
	Id          uint64  `json:"id"`
	UserId      uint64  `json:"userId"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	City        string  `json:"city"`
	Address     string  `json:"address"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
}

type OrganizationsDto struct {
	Items []OrganizationDto `json:"items"`
	Total uint64            `json:"total"`
	Pages uint              `json:"pages"`
}

func (d OrganizationDto) DomainToDto(org domain.Organization) OrganizationDto {
	return OrganizationDto{
		Id:          org.Id,
		UserId:      org.UserId,
		Name:        org.Name,
		Description: org.Description,
		City:        org.City,
		Address:     org.Address,
		Lat:         org.Lat,
		Lon:         org.Lon,
	}
}

func (d OrganizationDto) DomainToDtoCollection(orgs []domain.Organization) []OrganizationDto {
	result := make([]OrganizationDto, len(orgs))
	for i, o := range orgs {
		result[i] = d.DomainToDto(o)
	}
	return result
}
