package requests

import "github.com/remuser309-gif/Golovan_Serghiy_OOP/internal/domain"

type MeasurementRequest struct {
	DeviceId uint64  `json:"deviceId" validate:"required"`
	RoomId   uint64  `json:"roomId"`
	Value    float64 `json:"value" validate:"required"`
}

func (r MeasurementRequest) ToDomainModel() (interface{}, error) {
	return domain.Measurement{
		DeviceId: r.DeviceId,
		RoomId:   r.RoomId,
		Value:    r.Value,
	}, nil
}

type MeasurementUpdateRequest struct {
	DeviceId uint64  `json:"deviceId" validate:"required"`
	RoomId   uint64  `json:"roomId"`
	Value    float64 `json:"value" validate:"required"`
}

func (r MeasurementUpdateRequest) ToDomainModel() (interface{}, error) {
	return domain.Measurement{
		DeviceId: r.DeviceId,
		RoomId:   r.RoomId,
		Value:    r.Value,
	}, nil
}

type MeasurementFilterRequest struct {
	DeviceId uint64 `json:"deviceId"`
	RoomId   uint64 `json:"roomId"`
	From     string `json:"from"`
	To       string `json:"to"`
}
