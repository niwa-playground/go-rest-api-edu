package helper

import (
	"github.com/imniwa/go-rest-api-edu/model/entity"
	"github.com/imniwa/go-rest-api-edu/model/response"
)

func ToCategoryResponse(category entity.Category) response.CategoryResponse {
	return response.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func ToCategoryResponses(categories []entity.Category) []response.CategoryResponse {
	var categoryResponse []response.CategoryResponse
	for _, category := range categories {
		categoryResponse = append(categoryResponse, ToCategoryResponse(category))
	}
	return categoryResponse
}
