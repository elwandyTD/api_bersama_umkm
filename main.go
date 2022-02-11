package main

import (
	"net/http"

	"github.com/elwandyTD/api_bersama_umkm/app"
	"github.com/elwandyTD/api_bersama_umkm/controller"
	"github.com/elwandyTD/api_bersama_umkm/helper"
	"github.com/elwandyTD/api_bersama_umkm/repository"
	"github.com/elwandyTD/api_bersama_umkm/service"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func main() {
	db := app.NewDB()
	validate := validator.New()

	adminRepository := repository.NewAdminRepository()
	adminService := service.NewAdminService(adminRepository, db, validate)
	adminController := controller.NewAdminController(adminService)

	router := httprouter.New()

	router.POST("/api/v1/admin", adminController.Create)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
