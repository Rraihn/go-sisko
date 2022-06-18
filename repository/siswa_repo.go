package repository

import (
	"context"
	"database/sql"
	"github.com/Rraihn/go-sisko/model/domain"
)

type SiswaRepo interface {
	SaveSiswa(ctx context.Context, tx *sql.Tx, Siswa domain.Siswa) domain.Siswa
	UpdateSiswa(ctx context.Context, tx *sql.Tx, siswa domain.Siswa) domain.Siswa
	DeleteSiswa(ctx context.Context, tx *sql.Tx, siswa domain.Siswa)
	FindSiswaById(ctx context.Context, tx *sql.Tx, siswaId int) (domain.Siswa, error)
	FindAllSiswa(ctx context.Context, tx *sql.Tx) []domain.Siswa
}
