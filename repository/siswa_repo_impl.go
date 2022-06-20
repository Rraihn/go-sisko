package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/Rraihn/go-sisko/helper"
	"github.com/Rraihn/go-sisko/model/domain"
)

type SiswaRepoImpl struct {
}

func NewSiswaRepo() SiswaRepo {
	return &SiswaRepoImpl{}
}

func (s SiswaRepoImpl) SaveSiswa(ctx context.Context, tx *sql.Tx, siswa domain.Siswa) domain.Siswa {
	SQL := "insert into siswa(nama, alamat, tanggal_lahir, tempat_lahir, jenis_kelamin, agama, golongan_darah, no_telepon) values(?, ?, ?, ? ,? ,? ,? ,?)"
	result, err := tx.ExecContext(ctx, SQL, siswa.Nama, siswa.Alamat, siswa.TanggalLahir, siswa.TempatLahir, siswa.JenisKelamin, siswa.Agama, siswa.GolonganDarah, siswa.NoTelpon)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	siswa.Id = int(id)
	return siswa
}

func (s SiswaRepoImpl) UpdateSiswa(ctx context.Context, tx *sql.Tx, siswa domain.Siswa) domain.Siswa {
	SQL := "Update siswa set nama = ?, alamat = ?, tanggal_lahir =?, tempat_lahir=?, jenis_kelamin = ?, agama = ?, golongan_darah = ?, no_telepon = ? where id = ?"
	_, err := tx.ExecContext(ctx, SQL, siswa.Nama, siswa.Alamat, siswa.TanggalLahir, siswa.TempatLahir, siswa.JenisKelamin, siswa.Agama, siswa.GolonganDarah, siswa.NoTelpon, siswa.Id)
	helper.PanicIfError(err)

	return siswa
}

func (s SiswaRepoImpl) DeleteSiswa(ctx context.Context, tx *sql.Tx, siswa domain.Siswa) {
	SQL := "delete from siswa where id = ?"
	_, err := tx.ExecContext(ctx, SQL, siswa.Id)
	helper.PanicIfError(err)
}

func (s SiswaRepoImpl) FindSiswaById(ctx context.Context, tx *sql.Tx, siswaId int) (domain.Siswa, error) {
	SQL := "Select id, name from siswa where id  = ?"
	rows, err := tx.QueryContext(ctx, SQL, siswaId)
	helper.PanicIfError(err)

	siswa := domain.Siswa{}

	if rows.Next() {
		err := rows.Scan(&siswa.Id, &siswa.Nama)
		helper.PanicIfError(err)
		return siswa, nil
	} else {
		return siswa, errors.New("Siswa not found")
	}
}

func (s SiswaRepoImpl) FindAllSiswa(ctx context.Context, tx *sql.Tx) []domain.Siswa {
	SQL := "select id from siswa"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)

	var sisua []domain.Siswa
	for rows.Next() {
		siswa := domain.Siswa{}
		err := rows.Scan(&siswa.Id, &siswa.Nama)
		helper.PanicIfError(err)
		sisua = append(sisua, siswa)
	}
	return sisua
}
