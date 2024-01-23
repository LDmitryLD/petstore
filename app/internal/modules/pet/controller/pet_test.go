package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"projects/LDmitryLD/petstore/app/internal/infrastructure/responder"
	"projects/LDmitryLD/petstore/app/internal/models"
	"projects/LDmitryLD/petstore/app/internal/modules/pet/service"
	"projects/LDmitryLD/petstore/app/internal/modules/pet/service/mocks"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

const (
	testReqErr        = "ошибка при выполнении тестового запроса:"
	testRespDecodeErr = "ошибка при декодировании тестового ответа:"
)

var (
	id      = 1
	name    = "name"
	status  = "available"
	testPet = models.Pet{
		ID:     1,
		Name:   "name",
		Status: "available",
	}
)

func TestPetController_UploadImage_BadRequest(t *testing.T) {
	pet := PetController{
		Responder: &responder.Respond{},
	}

	rr := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/pet/id/uploadImage", nil)

	r := chi.NewRouter()
	r.MethodFunc("POST", "/pet/{petId}/uploadImage", pet.UploadImage)

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestPetController_UploadImage_BadRequest2(t *testing.T) {
	pet := PetController{
		Responder: &responder.Respond{},
	}

	rr := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/pet/1/uploadImage", nil)

	r := chi.NewRouter()
	r.MethodFunc("POST", "/pet/{petId}/uploadImage", pet.UploadImage)

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestPetController_AddPet_BadRequest(t *testing.T) {
	pet := PetController{
		Responder: &responder.Respond{},
	}
	req := map[string]interface{}{"ID": "1"}
	reqJSON, _ := json.Marshal(req)

	s := httptest.NewServer(http.HandlerFunc(pet.AddPet))
	defer s.Close()

	resp, err := http.Post(s.URL, "application/json", bytes.NewBuffer(reqJSON))
	if err != nil {
		t.Fatal(testReqErr, err.Error())
	}
	defer resp.Body.Close()

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestPetController_AddPet(t *testing.T) {
	in := service.CreateIn{
		Pet: testPet,
	}
	out := service.CreateOut{
		Pet: testPet,
	}
	reqJSON, _ := json.Marshal(testPet)

	serviceMock := mocks.NewPeterer(t)
	serviceMock.On("Create", in).Return(out)

	pet := NewPet(serviceMock)

	s := httptest.NewServer(http.HandlerFunc(pet.AddPet))
	defer s.Close()

	resp, err := http.Post(s.URL, "application/json", bytes.NewBuffer(reqJSON))
	if err != nil {
		t.Fatal(testReqErr, err.Error())
	}
	defer resp.Body.Close()

	var addResp models.Pet
	if err := json.NewDecoder(resp.Body).Decode(&addResp); err != nil {
		t.Fatal(testRespDecodeErr, err.Error())
	}

	assert.Equal(t, testPet, addResp)
}

func TestPetController_UpdateExistinPet_BadRequest(t *testing.T) {
	pet := PetController{
		Responder: &responder.Respond{},
	}
	req := map[string]interface{}{"ID": "1"}
	reqJSON, _ := json.Marshal(req)

	s := httptest.NewServer(http.HandlerFunc(pet.UpdateExistinPet))
	defer s.Close()

	resp, err := http.Post(s.URL, "application/json", bytes.NewBuffer(reqJSON))
	if err != nil {
		t.Fatal(testReqErr, err.Error())
	}
	defer resp.Body.Close()

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestPetController_UpdateExistinPet_NotFound(t *testing.T) {
	in := service.UpdateIn{
		Pet: testPet,
	}
	out := service.UpdateOut{
		Error: fmt.Errorf("not found"),
	}
	reqJSON, _ := json.Marshal(testPet)

	serviceMock := mocks.NewPeterer(t)
	serviceMock.On("Update", in).Return(out)

	pet := NewPet(serviceMock)

	s := httptest.NewServer(http.HandlerFunc(pet.UpdateExistinPet))
	defer s.Close()

	resp, err := http.Post(s.URL, "application/json", bytes.NewBuffer(reqJSON))
	if err != nil {
		t.Fatal(testReqErr, err.Error())
	}
	defer resp.Body.Close()

	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestPetController_FindByStatus_BadRequest(t *testing.T) {
	pet := PetController{
		Responder: &responder.Respond{},
	}

	rr := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)

	pet.FindByStatus(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestFindByStatus(t *testing.T) {
	serviceMock := mocks.NewPeterer(t)
	pets := []models.Pet{testPet}
	serviceMock.On("FindByStatus", []string{"available"}).Return(pets)

	pet := NewPet(serviceMock)

	rr := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/?status=available", nil)
	if err != nil {
		t.Fatal(testReqErr, err.Error())
	}

	pet.FindByStatus(rr, req)

	var resp []models.Pet
	if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
		t.Fatal(testRespDecodeErr)
	}

	assert.Equal(t, pets, resp)
}

func TestPetController_GetByID_BadRequest(t *testing.T) {
	pet := PetController{
		Responder: &responder.Respond{},
	}

	rr := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/pet/i", nil)

	r := chi.NewRouter()
	r.MethodFunc("GET", "/pet/{petId}", pet.GetByID)

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestPetController_GetByID_NotFound(t *testing.T) {
	in := service.GetByIDIn{
		PetID: id,
	}
	out := service.GetByIDOut{
		Error: fmt.Errorf("not found"),
	}

	serviceMock := mocks.NewPeterer(t)
	serviceMock.On("GetByID", in).Return(out)

	pet := NewPet(serviceMock)

	rr := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/pet/1", nil)

	r := chi.NewRouter()
	r.MethodFunc("GET", "/pet/{petId}", pet.GetByID)

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code)
}

func TestPetController_GetByID(t *testing.T) {
	in := service.GetByIDIn{
		PetID: id,
	}
	out := service.GetByIDOut{
		Pet: testPet,
	}

	serviceMock := mocks.NewPeterer(t)
	serviceMock.On("GetByID", in).Return(out)

	pet := NewPet(serviceMock)

	rr := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/pet/1", nil)

	r := chi.NewRouter()
	r.MethodFunc("GET", "/pet/{petId}", pet.GetByID)

	r.ServeHTTP(rr, req)

	var resp models.Pet
	if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
		t.Fatal(testRespDecodeErr, err.Error())
	}

	assert.Equal(t, testPet, resp)
}

func TestPetController_UpdateByID_BadRequest(t *testing.T) {
	pet := PetController{
		Responder: &responder.Respond{},
	}

	rr := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/pet/i", nil)

	r := chi.NewRouter()
	r.MethodFunc("GET", "/pet/{petId}", pet.UpdateByID)

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestPetController_UpdateByID_NotFound(t *testing.T) {
	expect := models.ApiResponse{
		Code:    http.StatusNotFound,
		Type:    "error",
		Message: "not found",
	}
	in := service.UpdateByIDIn{
		PetID:  id,
		Name:   "",
		Status: "",
	}

	serviceMock := mocks.NewPeterer(t)
	serviceMock.On("UpdateByID", in).Return(fmt.Errorf("not found"))

	pet := NewPet(serviceMock)

	rr := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/pet/1", nil)

	r := chi.NewRouter()
	r.MethodFunc("POST", "/pet/{petId}", pet.UpdateByID)

	r.ServeHTTP(rr, req)

	var resp models.ApiResponse
	if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
		t.Fatal(testRespDecodeErr, err.Error())
	}

	assert.Equal(t, expect, resp)
}

func TestPetController_UpdateByID(t *testing.T) {
	expect := models.ApiResponse{
		Code:    200,
		Type:    "unknown",
		Message: "1",
	}
	in := service.UpdateByIDIn{
		PetID:  id,
		Name:   "",
		Status: "",
	}

	serviceMock := mocks.NewPeterer(t)
	serviceMock.On("UpdateByID", in).Return(nil)

	pet := NewPet(serviceMock)

	rr := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/pet/1", nil)

	r := chi.NewRouter()
	r.MethodFunc("POST", "/pet/{petId}", pet.UpdateByID)

	r.ServeHTTP(rr, req)

	var resp models.ApiResponse
	if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
		t.Fatal(testRespDecodeErr, err.Error())
	}

	assert.Equal(t, expect, resp)
}

func TestPetController_DeletePet_BadRequest(t *testing.T) {
	pet := PetController{
		Responder: &responder.Respond{},
	}

	rr := httptest.NewRecorder()
	req := httptest.NewRequest("DELETE", "/pet/d", nil)

	r := chi.NewRouter()
	r.MethodFunc("DELETE", "/pet/{petId}", pet.DeletePet)

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestPetController_DeletePet_NotFound(t *testing.T) {
	serviceMock := mocks.NewPeterer(t)
	serviceMock.On("Delete", id).Return(fmt.Errorf("not found"))

	pet := NewPet(serviceMock)

	rr := httptest.NewRecorder()
	req := httptest.NewRequest("DELETE", "/pet/1", nil)

	r := chi.NewRouter()
	r.MethodFunc("DELETE", "/pet/{petId}", pet.DeletePet)

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code)
}

func TestPetController_DeletePet(t *testing.T) {
	expect := models.ApiResponse{
		Code:    200,
		Type:    "unknown",
		Message: "1",
	}

	serviceMock := mocks.NewPeterer(t)
	serviceMock.On("Delete", id).Return(nil)

	pet := NewPet(serviceMock)

	rr := httptest.NewRecorder()
	req := httptest.NewRequest("DELETE", "/pet/1", nil)

	r := chi.NewRouter()
	r.MethodFunc("DELETE", "/pet/{petId}", pet.DeletePet)

	r.ServeHTTP(rr, req)

	var resp models.ApiResponse
	if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
		t.Fatal(testRespDecodeErr, err.Error())
	}

	assert.Equal(t, expect, resp)
}
