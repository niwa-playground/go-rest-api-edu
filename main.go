package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator"
	"github.com/imniwa/go-rest-api-edu/app"
	"github.com/imniwa/go-rest-api-edu/controller"
	"github.com/imniwa/go-rest-api-edu/helper"
	repository "github.com/imniwa/go-rest-api-edu/repositories"
	service "github.com/imniwa/go-rest-api-edu/services"
	"github.com/julienschmidt/httprouter"
)

func main() {
	db := app.NewDB()
	validate := validator.New()

	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	server := http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
