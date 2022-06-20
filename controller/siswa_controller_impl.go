package controller

import (
	"github.com/Rraihn/go-sisko/helper"
	"github.com/Rraihn/go-sisko/model/web"
	"github.com/Rraihn/go-sisko/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type SiswaControllerImpl struct {
	SiswaService service.SiswaService
}

func NewSiswaController(siswaService service.SiswaService) SiswaController {
	return &SiswaControllerImpl{
		SiswaService: siswaService,
	}
}

func (s SiswaControllerImpl) CreateSiswa(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	siswaCreateRequest := web.SiswaCreateRequest{}
	helper.ReadFromRequestBody(request, &siswaCreateRequest)

	siswaResponse := s.SiswaService.CreateSiswa(request.Context(), siswaCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   siswaResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (s SiswaControllerImpl) UpdateSiswa(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	siswaUpdateRequest := web.SiswaUpdateRequest{}
	helper.ReadFromRequestBody(request, &siswaUpdateRequest)

	siswaId := params.ByName("siswaId")
	id, err := strconv.Atoi(siswaId)
	helper.PanicIfError(err)

	siswaUpdateRequest.Id = id

	siswaResponse := s.SiswaService.UpdateSiswa(request.Context(), siswaUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   siswaResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (s SiswaControllerImpl) DeleteSiswa(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	siswaId := params.ByName("siswaId")
	id, err := strconv.Atoi(siswaId)
	helper.PanicIfError(err)

	s.SiswaService.DeleteSiswa(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (s SiswaControllerImpl) FindSiswaById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	siswaId := params.ByName("siswaId")
	id, err := strconv.Atoi(siswaId)
	helper.PanicIfError(err)

	siswaResponse := s.SiswaService.FindSiswaById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   siswaResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (s SiswaControllerImpl) FindAllSiswa(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	siswaResponse := s.SiswaService.FindAllSiswa(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   siswaResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
