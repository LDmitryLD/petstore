package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"projects/LDmitryLD/petstore/app/internal/infrastructure/responder"
	"projects/LDmitryLD/petstore/app/internal/models"
	"projects/LDmitryLD/petstore/app/internal/modules/user/service/mocks"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

const (
	testReqErr        = "ошибка при выполнении тестового запроса:"
	testRespDecodeErr = "ошибка при декодировании тестового ответа:"
)

var (
	id       = 1
	username = "username"
	testUser = models.User{
		ID:        1,
		FirstName: "Name",
		Username:  "username",
	}
)

func TestUserController_Create_BadRequest(t *testing.T) {
	user := UserController{
		Responder: &responder.Respond{},
	}

	req := map[string]interface{}{"ID": "1"}
	reqJSON, _ := json.Marshal(req)

	s := httptest.NewServer(http.HandlerFunc(user.Create))
	defer s.Close()

	resp, err := http.Post(s.URL, "application/json", bytes.NewBuffer(reqJSON))
	if err != nil {
		t.Fatal(testReqErr, err.Error())
	}
	defer resp.Body.Close()

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestUserController_Create(t *testing.T) {
	expect := models.ApiResponse{
		Code:    200,
		Type:    "unknown",
		Message: fmt.Sprint(id),
	}

	serviceMock := mocks.NewUserer(t)
	serviceMock.On("Create", testUser).Return(id)

	user := NewUser(serviceMock)

	reqJSON, _ := json.Marshal(testUser)

	s := httptest.NewServer(http.HandlerFunc(user.Create))
	defer s.Close()

	resp, err := http.Post(s.URL, "application/json", bytes.NewBuffer(reqJSON))
	if err != nil {
		t.Fatal(testReqErr, err.Error())
	}
	defer resp.Body.Close()

	var createResp models.ApiResponse
	if err = json.NewDecoder(resp.Body).Decode(&createResp); err != nil {
		t.Fatal(testRespDecodeErr, err.Error())
	}

	assert.Equal(t, expect, createResp)
}

func TestUserController_CreateWithArray_BadRequest(t *testing.T) {
	user := UserController{
		Responder: &responder.Respond{},
	}

	req := map[string]interface{}{"id": "1"}
	reqJSON, _ := json.Marshal(req)

	s := httptest.NewServer(http.HandlerFunc(user.CreateWithArray))
	defer s.Close()

	resp, err := http.Post(s.URL, "application/json", bytes.NewBuffer(reqJSON))
	if err != nil {
		t.Fatal(testReqErr, err.Error())
	}
	defer resp.Body.Close()

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestUserController_CreateWithArray(t *testing.T) {
	users := []models.User{testUser, testUser}
	reqJSON, _ := json.Marshal(users)
	expect := models.ApiResponse{
		Code:    200,
		Type:    "unknown",
		Message: "ok",
	}

	serviceMock := mocks.NewUserer(t)
	serviceMock.On("Create", testUser).Return(id).Times(2)

	user := NewUser(serviceMock)

	s := httptest.NewServer(http.HandlerFunc(user.CreateWithArray))
	defer s.Close()

	resp, err := http.Post(s.URL, "application/json", bytes.NewBuffer(reqJSON))
	if err != nil {
		t.Fatal(testReqErr, err.Error())
	}
	defer resp.Body.Close()

	var withArrResp models.ApiResponse
	if err = json.NewDecoder(resp.Body).Decode(&withArrResp); err != nil {
		t.Fatal(testRespDecodeErr, err.Error())
	}

	assert.Equal(t, expect, withArrResp)
}

func TestUserController_GetByUsername_NotFound(t *testing.T) {
	serviceMock := mocks.NewUserer(t)
	serviceMock.On("GetByUsername", username).Return(models.User{}, fmt.Errorf("not found"))

	user := NewUser(serviceMock)

	rr := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/users/username", nil)

	r := chi.NewRouter()

	r.MethodFunc("GET", "/users/{username}", user.GetByUsername)

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code)
}

func TestUserController_GetByUsername(t *testing.T) {
	serviceMock := mocks.NewUserer(t)
	serviceMock.On("GetByUsername", username).Return(testUser, nil)

	user := NewUser(serviceMock)

	rr := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/users/username", nil)

	r := chi.NewRouter()

	r.MethodFunc("GET", "/users/{username}", user.GetByUsername)

	r.ServeHTTP(rr, req)

	var resp models.User
	if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
		t.Fatal(testRespDecodeErr, err.Error())
	}

	assert.Equal(t, testUser, resp)
}

func TestUserController_Update_BadRequest(t *testing.T) {
	user := UserController{
		Responder: &responder.Respond{},
	}
	req := map[string]interface{}{"id": "1"}
	reqJSON, _ := json.Marshal(req)

	s := httptest.NewServer(http.HandlerFunc(user.Update))
	defer s.Close()

	resp, err := http.Post(s.URL, "application/json", bytes.NewBuffer(reqJSON))
	if err != nil {
		t.Fatal(testReqErr, err.Error())
	}

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestUserController_Update_NotFound(t *testing.T) {
	reqJSON, _ := json.Marshal(testUser)

	serviceMock := mocks.NewUserer(t)
	serviceMock.On("Update", username, testUser).Return(fmt.Errorf("not found"))

	user := NewUser(serviceMock)

	rr := httptest.NewRecorder()
	req := httptest.NewRequest("PUT", "/users/username", bytes.NewBuffer(reqJSON))

	r := chi.NewRouter()
	r.MethodFunc("PUT", "/users/{username}", user.Update)

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code)
}

func TestUserController_Update(t *testing.T) {
	expect := models.ApiResponse{
		Code:    200,
		Type:    "unknown",
		Message: "ok",
	}
	reqJSON, _ := json.Marshal(testUser)

	serviceMock := mocks.NewUserer(t)
	serviceMock.On("Update", username, testUser).Return(nil)

	user := NewUser(serviceMock)

	rr := httptest.NewRecorder()
	req := httptest.NewRequest("PUT", "/users/username", bytes.NewBuffer(reqJSON))

	r := chi.NewRouter()
	r.MethodFunc("PUT", "/users/{username}", user.Update)

	r.ServeHTTP(rr, req)

	var resp models.ApiResponse
	if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
		t.Fatal(testRespDecodeErr, err.Error())
	}

	assert.Equal(t, expect, resp)
}

func TestUserController_Delete_NotFound(t *testing.T) {
	serviceMock := mocks.NewUserer(t)
	serviceMock.On("Delete", username).Return(fmt.Errorf("not found"))

	user := NewUser(serviceMock)

	rr := httptest.NewRecorder()
	req := httptest.NewRequest("DELETE", "/users/username", nil)

	r := chi.NewRouter()
	r.MethodFunc("DELETE", "/users/{username}", user.Delete)

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code)
}

func TestUserController_Delete(t *testing.T) {
	expect := models.ApiResponse{
		Code:    200,
		Type:    "unknown",
		Message: "user deleted",
	}

	serviceMock := mocks.NewUserer(t)
	serviceMock.On("Delete", username).Return(nil)

	user := NewUser(serviceMock)

	rr := httptest.NewRecorder()
	req := httptest.NewRequest("DELETE", "/users/username", nil)

	r := chi.NewRouter()
	r.MethodFunc("DELETE", "/users/{username}", user.Delete)

	r.ServeHTTP(rr, req)

	var resp models.ApiResponse
	if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
		t.Fatal(testRespDecodeErr, err.Error())
	}

	assert.Equal(t, expect, resp)
}
