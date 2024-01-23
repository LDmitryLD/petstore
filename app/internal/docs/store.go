package docs

import "projects/LDmitryLD/petstore/app/internal/models"

// swagger:route POST /store/order store CreateOrderRequest
// Разместить заказ питомца.
// responses:
// 	200: CreateOrderResponse
// 	400: description: Invalid Order.

// swagger:parameters CreateOrderRequest
type CreateOrderRequest struct {
	// Разместить заказ на покупку питомца
	//
	// required: true
	// in:body
	Body models.Order `json:"body"`
}

// swagger:response CreateOrderResponse
type CreateOrderResponse struct {
	// successful operation
	//
	// in:body
	Body models.Order `json:"body"`
}

// swagger:route GET /store/order/{orderId} store GetOrderByIDRequest
// Найти заказ на покупку питомца по ID.
// responses:
// 	200: GetOrderByIDResponse
// 	400: description: Invalid ID supplied.
// 	404: description: Order not found.

// swagger:parameters GetOrderByIDRequest
type GetOrderByIDRequest struct {
	// ID заказа
	//
	// required: true
	// in:path
	OrderID int `json:"orderId"`
}

// swagger:response GetOrderByIDResponse
type GetOrderByIDResponse struct {
	// in:body
	Body models.Order `json:"body"`
}

// swagger:route DELETE /store/order/{orderId} store DeleteOrderRequest
// Удаление заказа по ID.
// Используйте только положительные целые числа
// responses:
// 	200: DeleteOrderResponse
// 	400: description: Invalid ID supplied.
// 	404: description: Order not found.

// swagger:parameters DeleteOrderRequest
type DeleteOrderRequest struct {
	// ID заказа который нужно удалить
	//
	// required: true
	// in:path
	OrderID int `json:"orderId"`
}

// swagger:response DeleteOrderResponse
type DeleteOrderResponse struct {
	// in:body
	Body models.ApiResponse `json:"body"`
}

// swagger:route GET /store/inventory store GetInventoryRequest
// Получить список заказов по статусу.
// Возвращает сопоставление статусов с их колличеством
// responses:
// 	200: GetInventoryResponse

// swagger:response GetInventoryResponse
type GetInventoryResponse struct {
	// in:body
	Inventory map[string]interface{} `json:"inventory"`
}
