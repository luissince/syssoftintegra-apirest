package model

import (
	"time"
)

type Banco struct {
	Id int `json:"id,omitempty"`

	IdBanco       string    `json:"idBanco,omitempty"`
	NombreCuenta  string    `json:"nombreCuenta,omitempty"`
	NumeroCuenta  string    `json:"numeroCuenta,omitempty"`
	IdMoneda      int       `json:"idMoneda,omitempty"`
	SaldoInicial  float64   `json:"saldoInicial,omitempty"`
	FechaCreacion time.Time `json:"fechaCreacion,omitempty"`
	HoraCreacion  time.Time `json:"horaCreacion,omitempty"`
	Descripcion   string    `json:"descripcion,omitempty"`
	Sistema       bool      `json:"sistema,omitempty"`
	FormaPago     int16     `json:"formaPago,omitempty"`
	Mostrar       bool      `json:"mostrar,omitempty"`

	Moneda *Moneda `json:"Moneda,omitempty"`
}
