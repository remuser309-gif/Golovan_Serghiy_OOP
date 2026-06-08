package requests

import "github.com/remuser309-gif/Golovan_Serghiy_OOP/internal/domain"

type OrganizationRequest struct {
	Name        string  `json:"name" validate:"required,gte=1,max=100"`
	Description string  `json:"description"`
	City        string  `json:"city" validate:"required,gte=1,max=100"`
	Address     string  `json:"address" validate:"required,gte=1,max=200"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
}

func (r OrganizationRequest) ToDomainModel() (interface{}, error) {
	return domain.Organization{
		Name:        r.Name,
		Description: r.Description,
		City:        r.City,
		Address:     r.Address,
		Lat:         r.Lat,
		Lon:         r.Lon,
	}, nil
}

type OrganizationUpdateRequest struct {
	Name        string  `json:"name" validate:"required,gte=1,max=100"`
	Description string  `json:"description"`
	City        string  `json:"city" validate:"required,gte=1,max=100"`
	Address     string  `json:"address" validate:"required,gte=1,max=200"`
	Lat         float64 `json:"lat"`
	Lon         float64 `json:"lon"`
}

func (r OrganizationUpdateRequest) ToDomainModel() (interface{}, error) {
	return domain.Organization{
		Name:        r.Name,
		Description: r.Description,
		City:        r.City,
		Address:     r.Address,
		Lat:         r.Lat,
		Lon:         r.Lon,
	}, nil
}
