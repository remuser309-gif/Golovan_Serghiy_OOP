package resources

import "github.com/BohdanBoriak/boilerplate-go-back/internal/domain"

type EventDto struct {
	Id       uint64 `json:"id"`
	DeviceId uint64 `json:"deviceId"`
	RoomId   uint64 `json:"roomId"`
	Action   string `json:"action"`
}

type EventsDto struct {
	Items []EventDto `json:"items"`
	Total uint64     `json:"total"`
	Pages uint       `json:"pages"`
}

func (d EventDto) DomainToDto(e domain.Event) EventDto {
	return EventDto{
		Id:       e.Id,
		DeviceId: e.DeviceId,
		RoomId:   e.RoomId,
		Action:   e.Action,
	}
}

func (d EventDto) DomainToDtoCollection(events []domain.Event) []EventDto {
	result := make([]EventDto, len(events))
	for i, e := range events {
		result[i] = d.DomainToDto(e)
	}
	return result
}
