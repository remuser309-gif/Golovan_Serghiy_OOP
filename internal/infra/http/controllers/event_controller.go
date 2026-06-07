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

type EventController struct {
	eventService        app.EventService
	deviceService       app.DeviceService
	organizationService app.OrganizationService
}

func NewEventController(es app.EventService, ds app.DeviceService, os app.OrganizationService) EventController {
	return EventController{
		eventService:        es,
		deviceService:       ds,
		organizationService: os,
	}
}

func (c EventController) Save() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		devId, err := strconv.ParseUint(chi.URLParam(r, "devId"), 10, 64)
		if err != nil {
			log.Printf("EventController: %s", err)
			BadRequest(w, err)
			return
		}

		user := r.Context().Value(UserKey).(domain.User)
		dev, err := c.deviceService.Find(devId)
		if err != nil {
			log.Printf("EventController: %s", err)
			NotFound(w, err)
			return
		}
		org, err := c.organizationService.Find(dev.OrganizationId)
		if err != nil || org.UserId != user.Id {
			Forbidden(w, err)
			return
		}

		e, err := requests.Bind(r, requests.EventRequest{}, domain.Event{})
		if err != nil {
			log.Printf("EventController: %s", err)
			BadRequest(w, err)
			return
		}

		e.DeviceId = devId
		e.RoomId = dev.RoomId
		e, err = c.eventService.Save(e)
		if err != nil {
			log.Printf("EventController: %s", err)
			InternalServerError(w, err)
			return
		}

		var eDto resources.EventDto
		Created(w, eDto.DomainToDto(e))
	}
}

func (c EventController) FindByDevice() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		devId, err := strconv.ParseUint(chi.URLParam(r, "devId"), 10, 64)
		if err != nil {
			log.Printf("EventController: %s", err)
			BadRequest(w, err)
			return
		}

		user := r.Context().Value(UserKey).(domain.User)
		dev, err := c.deviceService.Find(devId)
		if err != nil {
			log.Printf("EventController: %s", err)
			NotFound(w, err)
			return
		}
		org, err := c.organizationService.Find(dev.OrganizationId)
		if err != nil || org.UserId != user.Id {
			Forbidden(w, err)
			return
		}

		action := r.URL.Query().Get("action")
		fromStr := r.URL.Query().Get("from")
		toStr := r.URL.Query().Get("to")

		var events []domain.Event
		if fromStr != "" && toStr != "" {
			from, err := time.Parse(time.RFC3339, fromStr)
			if err != nil {
				log.Printf("EventController: %s", err)
				BadRequest(w, err)
				return
			}
			to, err := time.Parse(time.RFC3339, toStr)
			if err != nil {
				log.Printf("EventController: %s", err)
				BadRequest(w, err)
				return
			}
			events, err = c.eventService.FindByDeviceAndDateRange(devId, from, to)
			if err != nil {
				log.Printf("EventController: %s", err)
				InternalServerError(w, err)
				return
			}
		} else if action != "" {
			events, err = c.eventService.FindByAction(action)
			if err != nil {
				log.Printf("EventController: %s", err)
				InternalServerError(w, err)
				return
			}
		} else {
			events, err = c.eventService.FindByDevice(devId)
			if err != nil {
				log.Printf("EventController: %s", err)
				InternalServerError(w, err)
				return
			}
		}

		var eDto resources.EventDto
		Success(w, eDto.DomainToDtoCollection(events))
	}
}

func (c EventController) Find() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := parseIdFromUrl(r)
		if err != nil {
			log.Printf("EventController: %s", err)
			BadRequest(w, err)
			return
		}

		e, err := c.eventService.Find(id)
		if err != nil {
			log.Printf("EventController: %s", err)
			NotFound(w, err)
			return
		}

		user := r.Context().Value(UserKey).(domain.User)
		dev, err := c.deviceService.Find(e.DeviceId)
		if err != nil {
			log.Printf("EventController: %s", err)
			NotFound(w, err)
			return
		}
		org, err := c.organizationService.Find(dev.OrganizationId)
		if err != nil || org.UserId != user.Id {
			Forbidden(w, err)
			return
		}

		var eDto resources.EventDto
		Success(w, eDto.DomainToDto(e))
	}
}

func (c EventController) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := parseIdFromUrl(r)
		if err != nil {
			log.Printf("EventController: %s", err)
			BadRequest(w, err)
			return
		}

		existingE, err := c.eventService.Find(id)
		if err != nil {
			log.Printf("EventController: %s", err)
			NotFound(w, err)
			return
		}

		user := r.Context().Value(UserKey).(domain.User)
		dev, err := c.deviceService.Find(existingE.DeviceId)
		if err != nil {
			log.Printf("EventController: %s", err)
			NotFound(w, err)
			return
		}
		org, err := c.organizationService.Find(dev.OrganizationId)
		if err != nil || org.UserId != user.Id {
			Forbidden(w, err)
			return
		}

		e, err := requests.Bind(r, requests.EventUpdateRequest{}, domain.Event{})
		if err != nil {
			log.Printf("EventController: %s", err)
			BadRequest(w, err)
			return
		}

		e.Id = existingE.Id
		e.DeviceId = existingE.DeviceId
		e.RoomId = existingE.RoomId
		e, err = c.eventService.Update(e)
		if err != nil {
			log.Printf("EventController: %s", err)
			InternalServerError(w, err)
			return
		}

		var eDto resources.EventDto
		Success(w, eDto.DomainToDto(e))
	}
}

func (c EventController) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := parseIdFromUrl(r)
		if err != nil {
			log.Printf("EventController: %s", err)
			BadRequest(w, err)
			return
		}

		existingE, err := c.eventService.Find(id)
		if err != nil {
			log.Printf("EventController: %s", err)
			NotFound(w, err)
			return
		}

		user := r.Context().Value(UserKey).(domain.User)
		dev, err := c.deviceService.Find(existingE.DeviceId)
		if err != nil {
			log.Printf("EventController: %s", err)
			NotFound(w, err)
			return
		}
		org, err := c.organizationService.Find(dev.OrganizationId)
		if err != nil || org.UserId != user.Id {
			Forbidden(w, err)
			return
		}

		err = c.eventService.Delete(id)
		if err != nil {
			log.Printf("EventController: %s", err)
			InternalServerError(w, err)
			return
		}

		Ok(w)
	}
}
