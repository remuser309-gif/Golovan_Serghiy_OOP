package requests

import "github.com/remuser309-gif/Golovan_Serghiy_OOP/internal/domain"

type EventRequest struct {
	DeviceId uint64 `json:"deviceId" validate:"required"`
	RoomId   uint64 `json:"roomId"`
	Action   string `json:"action" validate:"required"`
}

func (r EventRequest) ToDomainModel() (interface{}, error) {
	return domain.Event{
		DeviceId: r.DeviceId,
		RoomId:   r.RoomId,
		Action:   r.Action,
	}, nil
}

type EventUpdateRequest struct {
	DeviceId uint64 `json:"deviceId" validate:"required"`
	RoomId   uint64 `json:"roomId"`
	Action   string `json:"action" validate:"required"`
}

func (r EventUpdateRequest) ToDomainModel() (interface{}, error) {
	return domain.Event{
		DeviceId: r.DeviceId,
		RoomId:   r.RoomId,
		Action:   r.Action,
	}, nil
}

type EventFilterRequest struct {
	DeviceId uint64 `json:"deviceId"`
	RoomId   uint64 `json:"roomId"`
	Action   string `json:"action"`
	From     string `json:"from"`
	To       string `json:"to"`
}
