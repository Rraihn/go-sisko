package web

type SiswaUpdateRequest struct {
	Id            int    `validate:"required"`
	Nama          string `validate:"required,max=200,min=1" json:"nama"`
	Alamat        string `validate:"required,max=40,min=1" json:"alamat"`
	TanggalLahir  string `validate:"required,max=20,min=1" json:"tanggalLahir"`
	TempatLahir   string `validate:"required,max=40,min=1" json:"tempatLahir"`
	JenisKelamin  string `validate:"required,max=10,min=1" json:"jenisKelamin"`
	Agama         string `validate:"required,max=20,min=1" json:"agama"`
	GolonganDarah string `valifate:"required,max=2,min=1" json:"golonganDarah"`
	NoTelepon     string `validate:"required,max=20,min=1" json:"noTelepon"`
}
