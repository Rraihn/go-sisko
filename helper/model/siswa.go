package model

import (
	"github.com/Rraihn/go-sisko/model/domain"
	"github.com/Rraihn/go-sisko/model/web"
)

func ToSiswaResponse(s domain.Siswa) web.SiswaResponse {
	return web.SiswaResponse{
		Id:           s.Id,
		Nama:         s.Nama,
		Alamat:       s.Alamat,
		TanggalLahir: s.Alamat,
		TempatLahir:  s.TempatLahir,
		JenisKelamin: s.JenisKelamin,
		Agama:        s.Agama,
		NoTelpon:     s.NoTelpon,
	}
}

func ToSiswaResponses(s []domain.Siswa) []web.SiswaResponse {
	var siswaResponses []web.SiswaResponse
	for _, siswa := range s {
		siswaResponses = append(siswaResponses, ToSiswaResponse(siswa))
	}
	return siswaResponses
}
