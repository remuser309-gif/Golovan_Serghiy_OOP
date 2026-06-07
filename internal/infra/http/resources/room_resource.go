package resources

import "github.com/BohdanBoriak/boilerplate-go-back/internal/domain"

type RoomDto struct {
	Id             uint64 `json:"id"`
	OrganizationId uint64 `json:"organizationId"`
	Name           string `json:"name"`
	Description    string `json:"description"`
}

type RoomsDto struct {
	Items []RoomDto `json:"items"`
	Total uint64    `json:"total"`
	Pages uint      `json:"pages"`
}

func (d RoomDto) DomainToDto(room domain.Room) RoomDto {
	return RoomDto{
		Id:             room.Id,
		OrganizationId: room.OrganizationId,
		Name:           room.Name,
		Description:    room.Description,
	}
}

func (d RoomDto) DomainToDtoCollection(rooms []domain.Room) []RoomDto {
	result := make([]RoomDto, len(rooms))
	for i, r := range rooms {
		result[i] = d.DomainToDto(r)
	}
	return result
}
