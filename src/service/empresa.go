package service

import (
	"context"
	"database/sql"
	"syssoftintegra-api/src/database"
	"syssoftintegra-api/src/model"
)

var contx_empresa = context.Background()

func ObtenerEmpresa() (model.Empresa, string) {
	empresa := model.Empresa{}

	// Obtiene la conexión actual
	db, err := database.CreateConnection()
	if err != nil {
		return empresa, err.Error()
	}
	defer db.Close()

	query := `SELECT 
	IdEmpresa,
	NumeroDocumento,
	RazonSocial,
	NombreComercial
	FROM EmpresaTB`
	row := db.QueryRowContext(contx_empresa, query)
	err = row.Scan(&empresa.IdEmpresa, &empresa.NumeroDocumento, &empresa.RazonSocial, &empresa.NombreComercial)

	if err == sql.ErrNoRows {
		return empresa, "empty"
	}
	if err != nil {
		return empresa, err.Error()
	}

	return empresa, "ok"
}
