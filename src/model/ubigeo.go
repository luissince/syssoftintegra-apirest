package model

type Ubigeo struct {
	IdUbigeo     int    `json:"id,omitempty"`
	Ubigeo       string `json:"ubigeo,omitempty"`
	Departamento string `json:"departamento,omitempty"`
	Provincia    string `json:"provincia,omitempty"`
	Distrito     string `json:"distrito,omitempty"`
}
