package dto

type CreateSubscriptionReq struct {
	UserId      string `json:"user_id" binding:"required"`
	ServiceName string `json:"service_name" binding:"required"`
	Price       int    `json:"price" binding:"required"`
	StartDate   string `json:"start_date" binding:"required"`
	EndDate     string `json:"end_date"`
}

type CreateSubscriptionResp struct {
	Message string `json:"message"`
}

type GetCostReq struct {
	UserId      string `json:"user_id" binding:"required"`
	ServiceName string `json:"service_name" binding:"required"`
	StartDate   string `json:"start_date" binding:"required"`
	EndDate     string `json:"end_date" binding:"required"`
}

type GetCostResp struct {
	Total int `json:"total_cost"`
}
