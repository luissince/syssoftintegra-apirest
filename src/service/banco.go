package service

import (
	"context"
	"database/sql"
	"syssoftintegra-api/src/database"
	"syssoftintegra-api/src/model"
)

var contx_banco = context.Background()

func GetBancoById(id string) (model.Banco, string) {
	banco := model.Banco{}

	db, err := database.CreateConnection()
	if err != nil {
		return banco, err.Error()
	}
	defer db.Close()

	query := `SELECT TOP(1)
	IdBanco,
	NombreCuenta,
	NumeroCuenta,
	IdMoneda,            
	SaldoInicial,
	Descripcion,
	FormaPago,
	Mostrar 
	FROM Banco WHERE IdBanco = @IdBanco`

	row := db.QueryRowContext(contx_banco, query, sql.Named("IdBanco", id))

	err = row.Scan(
		&banco.IdBanco,
		&banco.NombreCuenta,
		&banco.NumeroCuenta,
		&banco.Moneda.IdMoneda,
		&banco.SaldoInicial,
		&banco.Descripcion,
		&banco.FormaPago,
		&banco.Mostrar,
	)
	if err == sql.ErrNoRows {
		return banco, "empty"
	}
	if err != nil {
		return banco, err.Error()
	}

	return banco, "ok"

}
