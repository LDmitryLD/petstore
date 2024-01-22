package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"projects/LDmitryLD/petstore/app/internal/infrastructure/responder"
	"projects/LDmitryLD/petstore/app/internal/models"
	"projects/LDmitryLD/petstore/app/internal/modules/pet/service"
	"strconv"

	"github.com/go-chi/chi/v5"
)

const (
	LogErrDecodeReq = "ошибка при декодировании запроса"
)

type Peterer interface {
	UploadImage(w http.ResponseWriter, r *http.Request)
	AddPet(w http.ResponseWriter, r *http.Request)
	UpdateExistinPet(w http.ResponseWriter, r *http.Request)
	FindByStatus(w http.ResponseWriter, r *http.Request)
	GetByID(w http.ResponseWriter, r *http.Request)
	UpdateByID(w http.ResponseWriter, r *http.Request)
	DeletePet(w http.ResponseWriter, r *http.Request)
}

type PetController struct {
	service service.Peterer
	responder.Responder
}

func NewPet(service service.Peterer) Peterer {
	return &PetController{
		service:   service,
		Responder: &responder.Respond{},
	}
}

func (p *PetController) UploadImage(w http.ResponseWriter, r *http.Request) {
	var (
		err      error
		petIDRaw string
		petID    int64
	)

	petIDRaw = chi.URLParam(r, "petId")
	petID, err = strconv.ParseInt(petIDRaw, 0, 64)
	if err != nil {
		p.ErrorBadRequest(w, err)
		return
	}

	additionalMetadata := r.FormValue("additionalMetadata")

	file, handler, err := r.FormFile("file")
	if err != nil {
		p.ErrorBadRequest(w, err)
		return
	}
	defer file.Close()

	out := p.service.UploadImage(service.UploadIn{PetID: int(petID), File: file, Handler: handler})

	if out.Error != nil {
		resp := models.ApiResponse{
			Code:    http.StatusInternalServerError,
			Message: out.Error.Error(),
		}
		p.OutputJSON(w, resp)
		return
	}

	resp := models.ApiResponse{
		Code:    http.StatusOK,
		Type:    out.FileType,
		Message: fmt.Sprintf("additionalMetadata: %s\nFile uploaded to ./%s", additionalMetadata, out.FileName),
	}

	p.OutputJSON(w, resp)
}

func (p *PetController) AddPet(w http.ResponseWriter, r *http.Request) {
	var pet models.Pet
	if err := json.NewDecoder(r.Body).Decode(&pet); err != nil {
		log.Println(LogErrDecodeReq)
		p.ErrorBadRequest(w, err)
		return
	}

	out := p.service.Create(service.CreateIn{Pet: pet})

	p.OutputJSON(w, out.Pet)
}

func (p *PetController) UpdateExistinPet(w http.ResponseWriter, r *http.Request) {
	var pet models.Pet
	if err := json.NewDecoder(r.Body).Decode(&pet); err != nil {
		p.ErrorBadRequest(w, err)
		return
	}

	out := p.service.Update(service.UpdateIn{Pet: pet})
	if out.Error != nil {
		http.Error(w, out.Error.Error(), http.StatusNotFound)
	}
}

func (p *PetController) FindByStatus(w http.ResponseWriter, r *http.Request) {
	statusValue := r.URL.Query()["status"]
	if statusValue == nil {
		p.ErrorBadRequest(w, fmt.Errorf("invalid status values"))
		return
	}
	pets := p.service.FindByStatus(statusValue)

	p.OutputJSON(w, pets)
}

func (p *PetController) GetByID(w http.ResponseWriter, r *http.Request) {
	var (
		err      error
		petIDRaw string
		petID    int64
	)

	petIDRaw = chi.URLParam(r, "petId")

	petID, err = strconv.ParseInt(petIDRaw, 0, 64)
	if err != nil {
		p.ErrorBadRequest(w, err)
		return
	}

	out := p.service.GetByID(service.GetByIDIn{PetID: int(petID)})

	if out.Error != nil {
		p.ErrorNotFound(w, out.Error)
		return
	}

	p.OutputJSON(w, out.Pet)
}

func (p *PetController) UpdateByID(w http.ResponseWriter, r *http.Request) {
	var (
		err      error
		petIDRaw string
		petID    int
		name     string
		status   string
	)

	petIDRaw = chi.URLParam(r, "petId")
	petID, err = strconv.Atoi(petIDRaw)
	if err != nil {
		p.ErrorBadRequest(w, err)
		return
	}
	name = r.FormValue("name")
	status = r.FormValue("status")

	if err := p.service.UpdateByID(service.UpdateByIDIn{PetID: petID, Name: name, Status: status}); err != nil {
		resp := models.ApiResponse{
			Code:    http.StatusNotFound,
			Type:    "error",
			Message: err.Error(),
		}
		p.OutputJSON(w, resp)
		return
	}

	resp := models.ApiResponse{
		Code:    http.StatusOK,
		Type:    "unknown",
		Message: petIDRaw,
	}

	p.OutputJSON(w, resp)
}

func (p *PetController) DeletePet(w http.ResponseWriter, r *http.Request) {
	var (
		err      error
		petIDRaw string
		petID    int
	)

	petIDRaw = chi.URLParam(r, "petId")
	petID, err = strconv.Atoi(petIDRaw)
	if err != nil {
		p.ErrorBadRequest(w, err)
		return
	}

	if err := p.service.Delete(petID); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	resp := models.ApiResponse{
		Code:    http.StatusOK,
		Type:    "unknown",
		Message: petIDRaw,
	}

	p.OutputJSON(w, resp)
}
