package dto

// @Description Create subscription request
type CreateSubscriptionReq struct {
	UserId      string `json:"user_id" binding:"required"`
	ServiceName string `json:"service_name" binding:"required"`
	Price       int    `json:"price" binding:"required"`
	StartDate   string `json:"start_date" binding:"required"`
	EndDate     string `json:"end_date"`
}

// @Description Create subscription response
type CreateSubscriptionResp struct {
	Message string `json:"message"`
}

// @Description Get cost request
type GetCostReq struct {
	UserId      string `json:"user_id" binding:"required"`
	ServiceName string `json:"service_name" binding:"required"`
	StartDate   string `json:"start_date" binding:"required"`
	EndDate     string `json:"end_date" binding:"required"`
}

// @Description Get cost response
type GetCostResp struct {
	Total int `json:"total_cost"`
}
