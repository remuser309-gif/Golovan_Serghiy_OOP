package controllers

import (
	"github.com/remuser309-gif/Golovan_Serghiy_OOP/internal/app"
	"github.com/remuser309-gif/Golovan_Serghiy_OOP/internal/domain"
	"github.com/remuser309-gif/Golovan_Serghiy_OOP/internal/infra/http/requests"
	"github.com/remuser309-gif/Golovan_Serghiy_OOP/internal/infra/http/resources"
	"log"
	"net/http"
)

type OrganizationController struct {
	organizationService app.OrganizationService
}

func NewOrganizationController(os app.OrganizationService) OrganizationController {
	return OrganizationController{
		organizationService: os,
	}
}

func (c OrganizationController) Save() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(UserKey).(domain.User)
		org, err := requests.Bind(r, requests.OrganizationRequest{}, domain.Organization{})
		if err != nil {
			log.Printf("OrganizationController: %s", err)
			BadRequest(w, err)
			return
		}

		org.UserId = user.Id
		org, err = c.organizationService.Save(org)
		if err != nil {
			log.Printf("OrganizationController: %s", err)
			InternalServerError(w, err)
			return
		}

		var orgDto resources.OrganizationDto
		Created(w, orgDto.DomainToDto(org))
	}
}

func (c OrganizationController) Find() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := parseIdFromUrl(r)
		if err != nil {
			log.Printf("OrganizationController: %s", err)
			BadRequest(w, err)
			return
		}

		org, err := c.organizationService.Find(id)
		if err != nil {
			log.Printf("OrganizationController: %s", err)
			NotFound(w, err)
			return
		}

		user := r.Context().Value(UserKey).(domain.User)
		if org.UserId != user.Id {
			Forbidden(w, err)
			return
		}

		var orgDto resources.OrganizationDto
		Success(w, orgDto.DomainToDto(org))
	}
}

func (c OrganizationController) FindAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(UserKey).(domain.User)
		orgs, err := c.organizationService.FindByUser(user.Id)
		if err != nil {
			log.Printf("OrganizationController: %s", err)
			InternalServerError(w, err)
			return
		}

		var orgDto resources.OrganizationDto
		Success(w, orgDto.DomainToDtoCollection(orgs))
	}
}

func (c OrganizationController) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := parseIdFromUrl(r)
		if err != nil {
			log.Printf("OrganizationController: %s", err)
			BadRequest(w, err)
			return
		}

		org, err := requests.Bind(r, requests.OrganizationUpdateRequest{}, domain.Organization{})
		if err != nil {
			log.Printf("OrganizationController: %s", err)
			BadRequest(w, err)
			return
		}

		existingOrg, err := c.organizationService.Find(id)
		if err != nil {
			log.Printf("OrganizationController: %s", err)
			NotFound(w, err)
			return
		}

		user := r.Context().Value(UserKey).(domain.User)
		if existingOrg.UserId != user.Id {
			Forbidden(w, err)
			return
		}

		org.Id = existingOrg.Id
		org.UserId = user.Id
		org, err = c.organizationService.Update(org)
		if err != nil {
			log.Printf("OrganizationController: %s", err)
			InternalServerError(w, err)
			return
		}

		var orgDto resources.OrganizationDto
		Success(w, orgDto.DomainToDto(org))
	}
}

func (c OrganizationController) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := parseIdFromUrl(r)
		if err != nil {
			log.Printf("OrganizationController: %s", err)
			BadRequest(w, err)
			return
		}

		org, err := c.organizationService.Find(id)
		if err != nil {
			log.Printf("OrganizationController: %s", err)
			NotFound(w, err)
			return
		}

		user := r.Context().Value(UserKey).(domain.User)
		if org.UserId != user.Id {
			Forbidden(w, err)
			return
		}

		err = c.organizationService.Delete(id)
		if err != nil {
			log.Printf("OrganizationController: %s", err)
			InternalServerError(w, err)
			return
		}

		Ok(w)
	}
}
