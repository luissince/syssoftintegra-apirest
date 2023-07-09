package service

import (
	"context"
	"database/sql"
	"syssoftintegra-api/src/database"
	"syssoftintegra-api/src/model"
)

var contx_banco = context.Background()

func GetAllBanco(search string, posicionPagina int, filasPorPagina int) ([]model.Banco, int, string) {
	bancos := []model.Banco{}
	db, err := database.CreateConnection()
	if err != nil {
		return nil, 0, err.Error()
	}
	defer db.Close()

	queryStoreOne := `exec Sp_Listar_Bancos @search, @posicionPagina, @filasPorPagina`
	rows, err := db.QueryContext(contx_banco, queryStoreOne, sql.Named("@search", search), sql.Named("posicionPagina", posicionPagina), sql.Named("filasPorPagina", filasPorPagina))
	if err == sql.ErrNoRows {
		return nil, 0, "empty"
	}
	if err != nil {
		return nil, 0, err.Error()
	}
	defer rows.Close()

	count := 0
	for rows.Next() {
		count++
		banco := model.Banco{}
		banco.Id = count + posicionPagina

		err := rows.Scan(
			&banco.Id,
			&banco.IdBanco,
			&banco.NombreCuenta,
			&banco.NumeroCuenta,
			&banco.IdMoneda,
			&banco.SaldoInicial,
			&banco.FechaCreacion,
			&banco.HoraCreacion,
			&banco.Descripcion,
			&banco.Sistema,
			&banco.FormaPago,
			&banco.Mostrar,
		)
		if err != nil {
			return nil, 0, err.Error()
		}
		bancos = append(bancos, banco)
	}

	var total int
	queryStoreTwo := `exec Sp_Listar_Bancos_Count @search`
	row := db.QueryRowContext(contx_banco, queryStoreTwo, sql.Named("search", search))
	err = row.Scan(&total)
	if err == sql.ErrNoRows {
		return nil, 0, "empty"
	}
	if err != nil {
		return nil, 0, err.Error()
	}

	return bancos, total, "ok"

}

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

func DeleteBanco(id string) string {
	db, err := database.CreateConnection()
	if err != nil {
		return err.Error()
	}
	defer db.Close()

	tx, err := db.BeginTx(contx_banco, nil)
	if err != nil {
		tx.Rollback()
		return err.Error()
	}

	query := `DELETE FROM Banco WHERE IdBanco = @IdBanco`
	result, err := tx.ExecContext(contx_banco, query, sql.Named("IdBanco", id))
	if err != nil {
		tx.Rollback()
		return err.Error()
	}

	value, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		return err.Error()
	}

	if value == 0 {
		tx.Rollback()
		return "empty"
	}

	tx.Commit()

	return "delete"
}

func ValidarSistemaBanco(id string) string {
	db, err := database.CreateConnection()
	if err != nil {
		return err.Error()
	}
	defer db.Close()

	var idBanco, rpta string
	query := `SELECT TOP(1) IdBanco FROM Banco WHERE IdBanco = @IdBanco AND Sistema = 1`
	row := db.QueryRow(query, sql.Named("IdBanco", id))
	err = row.Scan(&idBanco)
	if err == sql.ErrNoRows {
		return "empty"
	}
	if err != nil {
		return err.Error()
	}

	if idBanco == id {
		rpta = "sistema"
	}

	return rpta
}

func ValidarHistorialBanco(id string) string {
	db, err := database.CreateConnection()
	if err != nil {
		return err.Error()
	}
	defer db.Close()

	var idBanco, rpta string
	query := `SELECT IdBanco FROM BancoHistorialTB WHERE IdBanco = @IdBanco`
	row := db.QueryRow(query, sql.Named("IdBanco", id))
	err = row.Scan(&idBanco)
	if err == sql.ErrNoRows {
		return "empty"
	}
	if err != nil {
		return err.Error()
	}

	if idBanco == id {
		rpta = "history"
	}

	return rpta
}
