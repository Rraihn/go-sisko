package service

import (
	"context"
	"github.com/Rraihn/go-sisko/model/web"
)

type SiswaService interface {
	CreateSiswa(ctx context.Context, request web.SiswaCreateRequest) web.SiswaResponse
	UpdateSiswa(ctx context.Context, request web.SiswaUpdateRequest) web.SiswaResponse
	DeleteSiswa(ctx context.Context, categoryId int)
	FindSiswaById(ctx context.Context, categoryId int) web.SiswaResponse
	FindAllSiswa(ctx context.Context) []web.SiswaResponse
}
