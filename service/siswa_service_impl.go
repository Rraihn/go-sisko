package service

import (
	"context"
	"database/sql"
	"github.com/Rraihn/go-sisko/helper"
	"github.com/Rraihn/go-sisko/helper/helper_model"
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
	siswa = s.SiswaRepository.SaveSiswa(ctx, tx, siswa)
	return helper_model.ToSiswaResponse(siswa)
}

func (s SiswaServiceImpl) UpdateSiswa(ctx context.Context, request web.SiswaUpdateRequest) web.SiswaResponse {
	err := s.Validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	siswa, err := s.SiswaRepository.FindSiswaById(ctx, tx, request.Id)
	if err != nil {
		helper.PanicIfError(err)
	}

	siswa.Nama = request.Nama
	siswa.Alamat = request.Alamat
	siswa.TanggalLahir = request.TanggalLahir
	siswa.TempatLahir = request.TempatLahir
	siswa.JenisKelamin = request.JenisKelamin
	siswa.Agama = request.Agama
	siswa.GolonganDarah = request.GolonganDarah
	siswa.NoTelpon = request.NoTelepon

	siswa = s.SiswaRepository.UpdateSiswa(ctx, tx, siswa)

	return helper_model.ToSiswaResponse(siswa)
}

func (s SiswaServiceImpl) DeleteSiswa(ctx context.Context, categoryId int) {
	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	siswa, err := s.SiswaRepository.FindSiswaById(ctx, tx, categoryId)
	if err != nil {
		helper.PanicIfError(err)
	}

	s.SiswaRepository.DeleteSiswa(ctx, tx, siswa)
}

func (s SiswaServiceImpl) FindSiswaById(ctx context.Context, categoryId int) web.SiswaResponse {
	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	siswa, err := s.SiswaRepository.FindSiswaById(ctx, tx, categoryId)
	if err != nil {
		helper.PanicIfError(err)
	}

	return helper_model.ToSiswaResponse(siswa)
}

func (s SiswaServiceImpl) FindAllSiswa(ctx context.Context) []web.SiswaResponse {
	tx, err := s.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollBack(tx)

	siswaa := s.SiswaRepository.FindAllSiswa(ctx, tx)

	return helper_model.ToSiswaResponses(siswaa)
}
