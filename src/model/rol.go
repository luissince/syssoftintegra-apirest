package model

type Rol struct {
	IdRol   int    `json:"idRol,omitempty"`
	Nombre  string `json:"nombre,omitempty"`
	Sistema bool   `json:"sistema,omitempty"`
}
