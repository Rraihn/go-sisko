package web

type SiswaCreateRequest struct {
	Nama          string `validate:"required,min=1,max=100" json:"name"`
	Alamat        string `validate:"required,min=1,max=200" json:"alamat"`
	TanggalLahir  string `validate:"required,min=1,max=40" json:"tanggalLahir"`
	TempatLahir   string `validate:"required,min=1,max=50" json:"tempatLahir"`
	JenisKelamin  string `validate:"required,min=1,max=10" json:"jenisKelamin"`
	Agama         string `validate:"required,min=1,max=20" json:"agama"`
	GolonganDarah string `validate:"required,min=1,max=2" json:"golonganDarah"`
	NoTelepon     string `validate:"required,min=1,max=20" json:"noTelepon"`
}
