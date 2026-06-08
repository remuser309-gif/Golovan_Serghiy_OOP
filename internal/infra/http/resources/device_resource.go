package resources

import "github.com/remuser309-gif/Golovan_Serghiy_OOP/internal/domain"

type DeviceDto struct {
	Id               uint64          `json:"id"`
	OrganizationId   uint64          `json:"organizationId"`
	RoomId           uint64          `json:"roomId"`
	GUID             string          `json:"guid"`
	InventoryNumber  string          `json:"inventoryNumber"`
	SerialNumber     string          `json:"serialNumber"`
	Characteristics  string          `json:"characteristics"`
	Category         domain.Category `json:"category"`
	Units            string          `json:"units"`
	PowerConsumption float64         `json:"powerConsumption"`
}

type DevicesDto struct {
	Items []DeviceDto `json:"items"`
	Total uint64      `json:"total"`
	Pages uint        `json:"pages"`
}

func (d DeviceDto) DomainToDto(dev domain.Device) DeviceDto {
	return DeviceDto{
		Id:               dev.Id,
		OrganizationId:   dev.OrganizationId,
		RoomId:           dev.RoomId,
		GUID:             dev.GUID.String(),
		InventoryNumber:  dev.InventoryNumber,
		SerialNumber:     dev.SerialNumber,
		Characteristics:  dev.Characteristics,
		Category:         dev.Category,
		Units:            dev.Units,
		PowerConsumption: dev.PowerConsumption,
	}
}

func (d DeviceDto) DomainToDtoCollection(devices []domain.Device) []DeviceDto {
	result := make([]DeviceDto, len(devices))
	for i, dev := range devices {
		result[i] = d.DomainToDto(dev)
	}
	return result
}
