package docs

import "projects/LDmitryLD/petstore/app/internal/models"

// swagger:route POST /user/createWithArray user CreateArrayRequest
// Создаёт список пользователей по полученному массиву.
// responses:
// 	200: CreateArrayResponse

// swagger:parameters CreateArrayRequest
type CreateArrayRequest struct {
	// Список объектов user
	//
	// required: true
	// in:body
	Body []models.User `json:"body"`
}

// swagger:response CreateArrayResponse
type CreateArrayResponse struct {
	// in:body
	Body models.ApiResponse `json:"body"`
}

// swagger:route POST /user/createWithList user CreateListRequest
// Создаёт список пользователей по полученному массиву.
// responses:
// 	200: CreateListResponse

// swagger:parameters CreateListRequest
type CreateListRequest struct {
	// Список объектов user
	//
	// required: true
	// in:body
	Body []models.User `json:"body"`
}

// swagger:response CreateListResponse
type CreateListResponse struct {
	// in:body
	Body models.ApiResponse `json:"body"`
}

// swagger:route GET /user/{username} user GetUserRequest
// Получения пользователя по никнейму.
// responses:
// 	200: GetUserResponse

// swagger:parameters GetUserRequest
type GetUserRequest struct {
	// Никнейм пользователя
	//
	// required: true
	// in:path
	Username string `json:"username"`
}

// swagger:response GetUserResponse
type GetUserResponse struct {
	// in:body
	Body models.User `json:"body"`
}

// swagger:route PUT /user/{username} user UpdateUserRequest
// Обновляет пользователя.
// responses:
// 	200: UpdateUserResponse

// swagger:parameters UpdateUserRequest
type UpdateUserRequest struct {
	// Никнейм пользователя которого надо обновить
	//
	// required: true
	// in:path
	Username string `json:"username"`
	// Новый объект user
	//
	// required: true
	// in:body
	Body models.User `json:"body"`
}

// swagger:response UpdateUserResponse
type UpdateUserResponse struct {
	// in:body
	Body models.ApiResponse `json:"body"`
}

// swagger:route DELETE /user/{username} user DeleteUserRequest
// Удаление пользователя.
// responses:
// 	200: DeleteUserResponse

// swagger:parameters DeleteUserRequest
type DeleteUserRequest struct {
	// Никнейм пользователя
	//
	// required: true
	// in:path
	Username string `json:"username"`
}

// swagger:response DeleteUserResponse
type DeleteUserResponse struct {
	// in:body
	Body models.ApiResponse `json:"body"`
}

// swagger:route POST /user user CreateUserRequest
// Создание пользователя.
// responses:
// 	200: CreateUserResponse

// swagger:parameters CreateUserRequest
type CreateUserRequest struct {
	// in:body
	Body models.User `json:"body"`
}

// swagger:response CreateUserResponse
type CreateUserResponse struct {
	// in:body
	Body models.ApiResponse `json:"body"`
}
