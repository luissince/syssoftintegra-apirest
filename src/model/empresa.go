package model

type Empresa struct {
	IdEmpresa        int    `json:"idEmpresa,omitempty"`
	GiroComercial    int    `json:"giroComercial,omitempty"`
	Nombre           string `json:"nombre,omitempty"`
	Telefono         string `json:"telefono,omitempty"`
	Celular          string `json:"celular,omitempty"`
	PaginaWeb        string `json:"paginaWeb,omitempty"`
	Email            string `json:"email,omitempty"`
	Domicilio        string `json:"domicilio,omitempty"`
	TipoDocumento    int    `json:"tipoDocumento,omitempty"`
	NumeroDocumento  string `json:"numeroDocumento,omitempty"`
	RazonSocial      string `json:"razonSocial,omitempty"`
	NombreComercial  string `json:"nombreComercial,omitempty"`
	Image            []byte `json:"image,omitempty"`
	ImagenRuta       string `json:"imagenRuta,omitempty"`
	Ubigeo           int    `json:"ubigeo,omitempty"`
	UsuarioSol       string `json:"usuarioSol,omitempty"`
	ClaveSol         string `json:"claveSol,omitempty"`
	CertificadoRuta  string `json:"certificadoRuta,omitempty"`
	CertificadoClave string `json:"certificadoClave,omitempty"`
	Terminos         string `json:"terminos,omitempty"`
	Condiciones      string `json:"condiciones,omitempty"`
	IdApiSunat       string `json:"idApiSunat,omitempty"`
	ClaveApiSunat    string `json:"claveApiSunat,omitempty"`
}
