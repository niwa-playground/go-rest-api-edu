package services

import (
	"context"
	"database/sql"

	"github.com/go-playground/validator"
	"github.com/imniwa/go-rest-api-edu/helper"
	"github.com/imniwa/go-rest-api-edu/model/entity"
	"github.com/imniwa/go-rest-api-edu/model/request"
	"github.com/imniwa/go-rest-api-edu/model/response"
	repository "github.com/imniwa/go-rest-api-edu/repositories"
)

type CategoryService struct {
	Repository repository.BaseCategoryRepository
	DB         *sql.DB
	Validate   *validator.Validate
}

func NewCategoryService(CategoryRepository repository.BaseCategoryRepository, DB *sql.DB, Validate *validator.Validate) BaseCategoryService {
	return &CategoryService{
		Repository: CategoryRepository,
		DB:         DB,
		Validate:   Validate,
	}
}

func (service *CategoryService) Create(ctx context.Context, request request.CategoryCreateRequest) response.CategoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category := entity.Category{
		Name: request.Name,
	}

	category = service.Repository.Save(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}

func (service *CategoryService) Update(ctx context.Context, request request.CategoryUpdateRequest) response.CategoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.Repository.FindById(ctx, tx, request.Id)
	helper.PanicIfError(err)

	category.Name = request.Name

	category = service.Repository.Update(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}

func (service *CategoryService) Delete(ctx context.Context, categoryId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	_, err = service.Repository.FindById(ctx, tx, categoryId)
	helper.PanicIfError(err)

	service.Repository.Delete(ctx, tx, categoryId)
}

func (service *CategoryService) FindById(ctx context.Context, categoryId int) response.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := service.Repository.FindById(ctx, tx, categoryId)
	helper.PanicIfError(err)

	return helper.ToCategoryResponse(category)
}

func (service *CategoryService) FindAll(ctx context.Context) []response.CategoryResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	categories := service.Repository.FindAll(ctx, tx)

	return helper.ToCategoryResponses(categories)
}
