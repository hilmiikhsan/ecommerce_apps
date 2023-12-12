package product

import (
	"context"
	"database/sql"

	"github.com/ecommerce/domain/category"
	"github.com/ecommerce/domain/merchant"
	"github.com/ecommerce/dto"
	"github.com/ecommerce/entity"
)

type ProductService struct {
	repository         Repository
	merchantRepository merchant.Repository
	categoryRepository category.Repository
}

func NewProductService(repository Repository, merchantRepository merchant.Repository, categoryRepository category.Repository) ProductService {
	return ProductService{
		repository:         repository,
		merchantRepository: merchantRepository,
		categoryRepository: categoryRepository,
	}
}

func (p ProductService) CreateProduct(ctx context.Context, req entity.Product, token string) (err error) {
	merchant, err := p.merchantRepository.GetByCreatedBy(ctx, token)
	if err != nil {
		return
	}
	req.MerchantId = merchant.ID

	if err = entity.NewProduct().CheckUserRole(merchant.Role); err != nil {
		return
	}

	if _, err = p.categoryRepository.GetById(ctx, req.CategoryId); err != nil {
		if sql.ErrNoRows == err {
			err = entity.ErrCategoryNotFound
		}
		return
	}

	if err = p.repository.Create(ctx, req); err != nil {
		return
	}

	return
}

func (p ProductService) GetListProduct(ctx context.Context, token, queryParam string, limit, page int) (response []dto.GetListProductResponse, totalData int, err error) {
	merchant, err := p.merchantRepository.GetByCreatedBy(ctx, token)
	if err != nil {
		return
	}

	if err = entity.NewProduct().CheckUserRole(merchant.Role); err != nil {
		return
	}

	products, totalData, err := p.repository.GetByMerchantId(ctx, queryParam, limit, page, merchant.ID)
	if err != nil {
		return
	}

	response = entity.NewProduct().ProductResponse(products)

	return
}

func (p ProductService) GetDetailProduct(ctx context.Context, id int, token string) (response dto.GetDetailProductResponse, err error) {
	merchant, err := p.merchantRepository.GetByCreatedBy(ctx, token)
	if err != nil {
		return
	}

	if err = entity.NewProduct().CheckUserRole(merchant.Role); err != nil {
		return
	}

	product, err := p.repository.GetById(ctx, id)
	if err != nil {
		if sql.ErrNoRows == err {
			err = entity.ErrProductNotFound
		}
		return
	}

	response = entity.NewProduct().ProductDetailResponse(product)

	return
}

func (p ProductService) UpdateProduct(ctx context.Context, req entity.Product, token string) (err error) {
	merchant, err := p.merchantRepository.GetByCreatedBy(ctx, token)
	if err != nil {
		return
	}

	if err = entity.NewProduct().CheckUserRole(merchant.Role); err != nil {
		return
	}

	if _, err = p.repository.GetById(ctx, req.ID); err != nil {
		if sql.ErrNoRows == err {
			err = entity.ErrProductNotFound
		}
		return
	}

	if _, err = p.categoryRepository.GetById(ctx, req.CategoryId); err != nil {
		if sql.ErrNoRows == err {
			err = entity.ErrCategoryNotFound
		}
		return
	}

	if err = p.repository.Update(ctx, req); err != nil {
		return
	}

	return
}

func (p ProductService) GetDetailProductUserPerspective(ctx context.Context, sku string) (response dto.GetDetailProductUserPerspectiveResponse, err error) {
	product, err := p.repository.GetBySku(ctx, sku)
	if err != nil {
		if sql.ErrNoRows == err {
			err = entity.ErrProductNotFound
		}
		return
	}

	response = entity.NewProduct().ProductDetailUserPerspectiveResponse(product)

	return
}
