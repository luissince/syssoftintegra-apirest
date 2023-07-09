package model

type Impuesto struct {
	Id int `json:"id,omitempty"`

	IdImpuesto         int     `json:"idImpuesto,omitempty"`
	IdDetalleOperacion int     `json:"idDetalleOperacion,omitempty"`
	Nombre             string  `json:"nombre,omitempty"`
	Valor              float64 `json:"valor,omitempty"`
	Codigo             string  `json:"codigo,omitempty"`
	Numeracion         string  `json:"numeracion,omitempty"`
	NombreImpuesto     string  `json:"nombreImpuesto,omitempty"`
	Letra              string  `json:"letra,omitempty"`
	Categoria          string  `json:"categoria,omitempty"`
	Predeterminado     bool    `json:"predeterminado,omitempty"`
	Sistema            bool    `json:"sistema,omitempty"`

	DetalleOperacion *Detalle `json:"detalleOperacion"`
}
