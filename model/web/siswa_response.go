package web

type SiswaResponse struct {
	Id           int    `json:"id"`
	Nama         string `json:"nama"`
	Alamat       string `json:"alamat"`
	TanggalLahir string `json:"tanggalLahir"`
	TempatLahir  string `json:"tempatLahir"`
	JenisKelamin string `json:"jenisKelamin"`
	Agama        string `json:"agama"`
	NoTelpon     string `json:"noTelpon"`
}
