package main

import (
	"golang-restfull/app"
	"golang-restfull/controller"
	"golang-restfull/helper"
	"golang-restfull/middleware"
	"golang-restfull/repository"
	"golang-restfull/service"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)
	authMiddleware := middleware.NewAuthMiddleware(router)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: authMiddleware,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
