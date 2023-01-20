package services

import (
	"context"

	"github.com/imniwa/go-rest-api-edu/model/request"
	"github.com/imniwa/go-rest-api-edu/model/response"
)

type BaseCategoryService interface {
	Create(ctx context.Context, request request.CategoryCreateRequest) response.CategoryResponse
	Update(ctx context.Context, request request.CategoryUpdateRequest) response.CategoryResponse
	Delete(ctx context.Context, categoryId int)
	FindById(ctx context.Context, categoryId int) response.CategoryResponse
	FindAll(ctx context.Context) []response.CategoryResponse
}
