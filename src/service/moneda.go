package service

import (
	"context"
	"database/sql"
	"syssoftintegra-api/src/database"
	"syssoftintegra-api/src/model"
)

var contx_moneda = context.Background()

func GetMonedaComboBox() ([]model.Moneda, string) {
	monedas := []model.Moneda{}

	db, err := database.CreateConnection()
	if err != nil {
		return nil, err.Error()
	}
	defer db.Close()

	query := `SELECT IdMoneda, Nombre, Simbolo, Predeterminado FROM MonedaTB`
	rows, err := db.QueryContext(contx_moneda, query)
	if err == sql.ErrNoRows {
		return nil, "empty"
	}
	if err != nil {
		return nil, err.Error()
	}
	defer rows.Close()

	for rows.Next() {
		moneda := model.Moneda{}
		err = rows.Scan(&moneda.IdMoneda, &moneda.Nombre, &moneda.Simbolo, &moneda.Predeterminado)
		if err != nil {
			return nil, err.Error()
		}
		monedas = append(monedas, moneda)
	}

	return monedas, "ok"

}

func GetAllMoneda(opcion int, search string, posicionPagina int, filasPorPagina int) ([]model.Moneda, int, string) {

	monedas := []model.Moneda{}

	db, err := database.CreateConnection()
	if err != nil {
		return nil, 0, err.Error()
	}
	defer db.Close()

	query := `exec Sp_Listar_Monedas @opcion, @search, @posicionPagina, @filasPorPagina`

	rows, err := db.QueryContext(contx_moneda, query, sql.Named("opcion", opcion), sql.Named("search", search), sql.Named("posicionPagina", posicionPagina), sql.Named("filasPorPagina", filasPorPagina))
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

		moneda := model.Moneda{}
		moneda.Count = count

		err := rows.Scan(
			&moneda.IdMoneda,
			&moneda.Nombre,
			&moneda.Abreviado,
			&moneda.Simbolo,
			&moneda.TipoCambio,
			&moneda.Predeterminado,
		)
		if err != nil {
			return nil, 0, err.Error()
		}
		monedas = append(monedas, moneda)
	}

	var total int
	queryTotal := `exec Sp_Listar_Monedas_Count @opcion, @search`

	row := db.QueryRowContext(contx_moneda, queryTotal, sql.Named("opcion", opcion), sql.Named("search", search))
	err = row.Scan(&total)
	if err == sql.ErrNoRows {
		return nil, 0, "empty"
	}
	if err != nil {
		return nil, 0, err.Error()
	}

	return monedas, total, "ok"
}

func GetMonedaById(id int) (model.Moneda, string) {
	moneda := model.Moneda{}

	db, err := database.CreateConnection()
	if err != nil {
		return moneda, err.Error()
	}
	defer db.Close()

	query := `SELECT TOP 1 IdMoneda, Nombre, Abreviado, Simbolo, TipoCambio FROM MonedaTB WHERE IdMoneda = @IdMoneda`
	row := db.QueryRowContext(contx_moneda, query, sql.Named("IdMoneda", id))
	err = row.Scan(&moneda.IdMoneda, &moneda.Nombre, &moneda.Abreviado, &moneda.Simbolo, &moneda.TipoCambio)
	if err == sql.ErrNoRows {
		return moneda, "empty"
	}
	if err != nil {
		return moneda, err.Error()
	}

	return moneda, "ok"
}

func InsertUpdateMoneda(moneda *model.Moneda) string {

	db, err := database.CreateConnection()
	if err != nil {
		return err.Error()
	}
	defer db.Close()

	if moneda.IdMoneda == 0 {
		// INSERT
		tx, err := db.BeginTx(contx_moneda, nil)
		if err != nil {
			//tx.Rollback()
			return err.Error()
		}

		query := `INSERT INTO MonedaTB (Nombre, Abreviado, Simbolo, TipoCambio, Predeterminado, Sistema)
				VALUES (@Nombre, @Abreviado, @Simbolo, @TipoCambio, @Predeterminado, @Sistema)`
		result, err := tx.ExecContext(
			contx_moneda,
			query,
			sql.Named("Nombre", moneda.Nombre),
			sql.Named("Abreviado", moneda.Abreviado),
			sql.Named("Simbolo", moneda.Simbolo),
			sql.Named("TipoCambio", moneda.TipoCambio),
			sql.Named("Predeterminado", moneda.Predeterminado),
			sql.Named("Sistema", moneda.Sistema),
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
		tx, err := db.BeginTx(contx_moneda, nil)
		if err != nil {
			//tx.Rollback()
			return err.Error()
		}

		query := `UPDATE MonedaTB SET Nombre = @Nombre, 
					Abreviado = @Abreviado, 
					Simbolo = @Simbolo, 
					TipoCambio = @TipoCambio, 
					Predeterminado = @Predeterminado, 
					Sistema = @Sistema 
					WHERE IdMoneda = @IdMoneda`
		result, err := tx.ExecContext(
			contx_moneda,
			query,
			sql.Named("Nombre", moneda.Nombre),
			sql.Named("Abreviado", moneda.Abreviado),
			sql.Named("Simbolo", moneda.Simbolo),
			sql.Named("TipoCambio", moneda.TipoCambio),
			sql.Named("Predeterminado", moneda.Predeterminado),
			sql.Named("Sistema", moneda.Sistema),
			sql.Named("IdMoneda", moneda.IdMoneda),
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

func DeleteMoneda(id int) string {

	db, err := database.CreateConnection()
	if err != nil {
		return err.Error()
	}
	defer db.Close()

	tx, err := db.BeginTx(contx_moneda, nil)
	if err != nil {
		//tx.Rollback()
		return err.Error()
	}

	query := `DELETE FROM MonedaTB WHERE IdMoneda = @IdMoneda`
	result, err := tx.ExecContext(contx_moneda, query, sql.Named("IdMoneda", id))
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

func ValidarNombreMoneda(idMoneda int, nombreMoneda string) string {
	nombreConsulta := ""

	db, err := database.CreateConnection()
	if err != nil {
		return err.Error()
	}
	defer db.Close()

	if idMoneda == 0 {
		// INSERT
		query := `SELECT TOP 1 Nombre FROM MonedaTB WHERE Nombre = @Nombre`
		row := db.QueryRow(query, sql.Named("Nombre", nombreMoneda))
		err = row.Scan(&nombreConsulta)
		if err == sql.ErrNoRows {
			return "empty"
		}
		if err != nil {
			return err.Error()
		}

		if nombreMoneda == nombreConsulta {
			nombreConsulta = "exists"
		}

	} else {
		// UPDATE
		query := `SELECT Nombre FROM MonedaTB WHERE IdMoneda <> @IdMoneda AND Nombre = @Nombre`
		row := db.QueryRow(query, sql.Named("IdMoneda", idMoneda), sql.Named("Nombre", nombreMoneda))
		err = row.Scan(&nombreConsulta)
		if err == sql.ErrNoRows {
			return "empty"
		}
		if err != nil {
			return err.Error()
		}

		if nombreMoneda == nombreConsulta {
			nombreConsulta = "exists"
		}

	}

	return nombreConsulta

}

func ValidarSistemaMoneda(id int) string {

	db, err := database.CreateConnection()
	if err != nil {
		return err.Error()
	}
	defer db.Close()

	var (
		idMoneda int
		rpt      string
	)

	queryVerificar := `SELECT TOP 1 IdMoneda FROM MonedaTB WHERE IdMoneda = @IdMoneda AND Sistema = 1`
	row := db.QueryRow(queryVerificar, sql.Named("IdMoneda", id))
	err = row.Scan(&idMoneda)
	if err == sql.ErrNoRows {
		return "empty"
	}
	if err != nil {
		return err.Error()
	}

	if idMoneda == id {
		rpt = "sistema"
	}

	return rpt

}
