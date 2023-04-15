package model

type Moneda struct {
	IdMoneda       int     `json:"idMoneda,omitempty"`
	Nombre         string  `json:"nombre,omitempty"`
	Abreviado      string  `json:"abreviado,omitempty"`
	Simbolo        string  `json:"simbolo,omitempty"`
	TipoCambio     float64 `json:"tipoCambio,omitempty"`
	Predeterminado bool    `json:"predeterminado,omitempty"`
	Sistema        bool    `json:"sistema,omitempty"`
}
