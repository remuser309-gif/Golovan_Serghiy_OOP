package requests

import (
	"github.com/remuser309-gif/Golovan_Serghiy_OOP/internal/domain"
	"github.com/google/uuid"
)

type DeviceRequest struct {
	OrganizationId  uint64  `json:"organizationId" validate:"required"`
	RoomId          uint64  `json:"roomId" validate:"required"`
	InventoryNumber string  `json:"inventoryNumber"`
	SerialNumber    string  `json:"serialNumber"`
	Characteristics string  `json:"characteristics"`
	Category        string  `json:"category" validate:"required,oneof=SENSOR ACTUATOR"`
	Units           string  `json:"units"`
	PowerConsumption float64 `json:"powerConsumption"`
}

func (r DeviceRequest) ToDomainModel() (interface{}, error) {
	return domain.Device{
		OrganizationId:   r.OrganizationId,
		RoomId:           r.RoomId,
		GUID:             uuid.New(),
		InventoryNumber:  r.InventoryNumber,
		SerialNumber:     r.SerialNumber,
		Characteristics:  r.Characteristics,
		Category:         domain.Category(r.Category),
		Units:            r.Units,
		PowerConsumption: r.PowerConsumption,
	}, nil
}

type DeviceUpdateRequest struct {
	OrganizationId  uint64  `json:"organizationId" validate:"required"`
	RoomId          uint64  `json:"roomId" validate:"required"`
	InventoryNumber string  `json:"inventoryNumber"`
	SerialNumber    string  `json:"serialNumber"`
	Characteristics string  `json:"characteristics"`
	Category        string  `json:"category" validate:"required,oneof=SENSOR ACTUATOR"`
	Units           string  `json:"units"`
	PowerConsumption float64 `json:"powerConsumption"`
}

func (r DeviceUpdateRequest) ToDomainModel() (interface{}, error) {
	return domain.Device{
		OrganizationId:   r.OrganizationId,
		RoomId:           r.RoomId,
		InventoryNumber:  r.InventoryNumber,
		SerialNumber:     r.SerialNumber,
		Characteristics:  r.Characteristics,
		Category:         domain.Category(r.Category),
		Units:            r.Units,
		PowerConsumption: r.PowerConsumption,
	}, nil
}
