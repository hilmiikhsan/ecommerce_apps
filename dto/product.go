package dto

type CreateOrUpdateProductRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
	CategoryId  int    `json:"category_id"`
	ImageUrl    string `json:"image_url"`
}

type UpdateProductRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
	CategoryId  int    `json:"category_id"`
	ImageUrl    string `json:"image_url"`
}

type GetListProductResponse struct {
	ID          int    `json:"id"`
	Sku         string `json:"sku"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
	Category    string `json:"category"`
	ImageUrl    string `json:"image_url"`
}

type PaginationResponse struct {
	Query     string `json:"query"`
	Limit     int    `json:"limit"`
	Page      int    `json:"page"`
	TotalPage int    `json:"total_page"`
}

type GetDetailProductResponse struct {
	ID          int    `json:"id"`
	Sku         string `json:"sku"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
	Category    string `json:"category"`
	CategoryId  int    `json:"category_id"`
	ImageUrl    string `json:"image_url"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type GetDetailProductUserPerspectiveResponse struct {
	ID          int      `json:"id"`
	Sku         string   `json:"sku"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Price       int      `json:"price"`
	Stock       int      `json:"stock"`
	Category    string   `json:"category"`
	CategoryId  int      `json:"category_id"`
	Merchant    Merchant `json:"merchant"`
	ImageUrl    string   `json:"image_url"`
	CreatedAt   string   `json:"created_at"`
	UpdatedAt   string   `json:"updated_at"`
}

func CountTotalPage(total, limit int) int {
	if limit == 0 {
		return 0
	}
	return (total + limit - 1) / limit
}

func NewPaginationResponse(queryParam string, limit, page, totalData int) PaginationResponse {
	return PaginationResponse{
		Query:     queryParam,
		Limit:     limit,
		Page:      page,
		TotalPage: CountTotalPage(totalData, limit),
	}
}
