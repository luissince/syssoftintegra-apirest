package model

import (
	"time"
)

type Empleado struct {
	Count int `json:"count,omitempty"`

	IdEmpleado      string    `json:"idEmpleado,omitempty"`
	TipoDocumento   int       `json:"tipoDocumento,omitempty"`
	NumeroDocumento string    `json:"numeroDocumento,omitempty,omitempty"`
	Apellidos       string    `json:"apellidos,omitempty,omitempty"`
	Nombres         string    `json:"nombres,omitempty,omitempty"`
	Sexo            int       `json:"sexo,omitempty,omitempty"`
	FechaNacimiento time.Time `json:"fechaNacimiento,omitempty,omitempty"`
	Puesto          int       `json:"puesto,omitempty,omitempty"`
	IdRol           int       `json:"idRol,omitempty"`
	Estado          int       `json:"estado,omitempty"`
	Telefono        string    `json:"telefono,omitempty"`
	Celular         string    `json:"celular,omitempty"`
	Email           string    `json:"email,omitempty"`
	Direccion       string    `json:"direccion,omitempty"`
	Usuario         string    `json:"usuario,omitempty"`
	Clave           string    `json:"clave,omitempty"`
	Sistema         bool      `json:"sistema,omitempty"`
	Huella          string    `json:"huella,omitempty"`

	Rol     Rol
	Detalle Detalle
	// Huella sql.NullString `json:"huella"`
}
