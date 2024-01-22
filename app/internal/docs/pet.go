package docs

import (
	"io"
	"projects/LDmitryLD/petstore/app/internal/models"
)

// swagger:route POST /pet/{petId}/uploadImage pet UploadImageParametes
// загрузка изображения.
// responses:
// 	200: ApiResponse

// swagger:parameters UploadImageParametes
type UploadImageParametes struct {
	// ID питомца
	//
	// in:path
	// requierd:true
	PetID int `json:"petId"`

	// Дополнительная информация для передачи на сервер
	//
	// in:formData
	AdditionalMetadata string `json:"additional_metadata"`

	// файл для загрузки
	//
	// in:formData
	// required:true
	// swagger:file
	File *io.ReadCloser `json:"file"`
}

// swagger:response ApiResponse
type ApiResponse struct {
	// in:body
	Body models.ApiResponse `json:"body"`
}

// swagger:route POST /pet pet PetAddRequest
// Добавление нового питомца в магазин.
// responses:
// 	200: PetAddResponse

// swagger:parameters PetAddRequest
type PetAddRequest struct {
	// Питомец которого надо добавить в магазин
	//
	// in:body
	Body models.Pet `json:"body"`
}

// swagger:response PetAddResponse
type PetAddResponse struct {
	// in:body
	Body models.Pet `json:"body"`
}

// swagger:route PUT /pet pet PetUpdateRequest
// Обновление существующего питомца.
// responses:
// 	200: PetUpdateResponse

// swagger:parameters PetUpdateRequest
type PetUpdateRequest struct {
	// Питомец которого надо обновить
	//
	// in:body
	Body models.Pet `json:"body"`
}

// swagger:response PetUpdateResponse
type PetUpdateResponse struct {
	// in:body
	Body models.Pet `json:"body"`
}

// swagger:route GET /pet/findByStatus pet GetByStatusRequest
// Найти питомца по статусу.
// Несколько значений статуса могут быть представлены строками, разделенными запятыми
// responses:
// 	200: GetByStatusResponse

// swagger:parameters GetByStatusRequest
type GetByStatusRequest struct {
	// Значения статуса, которые необходимо учитывать при фильтрации
	// Доступные значения : available, pending, sold
	//
	// required: true
	// items.Enum: available, pending, sold
	// in:query
	Status []string `json:"status"`
}

// swagger:response GetByStatusResponse
type GetByStatusResponse struct {
	// in:body
	Body []models.Pet `json:"body"`
}

// swagger:route GET /pet/{petId} pet GetByIDRequest
// Найти питомца по ID.
// Возвращает одного питомца
// responses:
// 	200: GetByIDResponse

// swagger:parameters GetByIDRequest
type GetByIDRequest struct {
	// ID питомца
	//
	// in:path
	// required: true
	PetID int `json:"petId"`
}

// swagger:response GetByIDResponse
type GetByIDResponse struct {
	// in:body
	Body models.Pet `json:"body"`
}

// swagger:route POST /pet/{petId} pet UpdateByIDRequest
// Обновить питомца в магазине по ID.
// responses:
// 	200: UpdateByIDResponse

// swagger:parameters UpdateByIDRequest
type UpdateByIDRequest struct {
	// ID питомца которого нужно обновить
	//
	// required: true
	// in:path
	PetID int `json:"petId"`
	// Новое имя питомца
	//
	// in:formData
	Name string `json:"name"`
	// Новый статус питомца
	//
	// in:formData
	Status string `json:"status"`
}

// swagger:response UpdateByIDResponse
type UpdateByIDResponse struct {
	// in:body
	Body models.ApiResponse `json:"body"`
}

// swagger:route DELETE /pet/{petId} pet DeletePetRequest
// Удалить питомца.
// responses:
// 	404: description: Pet not found.

// swagger:parameters DeletePetRequest
type DeletePetRequest struct {
	// in:header
	APIKey string `json:"api_key"`
	// ID питомца
	//
	// required: true
	// in:path
	PetID int `json:"petId"`
}
