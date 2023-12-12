package dto

type CreateCategoryRequest struct {
	Name string `json:"name"`
}

type GetListCategoryResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
