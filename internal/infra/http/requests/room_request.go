package requests

import "github.com/remuser309-gif/Golovan_Serghiy_OOP/internal/domain"

type RoomRequest struct {
	Name        string `json:"name" validate:"required,gte=1,max=100"`
	Description string `json:"description"`
}

func (r RoomRequest) ToDomainModel() (interface{}, error) {
	return domain.Room{
		Name:        r.Name,
		Description: r.Description,
	}, nil
}

type RoomUpdateRequest struct {
	Name        string `json:"name" validate:"required,gte=1,max=100"`
	Description string `json:"description"`
}

func (r RoomUpdateRequest) ToDomainModel() (interface{}, error) {
	return domain.Room{
		Name:        r.Name,
		Description: r.Description,
	}, nil
}
