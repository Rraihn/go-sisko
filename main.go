package main

import (
	"github.com/Rraihn/go-sisko/app"
	"github.com/Rraihn/go-sisko/controller"
	"github.com/Rraihn/go-sisko/helper"
	"github.com/Rraihn/go-sisko/middleware"
	"github.com/Rraihn/go-sisko/repository"
	"github.com/Rraihn/go-sisko/service"
	"github.com/go-playground/validator"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	siswaRepository := repository.NewSiswaRepo()
	siswaService := service.NewSiswaService(siswaRepository, db, validate)
	siswaController := controller.NewSiswaController(siswaService)

	router := app.NewRouter(siswaController)

	server := http.Server{
		Addr:    "localhost:3030",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
