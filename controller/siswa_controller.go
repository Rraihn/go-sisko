package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type SiswaController interface {
	CreateSiswa(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	UpdateSiswa(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	DeleteSiswa(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindSiswaById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindAllSiswa(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
