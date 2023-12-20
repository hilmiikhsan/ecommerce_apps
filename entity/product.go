package entity

import (
	"errors"

	"github.com/ecommerce/dto"
	"github.com/google/uuid"
)

var (
	ErrProductNameIsRequired = errors.New("name is required")
	ErrDescriptionIsRequired = errors.New("description is required")
	ErrPriceIsRequired       = errors.New("price is required")
	ErrPriceIsInvalid        = errors.New("price is invalid")
	ErrStockIsRequired       = errors.New("stock is required")
	ErrStockIsInvalid        = errors.New("stock is invalid")
	ErrCategoryIdIsRequired  = errors.New("category_id is required")
	ErrImageUrlIsRequired    = errors.New("image_url is required")
	ErrInvalidRole           = errors.New("invalid role")
	ErrCategoryNotFound      = errors.New("category_id is not found")
	ErrProductNotFound       = errors.New("product not found in this resources")
)

type Product struct {
	ID           int     `db:"id"`
	Name         string  `db:"name"`
	Description  string  `db:"description"`
	Price        int     `db:"price"`
	Stock        int     `db:"stock"`
	CategoryId   int     `db:"category_id"`
	MerchantId   int     `db:"merchant_id"`
	ImageUrl     string  `db:"image_url"`
	Sku          string  `db:"sku"`
	Category     string  `db:"category"`
	MerchantName string  `db:"merchant_name"`
	MerchantCity string  `db:"merchant_city"`
	TotalData    int     `db:"total_data"`
	CreatedBy    string  `db:"created_by"`
	CreatedAt    string  `db:"created_at"`
	UpdatedAt    *string `db:"updated_at"`
}

func NewProduct() Product {
	return Product{}
}

func (p Product) Validate(req dto.CreateOrUpdateProductRequest, id string) (Product, error) {
	if req.Name == "" {
		return p, ErrProductNameIsRequired
	}

	if req.Description == "" {
		return p, ErrDescriptionIsRequired
	}

	if req.Price == 0 {
		return p, ErrPriceIsRequired
	}

	if req.Price < 0 {
		return p, ErrPriceIsInvalid
	}

	if req.Stock == 0 {
		return p, ErrStockIsRequired
	}

	if req.Stock < 0 {
		return p, ErrStockIsInvalid
	}

	if req.CategoryId == 0 {
		return p, ErrCategoryIdIsRequired
	}

	if req.ImageUrl == "" {
		return p, ErrImageUrlIsRequired
	}

	p.ID = req.ID
	p.Name = req.Name
	p.Description = req.Description
	p.Price = req.Price
	p.Stock = req.Stock
	p.CategoryId = req.CategoryId
	p.ImageUrl = req.ImageUrl
	p.Sku = uuid.New().String()
	p.CreatedBy = id

	return p, nil
}

func (p Product) CheckUserRole(role string) (err error) {
	if role != "merchant" {
		return ErrInvalidRole
	}

	return
}

func (p Product) ProductResponse(products []Product) []dto.GetListProductResponse {
	responses := []dto.GetListProductResponse{}

	for _, product := range products {
		response := dto.GetListProductResponse{
			ID:          product.ID,
			Sku:         product.Sku,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			Stock:       product.Stock,
			Category:    product.Category,
			ImageUrl:    product.ImageUrl,
		}

		responses = append(responses, response)
	}

	return responses
}

func (p Product) ProductDetailResponse(product Product) dto.GetDetailProductResponse {
	response := dto.GetDetailProductResponse{
		ID:          product.ID,
		Sku:         product.Sku,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
		Category:    product.Category,
		CategoryId:  product.CategoryId,
		ImageUrl:    product.ImageUrl,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   p.NullStringScan(product.UpdatedAt),
	}

	return response
}

func (p Product) ProductDetailUserPerspectiveResponse(product Product) dto.GetDetailProductUserPerspectiveResponse {
	response := dto.GetDetailProductUserPerspectiveResponse{
		ID:          product.ID,
		Sku:         product.Sku,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
		Category:    product.Category,
		CategoryId:  product.CategoryId,
		Merchant: dto.Merchant{
			ID:   product.MerchantId,
			Name: product.MerchantName,
			City: product.MerchantCity,
		},
		ImageUrl:  product.ImageUrl,
		CreatedAt: product.CreatedAt,
		UpdatedAt: p.NullStringScan(product.UpdatedAt),
	}

	return response
}

func (p Product) NullStringScan(value *string) string {
	if value == nil {
		return ""
	}
	return *value
}
