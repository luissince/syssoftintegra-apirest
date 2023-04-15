package model

import (
	"time"
)

type Banco struct {
	IdBanco       string `json:"idBanco,omitempty"`
	NombreCuenta  string `json:"nombreCuenta,omitempty"`
	NumeroCuenta  string `json:"numeroCuenta,omitempty"`
	Moneda        Moneda
	SaldoInicial  float64   `json:"saldoInicial,omitempty"`
	FechaCreacion time.Time `json:"fechaCreacion,omitempty"`
	HoraCreacion  time.Time `json:"horaCreacion,omitempty"`
	Descripcion   string    `json:"descripcion,omitempty"`
	Sistema       bool      `json:"sistema,omitempty"`
	FormaPago     int16     `json:"formaPago,omitempty"`
	Mostrar       bool      `json:"mostrar,omitempty"`
}
