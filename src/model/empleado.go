package model

type Empleado struct {
	Id              int    `json:"id,omitempty"`
	IdEmpleado      string `json:"idEmpleado,omitempty"`
	TipoDocumento   int    `json:"tipoDocumento,omitempty"`
	NumeroDocumento string `json:"numeroDocumento,omitempty"`
	Apellidos       string `json:"apellidos,omitempty"`
	Nombres         string `json:"nombres,omitempty"`
	Sexo            int    `json:"sexo,omitempty"`
	FechaNacimiento string `json:"fechaNacimiento,omitempty"`
	Puesto          int    `json:"puesto,omitempty"`
	IdRol           int    `json:"idRol,omitempty"`
	Estado          int    `json:"estado,omitempty"`
	Telefono        string `json:"telefono,omitempty"`
	Celular         string `json:"celular,omitempty"`
	Email           string `json:"email,omitempty"`
	Direccion       string `json:"direccion,omitempty"`
	Usuario         string `json:"usuario,omitempty"`
	Clave           string `json:"clave,omitempty"`
	Sistema         bool   `json:"sistema,omitempty"`
	Huella          string `json:"huella,omitempty"`

	Rol     *Rol     `json:"rol,omitempty"`
	Detalle *Detalle `json:"detalle,omitempty"`
}
