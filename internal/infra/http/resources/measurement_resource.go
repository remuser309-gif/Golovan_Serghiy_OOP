package resources

import "github.com/BohdanBoriak/boilerplate-go-back/internal/domain"

type MeasurementDto struct {
	Id       uint64  `json:"id"`
	DeviceId uint64  `json:"deviceId"`
	RoomId   uint64  `json:"roomId"`
	Value    float64 `json:"value"`
}

type MeasurementsDto struct {
	Items []MeasurementDto `json:"items"`
	Total uint64           `json:"total"`
	Pages uint             `json:"pages"`
}

func (d MeasurementDto) DomainToDto(m domain.Measurement) MeasurementDto {
	return MeasurementDto{
		Id:       m.Id,
		DeviceId: m.DeviceId,
		RoomId:   m.RoomId,
		Value:    m.Value,
	}
}

func (d MeasurementDto) DomainToDtoCollection(measurements []domain.Measurement) []MeasurementDto {
	result := make([]MeasurementDto, len(measurements))
	for i, m := range measurements {
		result[i] = d.DomainToDto(m)
	}
	return result
}
