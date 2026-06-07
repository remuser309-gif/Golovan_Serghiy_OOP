package controllers

import (
	"github.com/BohdanBoriak/boilerplate-go-back/internal/app"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/http/requests"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/http/resources"
	"log"
	"net/http"
	"strconv"
	"github.com/go-chi/chi/v5"
)

type RoomController struct {
	roomService         app.RoomService
	organizationService app.OrganizationService
}

func NewRoomController(rs app.RoomService, os app.OrganizationService) RoomController {
	return RoomController{
		roomService:         rs,
		organizationService: os,
	}
}

func (c RoomController) Save() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		orgId, err := strconv.ParseUint(chi.URLParam(r, "orgId"), 10, 64)
		if err != nil {
			log.Printf("RoomController: %s", err)
			BadRequest(w, err)
			return
		}

		user := r.Context().Value(UserKey).(domain.User)
		org, err := c.organizationService.Find(orgId)
		if err != nil {
			log.Printf("RoomController: %s", err)
			NotFound(w, err)
			return
		}
		if org.UserId != user.Id {
			Forbidden(w, err)
			return
		}

		room, err := requests.Bind(r, requests.RoomRequest{}, domain.Room{})
		if err != nil {
			log.Printf("RoomController: %s", err)
			BadRequest(w, err)
			return
		}

		room.OrganizationId = orgId
		room, err = c.roomService.Save(room)
		if err != nil {
			log.Printf("RoomController: %s", err)
			InternalServerError(w, err)
			return
		}

		var roomDto resources.RoomDto
		Created(w, roomDto.DomainToDto(room))
	}
}

func (c RoomController) FindByOrg() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		orgId, err := strconv.ParseUint(chi.URLParam(r, "orgId"), 10, 64)
		if err != nil {
			log.Printf("RoomController: %s", err)
			BadRequest(w, err)
			return
		}

		user := r.Context().Value(UserKey).(domain.User)
		org, err := c.organizationService.Find(orgId)
		if err != nil {
			log.Printf("RoomController: %s", err)
			NotFound(w, err)
			return
		}
		if org.UserId != user.Id {
			Forbidden(w, err)
			return
		}

		rooms, err := c.roomService.FindByOrg(orgId)
		if err != nil {
			log.Printf("RoomController: %s", err)
			InternalServerError(w, err)
			return
		}

		var roomDto resources.RoomDto
		Success(w, roomDto.DomainToDtoCollection(rooms))
	}
}

func (c RoomController) Find() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := parseIdFromUrl(r)
		if err != nil {
			log.Printf("RoomController: %s", err)
			BadRequest(w, err)
			return
		}

		room, err := c.roomService.Find(id)
		if err != nil {
			log.Printf("RoomController: %s", err)
			NotFound(w, err)
			return
		}

		user := r.Context().Value(UserKey).(domain.User)
		org, err := c.organizationService.Find(room.OrganizationId)
		if err != nil {
			log.Printf("RoomController: %s", err)
			Forbidden(w, err)
			return
		}
		if org.UserId != user.Id {
			Forbidden(w, err)
			return
		}

		var roomDto resources.RoomDto
		Success(w, roomDto.DomainToDto(room))
	}
}

func (c RoomController) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := parseIdFromUrl(r)
		if err != nil {
			log.Printf("RoomController: %s", err)
			BadRequest(w, err)
			return
		}

		existingRoom, err := c.roomService.Find(id)
		if err != nil {
			log.Printf("RoomController: %s", err)
			NotFound(w, err)
			return
		}

		user := r.Context().Value(UserKey).(domain.User)
		org, err := c.organizationService.Find(existingRoom.OrganizationId)
		if err != nil {
			log.Printf("RoomController: %s", err)
			Forbidden(w, err)
			return
		}
		if org.UserId != user.Id {
			Forbidden(w, err)
			return
		}

		room, err := requests.Bind(r, requests.RoomUpdateRequest{}, domain.Room{})
		if err != nil {
			log.Printf("RoomController: %s", err)
			BadRequest(w, err)
			return
		}

		room.Id = existingRoom.Id
		room.OrganizationId = existingRoom.OrganizationId
		room, err = c.roomService.Update(room)
		if err != nil {
			log.Printf("RoomController: %s", err)
			InternalServerError(w, err)
			return
		}

		var roomDto resources.RoomDto
		Success(w, roomDto.DomainToDto(room))
	}
}

func (c RoomController) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := parseIdFromUrl(r)
		if err != nil {
			log.Printf("RoomController: %s", err)
			BadRequest(w, err)
			return
		}

		existingRoom, err := c.roomService.Find(id)
		if err != nil {
			log.Printf("RoomController: %s", err)
			NotFound(w, err)
			return
		}

		user := r.Context().Value(UserKey).(domain.User)
		org, err := c.organizationService.Find(existingRoom.OrganizationId)
		if err != nil {
			log.Printf("RoomController: %s", err)
			Forbidden(w, err)
			return
		}
		if org.UserId != user.Id {
			Forbidden(w, err)
			return
		}

		err = c.roomService.Delete(id)
		if err != nil {
			log.Printf("RoomController: %s", err)
			InternalServerError(w, err)
			return
		}

		Ok(w)
	}
}
