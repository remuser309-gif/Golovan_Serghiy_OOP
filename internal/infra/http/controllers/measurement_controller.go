package controllers

import (
	"github.com/BohdanBoriak/boilerplate-go-back/internal/app"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/http/requests"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/http/resources"
	"log"
	"net/http"
	"strconv"
	"time"
	"github.com/go-chi/chi/v5"
)

type MeasurementController struct {
	measurementService  app.MeasurementService
	deviceService       app.DeviceService
	organizationService app.OrganizationService
}

func NewMeasurementController(ms app.MeasurementService, ds app.DeviceService, os app.OrganizationService) MeasurementController {
	return MeasurementController{
		measurementService:  ms,
		deviceService:       ds,
		organizationService: os,
	}
}

func (c MeasurementController) Save() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		devId, err := strconv.ParseUint(chi.URLParam(r, "devId"), 10, 64)
		if err != nil {
			log.Printf("MeasurementController: %s", err)
			BadRequest(w, err)
			return
		}

		user := r.Context().Value(UserKey).(domain.User)
		dev, err := c.deviceService.Find(devId)
		if err != nil {
			log.Printf("MeasurementController: %s", err)
			NotFound(w, err)
			return
		}
		org, err := c.organizationService.Find(dev.OrganizationId)
		if err != nil || org.UserId != user.Id {
			Forbidden(w, err)
			return
		}

		m, err := requests.Bind(r, requests.MeasurementRequest{}, domain.Measurement{})
		if err != nil {
			log.Printf("MeasurementController: %s", err)
			BadRequest(w, err)
			return
		}

		m.DeviceId = devId
		m.RoomId = dev.RoomId
		m, err = c.measurementService.Save(m)
		if err != nil {
			log.Printf("MeasurementController: %s", err)
			InternalServerError(w, err)
			return
		}

		var mDto resources.MeasurementDto
		Created(w, mDto.DomainToDto(m))
	}
}

func (c MeasurementController) FindByDevice() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		devId, err := strconv.ParseUint(chi.URLParam(r, "devId"), 10, 64)
		if err != nil {
			log.Printf("MeasurementController: %s", err)
			BadRequest(w, err)
			return
		}

		user := r.Context().Value(UserKey).(domain.User)
		dev, err := c.deviceService.Find(devId)
		if err != nil {
			log.Printf("MeasurementController: %s", err)
			NotFound(w, err)
			return
		}
		org, err := c.organizationService.Find(dev.OrganizationId)
		if err != nil || org.UserId != user.Id {
			Forbidden(w, err)
			return
		}

		fromStr := r.URL.Query().Get("from")
		toStr := r.URL.Query().Get("to")

		var measurements []domain.Measurement
		if fromStr != "" && toStr != "" {
			from, err := time.Parse(time.RFC3339, fromStr)
			if err != nil {
				log.Printf("MeasurementController: %s", err)
				BadRequest(w, err)
				return
			}
			to, err := time.Parse(time.RFC3339, toStr)
			if err != nil {
				log.Printf("MeasurementController: %s", err)
				BadRequest(w, err)
				return
			}
			measurements, err = c.measurementService.FindByDeviceAndDateRange(devId, from, to)
			if err != nil {
				log.Printf("MeasurementController: %s", err)
				InternalServerError(w, err)
				return
			}
		} else {
			measurements, err = c.measurementService.FindByDevice(devId)
			if err != nil {
				log.Printf("MeasurementController: %s", err)
				InternalServerError(w, err)
				return
			}
		}

		var mDto resources.MeasurementDto
		Success(w, mDto.DomainToDtoCollection(measurements))
	}
}

func (c MeasurementController) Find() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := parseIdFromUrl(r)
		if err != nil {
			log.Printf("MeasurementController: %s", err)
			BadRequest(w, err)
			return
		}

		m, err := c.measurementService.Find(id)
		if err != nil {
			log.Printf("MeasurementController: %s", err)
			NotFound(w, err)
			return
		}

		user := r.Context().Value(UserKey).(domain.User)
		dev, err := c.deviceService.Find(m.DeviceId)
		if err != nil {
			log.Printf("MeasurementController: %s", err)
			NotFound(w, err)
			return
		}
		org, err := c.organizationService.Find(dev.OrganizationId)
		if err != nil || org.UserId != user.Id {
			Forbidden(w, err)
			return
		}

		var mDto resources.MeasurementDto
		Success(w, mDto.DomainToDto(m))
	}
}

func (c MeasurementController) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := parseIdFromUrl(r)
		if err != nil {
			log.Printf("MeasurementController: %s", err)
			BadRequest(w, err)
			return
		}

		existingM, err := c.measurementService.Find(id)
		if err != nil {
			log.Printf("MeasurementController: %s", err)
			NotFound(w, err)
			return
		}

		user := r.Context().Value(UserKey).(domain.User)
		dev, err := c.deviceService.Find(existingM.DeviceId)
		if err != nil {
			log.Printf("MeasurementController: %s", err)
			NotFound(w, err)
			return
		}
		org, err := c.organizationService.Find(dev.OrganizationId)
		if err != nil || org.UserId != user.Id {
			Forbidden(w, err)
			return
		}

		m, err := requests.Bind(r, requests.MeasurementUpdateRequest{}, domain.Measurement{})
		if err != nil {
			log.Printf("MeasurementController: %s", err)
			BadRequest(w, err)
			return
		}

		m.Id = existingM.Id
		m.DeviceId = existingM.DeviceId
		m.RoomId = existingM.RoomId
		m, err = c.measurementService.Update(m)
		if err != nil {
			log.Printf("MeasurementController: %s", err)
			InternalServerError(w, err)
			return
		}

		var mDto resources.MeasurementDto
		Success(w, mDto.DomainToDto(m))
	}
}

func (c MeasurementController) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := parseIdFromUrl(r)
		if err != nil {
			log.Printf("MeasurementController: %s", err)
			BadRequest(w, err)
			return
		}

		existingM, err := c.measurementService.Find(id)
		if err != nil {
			log.Printf("MeasurementController: %s", err)
			NotFound(w, err)
			return
		}

		user := r.Context().Value(UserKey).(domain.User)
		dev, err := c.deviceService.Find(existingM.DeviceId)
		if err != nil {
			log.Printf("MeasurementController: %s", err)
			NotFound(w, err)
			return
		}
		org, err := c.organizationService.Find(dev.OrganizationId)
		if err != nil || org.UserId != user.Id {
			Forbidden(w, err)
			return
		}

		err = c.measurementService.Delete(id)
		if err != nil {
			log.Printf("MeasurementController: %s", err)
			InternalServerError(w, err)
			return
		}

		Ok(w)
	}
}
