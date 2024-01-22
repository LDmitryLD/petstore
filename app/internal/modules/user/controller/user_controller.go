package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"projects/LDmitryLD/petstore/app/internal/infrastructure/responder"
	"projects/LDmitryLD/petstore/app/internal/models"
	"projects/LDmitryLD/petstore/app/internal/modules/user/service"

	"github.com/go-chi/chi/v5"
)

const (
	ErrInvalidUsername = "invalid username supplied"
)

type Userer interface {
	CreateWithArray(w http.ResponseWriter, r *http.Request)
	GetByUsername(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	// Login(w http.ResponseWriter, r *http.Request)
	// Logout(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
}

type UserController struct {
	service service.Userer
	responder.Responder
}

func NewUser(service service.Userer) Userer {
	return &UserController{
		service:   service,
		Responder: &responder.Respond{},
	}
}

// func (u *UserController) Login(w http.ResponseWriter, r *http.Request) {
// 	username := r.URL.Query().Get("username")
// 	password := r.URL.Query().Get("password")

// 	if username == "" || password == "" {
// 		u.ErrorBadRequest(w, fmt.Errorf("invalid username/password supplied"))
// 	}

// 	u.OutputJSON(w, LoginResponse{Code: http.StatusOK, Type: "unknown", Message: })
// }

func (u *UserController) Create(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		u.ErrorBadRequest(w, err)
		return
	}

	id := u.service.Create(user)

	u.OutputJSON(w, models.ApiResponse{Code: 200, Type: "unknown", Message: fmt.Sprint(id)})
}

func (u *UserController) CreateWithArray(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	if err := json.NewDecoder(r.Body).Decode(&users); err != nil {
		u.ErrorBadRequest(w, err)
		return
	}

	for i := range users {
		u.service.Create(users[i])
	}

	u.OutputJSON(w, models.ApiResponse{Code: 200, Type: "unknown", Message: "ok"})
}

func (u *UserController) GetByUsername(w http.ResponseWriter, r *http.Request) {
	userName := chi.URLParam(r, "username")
	if userName == "" {
		u.ErrorBadRequest(
			w, fmt.Errorf(ErrInvalidUsername))
		return
	}

	user, err := u.service.GetByUsername(userName)
	if err != nil {
		msg, _ := json.Marshal(models.ApiResponse{Code: http.StatusNotFound, Type: "error", Message: err.Error()})
		http.Error(w, string(msg), http.StatusNotFound)
		return
	}

	u.OutputJSON(w, user)
}

func (u *UserController) Update(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")
	if username == "" {
		u.ErrorBadRequest(w, fmt.Errorf(ErrInvalidUsername))
		return
	}

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		u.ErrorBadRequest(w, err)
		return
	}

	if err := u.service.Update(username, user); err != nil {
		msg, _ := json.Marshal(models.ApiResponse{Code: http.StatusNotFound, Type: "error", Message: err.Error()})
		http.Error(w, string(msg), http.StatusNotFound)
		return
	}

	u.OutputJSON(w, models.ApiResponse{Code: 200, Type: "unknown", Message: ""}) // что передавать в message надо подумать
}

func (u *UserController) Delete(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")
	if username == "" {
		u.ErrorBadRequest(w, fmt.Errorf(ErrInvalidUsername))
		return
	}

	if err := u.service.Delete(username); err != nil {
		msg, _ := json.Marshal(models.ApiResponse{Code: http.StatusNotFound, Type: "error", Message: err.Error()})
		http.Error(w, string(msg), http.StatusNotFound)
		return
	}

	u.OutputJSON(w, models.ApiResponse{Code: 200, Type: "unknown", Message: "user deleted"})
}
