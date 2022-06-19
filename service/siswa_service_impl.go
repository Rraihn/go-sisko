package service

import (
	"context"
	"database/sql"
	"github.com/Rraihn/go-sisko/helper"
	"github.com/Rraihn/go-sisko/helper/model"
	"github.com/Rraihn/go-sisko/model/domain"
	"github.com/Rraihn/go-sisko/model/web"
	"github.com/Rraihn/go-sisko/repository"
	"github.com/go-playground/validator"
)

type SiswaServiceImpl struct {
	SiswaRepository repository.SiswaRepo
	DB              *sql.DB
	Validate        *validator.Validate
}

func NewSiswaService(SiswaRepo repository.SiswaRepo, DB *sql.DB, validate *validator.Validate) SiswaService {
	return &SiswaServiceImpl{
		SiswaRepository: SiswaRepo,
		DB:              DB,
		Validate:        validate,
	}
}

func (s SiswaServiceImpl) CreateSiswa(ctx context.Context, request web.SiswaCreateRequest) web.SiswaResponse {
	err := s.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	siswa := domain.Siswa{
		Nama:          request.Nama,
		Alamat:        request.Alamat,
		TanggalLahir:  request.TanggalLahir,
		TempatLahir:   request.TempatLahir,
		JenisKelamin:  request.JenisKelamin,
		Agama:         request.Agama,
		GolonganDarah: request.GolonganDarah,
		NoTelpon:      request.NoTelepon,
	}

	return model.ToSiswaResponse(siswa)
}

func (s SiswaServiceImpl) UpdateSiswa(ctx context.Context, request web.SiswaUpdateRequest) web.SiswaResponse {
	//TODO implement me
	panic("implement me")
}

func (s SiswaServiceImpl) DeleteSiswa(ctx context.Context, categoryId int) {
	//TODO implement me
	panic("implement me")
}

func (s SiswaServiceImpl) FindSiswaById(ctx context.Context, categoryId int) web.SiswaResponse {
	//TODO implement me
	panic("implement me")
}

func (s SiswaServiceImpl) FindAllSiswa(ctx context.Context) []web.SiswaResponse {
	//TODO implement me
	panic("implement me")
}
