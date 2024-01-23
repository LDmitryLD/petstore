package controller

import (
	"encoding/json"
	"net/http"
	"projects/LDmitryLD/petstore/app/internal/infrastructure/responder"
	"projects/LDmitryLD/petstore/app/internal/models"
	"projects/LDmitryLD/petstore/app/internal/modules/store/service"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type Storeger interface {
	Create(w http.ResponseWriter, r *http.Request)
	GetByID(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Inventory(w http.ResponseWriter, r *http.Request)
}

type StoreController struct {
	service service.Storeger
	responder.Responder
}

func NewStore(service service.Storeger) Storeger {
	return &StoreController{
		service:   service,
		Responder: &responder.Respond{},
	}
}

func (s *StoreController) Create(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		s.ErrorBadRequest(w, err)
		return
	}

	resp := s.service.Create(order)

	s.OutputJSON(w, resp)
}

func (s *StoreController) GetByID(w http.ResponseWriter, r *http.Request) {
	var (
		err   error
		idRaw string
		id    int
	)

	idRaw = chi.URLParam(r, "orderId")

	id, err = strconv.Atoi(idRaw)
	if err != nil {
		s.ErrorBadRequest(w, err)
		return
	}

	order, err := s.service.GetByID(id)
	if err != nil {
		http.Error(w, models.NotFoundBody(), http.StatusNotFound)
		return
	}

	s.OutputJSON(w, order)
}

func (s *StoreController) Delete(w http.ResponseWriter, r *http.Request) {
	var (
		err   error
		idRaw string
		id    int
	)

	idRaw = chi.URLParam(r, "orderId")

	id, err = strconv.Atoi(idRaw)
	if err != nil {
		s.ErrorBadRequest(w, err)
		return
	}

	err = s.service.Delete(id)
	if err != nil {
		http.Error(w, models.NotFoundBody(), http.StatusNotFound)
	}

	s.OutputJSON(w, models.ApiResponse{Code: 200, Type: "unknown", Message: idRaw})
}

func (s *StoreController) Inventory(w http.ResponseWriter, r *http.Request) {
	inventory := s.service.Inventory()

	s.OutputJSON(w, inventory)
}
