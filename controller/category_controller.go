package controller

import (
	"net/http"
	"strconv"

	"github.com/imniwa/go-rest-api-edu/helper"
	modelRequest "github.com/imniwa/go-rest-api-edu/model/request"
	"github.com/imniwa/go-rest-api-edu/model/response"
	service "github.com/imniwa/go-rest-api-edu/services"
	"github.com/julienschmidt/httprouter"
)

type CategoryController struct {
	Service service.CategoryService
}

func (controller *CategoryController) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryCreateRequest := modelRequest.CategoryCreateRequest{}
	helper.ReadFromRequestBody(request, &categoryCreateRequest)

	categoryResponse := controller.Service.Create(request.Context(), categoryCreateRequest)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryController) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	CategoryUpdateRequest := modelRequest.CategoryUpdateRequest{}
	helper.ReadFromRequestBody(request, &CategoryUpdateRequest)

	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	CategoryUpdateRequest.Id = id

	categoryResponse := controller.Service.Update(request.Context(), CategoryUpdateRequest)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryController) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	controller.Service.Delete(request.Context(), id)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryController) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	categoryResponse := controller.Service.FindById(request.Context(), id)
	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryController) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryResponses := controller.Service.FindAll(request.Context())
	webResponse := response.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponses,
	}
	helper.WriteToResponseBody(writer, webResponse)
}
