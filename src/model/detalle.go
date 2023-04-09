package model

type Detalle struct {
	IdDetalle       int    `json:"idDetalle,omitempty"`
	IdMantenimiento string `json:"idMantenimiento,omitempty"`
	IdAuxiliar      string `json:"idAuxiliar,omitempty"`
	Nombre          string `json:"nombre,omitempty"`
	Descripcion     string `json:"descripcion,omitempty"`
	Estado          string `json:"estado,omitempty"`
	UsuarioRegistro string `json:"usuarioRegistro,omitempty"`
}
