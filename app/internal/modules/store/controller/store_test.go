package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"projects/LDmitryLD/petstore/app/internal/infrastructure/responder"
	"projects/LDmitryLD/petstore/app/internal/models"
	"projects/LDmitryLD/petstore/app/internal/modules/store/service/mocks"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

const (
	testReqErr        = "ошибка при выполнении тестового запроса:"
	testRespDecodeErr = "ошибка при декодировании тестового ответа:"
)

var (
	id    = 1
	order = models.Order{
		ID:     1,
		Status: "status",
	}
)

func TestStoreController_Create_BadRequest(t *testing.T) {
	store := StoreController{
		Responder: &responder.Respond{},
	}
	req := map[string]interface{}{"ID": "1"}
	reqJSON, _ := json.Marshal(req)

	s := httptest.NewServer(http.HandlerFunc(store.Create))
	defer s.Close()

	resp, err := http.Post(s.URL, "application/json", bytes.NewBuffer(reqJSON))
	if err != nil {
		t.Fatal(testReqErr, err.Error())
	}
	defer resp.Body.Close()

	assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
}

func TestStoreController_Create(t *testing.T) {
	reqJSON, _ := json.Marshal(order)

	serviceMock := mocks.NewStoreger(t)
	serviceMock.On("Create", order).Return(order)

	store := NewStore(serviceMock)

	s := httptest.NewServer(http.HandlerFunc(store.Create))
	defer s.Close()

	resp, err := http.Post(s.URL, "application/json", bytes.NewBuffer(reqJSON))
	if err != nil {
		t.Fatal(testReqErr, err.Error())
	}
	defer resp.Body.Close()

	var createResp models.Order
	if err := json.NewDecoder(resp.Body).Decode(&createResp); err != nil {
		t.Fatal(testRespDecodeErr, err.Error())
	}

	assert.Equal(t, order, createResp)
}

func TestStoreController_GetByID_BadRequest(t *testing.T) {
	store := StoreController{
		Responder: &responder.Respond{},
	}

	rr := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/store/d", nil)

	r := chi.NewRouter()
	r.MethodFunc("GET", "/store/{orderId}", store.GetByID)

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestStoreController_GetByID_NotFound(t *testing.T) {
	serviceMock := mocks.NewStoreger(t)
	serviceMock.On("GetByID", id).Return(models.Order{}, fmt.Errorf("not found"))

	store := NewStore(serviceMock)

	rr := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/store/1", nil)

	r := chi.NewRouter()
	r.MethodFunc("GET", "/store/{orderId}", store.GetByID)

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code)
}

func TestStoreController_GetByID(t *testing.T) {
	serviceMock := mocks.NewStoreger(t)
	serviceMock.On("GetByID", id).Return(order, nil)

	store := NewStore(serviceMock)

	rr := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/store/1", nil)

	r := chi.NewRouter()
	r.MethodFunc("GET", "/store/{orderId}", store.GetByID)

	r.ServeHTTP(rr, req)

	var resp models.Order
	if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
		t.Fatal(testRespDecodeErr, err.Error())
	}

	assert.Equal(t, order, resp)
}

func TestStoreController_Delete_BadRequest(t *testing.T) {
	store := StoreController{
		Responder: &responder.Respond{},
	}

	rr := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/store/d", nil)

	r := chi.NewRouter()
	r.MethodFunc("GET", "/store/{orderId}", store.Delete)

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestStoreController_Delete_NotFound(t *testing.T) {
	serviceMock := mocks.NewStoreger(t)
	serviceMock.On("Delete", id).Return(fmt.Errorf("not found"))

	store := NewStore(serviceMock)

	rr := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/store/1", nil)

	r := chi.NewRouter()
	r.MethodFunc("GET", "/store/{orderId}", store.Delete)

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNotFound, rr.Code)
}

func TestStoreController_Delete(t *testing.T) {
	epxect := models.ApiResponse{
		Code:    200,
		Type:    "unknown",
		Message: "1",
	}
	serviceMock := mocks.NewStoreger(t)
	serviceMock.On("Delete", id).Return(nil)

	store := NewStore(serviceMock)

	rr := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/store/1", nil)

	r := chi.NewRouter()
	r.MethodFunc("GET", "/store/{orderId}", store.Delete)

	r.ServeHTTP(rr, req)

	var resp models.ApiResponse
	if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
		t.Fatal(testRespDecodeErr, err.Error())
	}

	assert.Equal(t, epxect, resp)
}

func TestStoreController_Inventory(t *testing.T) {
	inv := map[string]interface{}{"status": float64(10)}

	serviceMock := mocks.NewStoreger(t)
	serviceMock.On("Inventory").Return(inv)

	store := NewStore(serviceMock)

	rr := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/store/inventory", nil)

	r := chi.NewRouter()
	r.MethodFunc("GET", "/store/inventory", store.Inventory)

	r.ServeHTTP(rr, req)

	var resp map[string]interface{}
	if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
		t.Fatal(testRespDecodeErr, err.Error())
	}

	assert.Equal(t, inv, resp)
}
