package service

import (
	"syssoftintegra-api/src/database"
	"syssoftintegra-api/src/model"

	"context"
	"database/sql"
)

var contx_impuesto = context.Background()

func GetAllImpuesto(posicionPagina int, filasPorPagina int) ([]model.Impuesto, int, string) {

	impuestos := []model.Impuesto{}

	db, err := database.CreateConnection()
	if err != nil {
		return nil, 0, err.Error()
	}
	defer db.Close()

	queryStoreOne := `exec Sp_Listar_Impuestos @posicionPagina, @filasPorPagina`
	rows, err := db.QueryContext(contx_impuesto, queryStoreOne, sql.Named("posicionPagina", posicionPagina), sql.Named("filasPorPagina", filasPorPagina))
	if err == sql.ErrNoRows {
		return nil, 0, "empty"
	}
	if err != nil {
		return nil, 0, err.Error()
	}
	defer rows.Close()

	count := 0
	for rows.Next() {
		impuesto := model.Impuesto{}

		detOperacion := model.Detalle{}
		impuesto.DetalleOperacion = &detOperacion

		count++
		impuesto.Id = count + posicionPagina

		err := rows.Scan(
			&impuesto.IdImpuesto,
			&impuesto.DetalleOperacion.Nombre,
			&impuesto.Nombre,
			&impuesto.Valor,
			&impuesto.Predeterminado,
			&impuesto.Codigo,
			&impuesto.Sistema,
		)
		if err != nil {
			return nil, 0, err.Error()
		}
		impuestos = append(impuestos, impuesto)
	}

	var total int
	queryStoreTwo := `exec Sp_Listar_Impuestos_Count`
	row := db.QueryRowContext(contx_impuesto, queryStoreTwo)
	err = row.Scan(&total)
	if err == sql.ErrNoRows {
		return nil, 0, "empty"
	}
	if err != nil {
		return nil, 0, err.Error()
	}

	return impuestos, total, "ok"

}

func GetImpuestoById(id int) (model.Impuesto, string) {
	impuesto := model.Impuesto{}

	db, err := database.CreateConnection()
	if err != nil {
		return impuesto, err.Error()
	}
	defer db.Close()

	query := `SELECT TOP 1 
			IdImpuesto,
			Operacion,
			Nombre,
			Valor,
			Codigo,
			Numeracion,
			NombreImpuesto,
			Letra,
			Categoria,
			Predeterminado,
			Sistema 
			FROM ImpuestoTB WHERE IdImpuesto = @IdImpuesto`
	row := db.QueryRowContext(contx_impuesto, query, sql.Named("IdImpuesto", id))
	err = row.Scan(
		&impuesto.IdImpuesto,
		&impuesto.IdDetalleOperacion,
		&impuesto.Nombre,
		&impuesto.Valor,
		&impuesto.Codigo,
		&impuesto.Numeracion,
		&impuesto.NombreImpuesto,
		&impuesto.Letra,
		&impuesto.Categoria,
		&impuesto.Predeterminado,
		&impuesto.Sistema,
	)
	if err == sql.ErrNoRows {
		return impuesto, "empty"
	}
	if err != nil {
		return impuesto, err.Error()
	}

	return impuesto, "ok"
}

func InsertUpdateImpuesto(impuesto *model.Impuesto) string {

	db, err := database.CreateConnection()
	if err != nil {
		return err.Error()
	}
	defer db.Close()

	if impuesto.IdImpuesto == 0 {
		// INSERT
		tx, err := db.BeginTx(contx_impuesto, nil)
		if err != nil {
			tx.Rollback()
			return err.Error()
		}

		query := `INSERT INTO ImpuestoTB (
			Operacion,
			Nombre,
			Valor,
			Codigo,
			Numeracion,
			NombreImpuesto,
			Letra,
			Categoria,
			Predeterminado,
			Sistema)
			VALUES (
			@Operacion,
			@Nombre,
			@Valor,
			@Codigo,
			@Numeracion,
			@NombreImpuesto,
			@Letra,
			@Categoria,
			@Predeterminado,
			@Sistema)`
		result, err := tx.ExecContext(
			contx_impuesto,
			query,
			sql.Named("Operacion", impuesto.IdDetalleOperacion),
			sql.Named("Nombre", impuesto.Nombre),
			sql.Named("Valor", impuesto.Valor),
			sql.Named("Codigo", impuesto.Codigo),
			sql.Named("Numeracion", impuesto.Numeracion),
			sql.Named("NombreImpuesto", impuesto.NombreImpuesto),
			sql.Named("Letra", impuesto.Letra),
			sql.Named("Categoria", impuesto.Categoria),
			sql.Named("Predeterminado", impuesto.Predeterminado),
			sql.Named("Sistema", impuesto.Sistema),
		)
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

		return "insert"

	} else {
		// UPDATE
		tx, err := db.BeginTx(contx_impuesto, nil)
		if err != nil {
			tx.Rollback()
			return err.Error()
		}

		query := `UPDATE ImpuestoTB SET 
				Operacion = @Operacion,
				Nombre = @Nombre,
				Valor = @Valor,
				Codigo = @Codigo,
				Numeracion = @Numeracion,
				NombreImpuesto = @NombreImpuesto,
				Letra = @Letra,
				Categoria = @Categoria,
				Predeterminado = @Predeterminado,
				Sistema = @Sistema
				WHERE IdImpuesto = @IdImpuesto`
		result, err := tx.ExecContext(
			contx_impuesto,
			query,
			sql.Named("Operacion", impuesto.IdDetalleOperacion),
			sql.Named("Nombre", impuesto.Nombre),
			sql.Named("Valor", impuesto.Valor),
			sql.Named("Codigo", impuesto.Codigo),
			sql.Named("Numeracion", impuesto.Numeracion),
			sql.Named("NombreImpuesto", impuesto.NombreImpuesto),
			sql.Named("Letra", impuesto.Letra),
			sql.Named("Categoria", impuesto.Categoria),
			sql.Named("Predeterminado", impuesto.Predeterminado),
			sql.Named("Sistema", impuesto.Sistema),
			sql.Named("IdImpuesto", impuesto.IdImpuesto),
		)
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

		return "update"
	}

}

func DeleteImpuesto(id int) string {

	db, err := database.CreateConnection()
	if err != nil {
		return err.Error()
	}
	defer db.Close()

	tx, err := db.BeginTx(contx_impuesto, nil)
	if err != nil {
		tx.Rollback()
		return err.Error()
	}

	query := `DELETE FROM ImpuestoTB WHERE IdImpuesto = @IdImpuesto`
	result, err := tx.ExecContext(contx_impuesto, query, sql.Named("IdImpuesto", id))
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
