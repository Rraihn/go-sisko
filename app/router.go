package app

import (
	"github.com/Rraihn/go-sisko/controller"
	"github.com/Rraihn/go-sisko/exception"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(siswaController controller.SiswaController) *httprouter.Router {
	router := httprouter.New()

	router.GET("/api/siswa", siswaController.FindAllSiswa)
	router.GET("/api/siswa/:siswaId", siswaController.FindSiswaById)
	router.POST("/api/siswa", siswaController.CreateSiswa)
	router.PUT("/api/siswa/:siswaId", siswaController.UpdateSiswa)
	router.DELETE("/api/siswa/:siswaId", siswaController.DeleteSiswa)

	router.PanicHandler = exception.ErrorHandler

	return router
}
