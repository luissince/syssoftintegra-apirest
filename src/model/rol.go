package model

type Rol struct {
	Id      int    `json:"id,omitempty"`
	IdRol   int    `json:"idRol,omitempty"`
	Nombre  string `json:"nombre,omitempty"`
	Sistema bool   `json:"sistema,omitempty"`
}
