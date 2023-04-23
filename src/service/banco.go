package service

import (
	"context"
	"database/sql"
	"syssoftintegra-api/src/database"
	"syssoftintegra-api/src/model"
)

var contx_banco = context.Background()

func GetbancoByID(id string) (model.Banco, string) {
	banco := model.Banco{}

	db, err := database.CreateConnection()
	if err != nil {
		return banco, "error"
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
	FROM Banco WHERE IdBanco = @id`

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
	if err == sql.ErrNoRows || err != nil {
		return banco, "empty"
	}

	return banco, "ok"

}
