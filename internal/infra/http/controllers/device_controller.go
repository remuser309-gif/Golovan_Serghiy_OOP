package controllers

import (
	"github.com/remuser309-gif/Golovan_Serghiy_OOP/internal/app"
	"github.com/remuser309-gif/Golovan_Serghiy_OOP/internal/domain"
	"github.com/remuser309-gif/Golovan_Serghiy_OOP/internal/infra/http/requests"
	"github.com/remuser309-gif/Golovan_Serghiy_OOP/internal/infra/http/resources"
	"log"
	"net/http"
	"strconv"
	"github.com/go-chi/chi/v5"
)

type DeviceController struct {
	deviceService       app.DeviceService
	organizationService app.OrganizationService
}

func NewDeviceController(ds app.DeviceService, os app.OrganizationService) DeviceController {
	return DeviceController{
		deviceService:       ds,
		organizationService: os,
	}
}

func (c DeviceController) Save() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		orgId, err := strconv.ParseUint(chi.URLParam(r, "orgId"), 10, 64)
		if err != nil {
			log.Printf("DeviceController: %s", err)
			BadRequest(w, err)
			return
		}

		user := r.Context().Value(UserKey).(domain.User)
		org, err := c.organizationService.Find(orgId)
		if err != nil {
			log.Printf("DeviceController: %s", err)
			NotFound(w, err)
			return
		}
		if org.UserId != user.Id {
			Forbidden(w, err)
			return
		}

		dev, err := requests.Bind(r, requests.DeviceRequest{}, domain.Device{})
		if err != nil {
			log.Printf("DeviceController: %s", err)
			BadRequest(w, err)
			return
		}

		dev.OrganizationId = orgId
		dev, err = c.deviceService.Save(dev)
		if err != nil {
			log.Printf("DeviceController: %s", err)
			InternalServerError(w, err)
			return
		}

		var devDto resources.DeviceDto
		Created(w, devDto.DomainToDto(dev))
	}
}

func (c DeviceController) FindByOrg() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		orgId, err := strconv.ParseUint(chi.URLParam(r, "orgId"), 10, 64)
		if err != nil {
			log.Printf("DeviceController: %s", err)
			BadRequest(w, err)
			return
		}

		user := r.Context().Value(UserKey).(domain.User)
		org, err := c.organizationService.Find(orgId)
		if err != nil {
			log.Printf("DeviceController: %s", err)
			NotFound(w, err)
			return
		}
		if org.UserId != user.Id {
			Forbidden(w, err)
			return
		}

		category := r.URL.Query().Get("category")
		var devices []domain.Device
		if category != "" {
			devices, err = c.deviceService.FindByCategory(domain.Category(category))
		} else {
			devices, err = c.deviceService.FindByOrg(orgId)
		}
		if err != nil {
			log.Printf("DeviceController: %s", err)
			InternalServerError(w, err)
			return
		}

		var devDto resources.DeviceDto
		Success(w, devDto.DomainToDtoCollection(devices))
	}
}

func (c DeviceController) FindByRoom() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		roomId, err := strconv.ParseUint(chi.URLParam(r, "roomId"), 10, 64)
		if err != nil {
			log.Printf("DeviceController: %s", err)
			BadRequest(w, err)
			return
		}

		user := r.Context().Value(UserKey).(domain.User)
		devices, err := c.deviceService.FindByRoom(roomId)
		if err != nil {
			log.Printf("DeviceController: %s", err)
			InternalServerError(w, err)
			return
		}

		for _, dev := range devices {
			org, err := c.organizationService.Find(dev.OrganizationId)
			if err != nil || org.UserId != user.Id {
				Forbidden(w, err)
				return
			}
		}

		var devDto resources.DeviceDto
		Success(w, devDto.DomainToDtoCollection(devices))
	}
}

func (c DeviceController) Find() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := parseIdFromUrl(r)
		if err != nil {
			log.Printf("DeviceController: %s", err)
			BadRequest(w, err)
			return
		}

		dev, err := c.deviceService.Find(id)
		if err != nil {
			log.Printf("DeviceController: %s", err)
			NotFound(w, err)
			return
		}

		user := r.Context().Value(UserKey).(domain.User)
		org, err := c.organizationService.Find(dev.OrganizationId)
		if err != nil || org.UserId != user.Id {
			Forbidden(w, err)
			return
		}

		var devDto resources.DeviceDto
		Success(w, devDto.DomainToDto(dev))
	}
}

func (c DeviceController) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := parseIdFromUrl(r)
		if err != nil {
			log.Printf("DeviceController: %s", err)
			BadRequest(w, err)
			return
		}

		existingDev, err := c.deviceService.Find(id)
		if err != nil {
			log.Printf("DeviceController: %s", err)
			NotFound(w, err)
			return
		}

		user := r.Context().Value(UserKey).(domain.User)
		org, err := c.organizationService.Find(existingDev.OrganizationId)
		if err != nil || org.UserId != user.Id {
			Forbidden(w, err)
			return
		}

		dev, err := requests.Bind(r, requests.DeviceUpdateRequest{}, domain.Device{})
		if err != nil {
			log.Printf("DeviceController: %s", err)
			BadRequest(w, err)
			return
		}

		dev.Id = existingDev.Id
		dev.OrganizationId = existingDev.OrganizationId
		dev.GUID = existingDev.GUID
		dev, err = c.deviceService.Update(dev)
		if err != nil {
			log.Printf("DeviceController: %s", err)
			InternalServerError(w, err)
			return
		}

		var devDto resources.DeviceDto
		Success(w, devDto.DomainToDto(dev))
	}
}

func (c DeviceController) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := parseIdFromUrl(r)
		if err != nil {
			log.Printf("DeviceController: %s", err)
			BadRequest(w, err)
			return
		}

		existingDev, err := c.deviceService.Find(id)
		if err != nil {
			log.Printf("DeviceController: %s", err)
			NotFound(w, err)
			return
		}

		user := r.Context().Value(UserKey).(domain.User)
		org, err := c.organizationService.Find(existingDev.OrganizationId)
		if err != nil || org.UserId != user.Id {
			Forbidden(w, err)
			return
		}

		err = c.deviceService.Delete(id)
		if err != nil {
			log.Printf("DeviceController: %s", err)
			InternalServerError(w, err)
			return
		}

		Ok(w)
	}
}
