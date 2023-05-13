package service

import (
	"context"
	"database/sql"

	"syssoftintegra-api/src/database"
	"syssoftintegra-api/src/model"
)

var contx_empleado = context.Background()

func Login(usuario string, clave string) (model.Empleado, string) {
	empleado := model.Empleado{}
	rol := &model.Rol{}
	empleado.Rol = rol

	db, err := database.CreateConnection()
	if err != nil {
		return empleado, err.Error()
	}
	defer db.Close()

	queryStore := `exec Sp_Validar_Ingreso @usuario, @clave`
	row := db.QueryRowContext(contx_empleado, queryStore, sql.Named("usuario", usuario), sql.Named("clave", clave))

	err = row.Scan(&empleado.IdEmpleado, &empleado.Apellidos, &empleado.Nombres, &empleado.Rol.Nombre, &empleado.Estado, &empleado.Estado)
	if err == sql.ErrNoRows {
		return empleado, "empty"
	}
	if err != nil {
		return empleado, err.Error()
	}

	return empleado, "ok"
}

func GetAllEmpleado(opcion int, search string, posicionPagina int, filasPorPagina int) ([]model.Empleado, int, string) {

	empleados := []model.Empleado{}

	db, err := database.CreateConnection()
	if err != nil {
		return nil, 0, err.Error()
	}
	defer db.Close()

	queryStoreOne := `exec Sp_Listar_Empleados @opcion, @search, @posicionPagina, @filasPorPagina`
	rows, err := db.QueryContext(contx_empleado, queryStoreOne, sql.Named("opcion", opcion), sql.Named("search", search), sql.Named("posicionPagina", posicionPagina), sql.Named("filasPorPagina", filasPorPagina))
	if err == sql.ErrNoRows {
		return nil, 0, "empty"
	}
	if err != nil {
		return nil, 0, err.Error()
	}
	defer rows.Close()

	count := 0

	for rows.Next() {
		empleado := model.Empleado{}

		rol := &model.Rol{}
		empleado.Rol = rol

		detalle := &model.Detalle{}
		empleado.Detalle = detalle

		count++
		empleado.Id = count + posicionPagina

		err := rows.Scan(
			&empleado.IdEmpleado,
			&empleado.NumeroDocumento,
			&empleado.Apellidos,
			&empleado.Nombres,
			&empleado.Telefono,
			&empleado.Celular,
			&empleado.Direccion,
			&empleado.Rol.Nombre,
			&empleado.Detalle.Nombre,
		)
		if err != nil {
			return nil, 0, err.Error()
		}
		empleados = append(empleados, empleado)
	}

	var total int
	queryStoreTwo := `exec Sp_Listar_Empleados_Count @opcion, @search`
	row := db.QueryRowContext(contx_empleado, queryStoreTwo, sql.Named("opcion", opcion), sql.Named("search", search))
	err = row.Scan(&total)
	if err == sql.ErrNoRows {
		return nil, 0, "empty"
	}
	if err != nil {
		return nil, 0, err.Error()
	}

	return empleados, total, "ok"
}

func GetEmpleadoById(idEmpleado string) (model.Empleado, string) {

	empleado := model.Empleado{}

	db, err := database.CreateConnection()
	if err != nil {
		return empleado, err.Error()
	}
	defer db.Close()

	query := `SELECT
		IdEmpleado,
		TipoDocumento,
		NumeroDocumento,
		Apellidos,
		Nombres,
		Sexo,
		FechaNacimiento,
		Puesto,
		Rol,
		Estado,
		Telefono,
		Celular,
		Email,
		Direccion,
		Usuario,
		Clave,
		Sistema,
		ISNULL(Huella, '') as Huella
		FROM EmpleadoTB WHERE IdEmpleado = @idEmpleado`

	row := db.QueryRowContext(contx_empleado, query, sql.Named("IdEmpleado", idEmpleado))
	err = row.Scan(
		&empleado.IdEmpleado,
		&empleado.TipoDocumento,
		&empleado.NumeroDocumento,
		&empleado.Apellidos,
		&empleado.Nombres,
		&empleado.Sexo,
		&empleado.FechaNacimiento,
		&empleado.Puesto,
		&empleado.Rol,
		&empleado.Estado,
		&empleado.Telefono,
		&empleado.Celular,
		&empleado.Email,
		&empleado.Direccion,
		&empleado.Usuario,
		&empleado.Clave,
		&empleado.Sistema,
		&empleado.Huella,
	)
	if err == sql.ErrNoRows {
		return empleado, "empty"
	}
	if err != nil {
		return empleado, err.Error()
	}

	return empleado, "ok"
}

func InsertUpdateEmpledo(empleado *model.Empleado) string {

	db, err := database.CreateConnection()
	if err != nil {
		return err.Error()
	}
	defer db.Close()

	if empleado.IdEmpleado == "" {

		var newId string
		queryCodAlfa := `SELECT dbo.Fc_Empleado_Codigo_Alfanumerico()`
		row := db.QueryRowContext(contx_empleado, queryCodAlfa)
		err = row.Scan(&newId)
		if err == sql.ErrNoRows || err != nil {
			return err.Error()
		}

		tx, err := db.BeginTx(contx_empleado, nil)
		if err != nil {
			tx.Rollback()
			return err.Error()
		}

		query := `INSERT INTO EmpleadoTB (IdEmpleado, TipoDocumento, NumeroDocumento, Apellidos, Nombres, Sexo, FechaNacimiento, Puesto, Rol, Estado, Telefono, Celular, Email, Direccion, Usuario, Clave, Sistema, Huella)
				VALUES (IdEmpleado @TipoDocumento, @NumeroDocumento, @Apellidos, @Nombres, @Sexo, @FechaNacimiento, @Puesto, @Rol, @Estado, @Telefono, @Celular, @Email, @Direccion, @Usuario, @Clave, @Sistema, @Huella)`

		result, err := tx.ExecContext(
			contx_empleado,
			query,
			sql.Named("IdEmpleado", newId),
			sql.Named("TipoDocumento", empleado.TipoDocumento),
			sql.Named("NumeroDocumento", empleado.NumeroDocumento),
			sql.Named("Apellidos", empleado.Apellidos),
			sql.Named("Nombres", empleado.Nombres),
			sql.Named("Sexo", empleado.Sexo),
			sql.Named("FechaNacimiento", empleado.FechaNacimiento),
			sql.Named("Puesto", empleado.Puesto),
			sql.Named("Rol", empleado.Rol),
			sql.Named("Estado", empleado.Estado),
			sql.Named("Telefono", empleado.Telefono),
			sql.Named("Celular", empleado.Celular),
			sql.Named("Email", empleado.Email),
			sql.Named("Direccion", empleado.Direccion),
			sql.Named("Usuario", empleado.Usuario),
			sql.Named("Clave", empleado.Clave),
			sql.Named("Sistema", empleado.Sistema),
			sql.Named("Huella", empleado.Huella),
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

		tx, err := db.BeginTx(contx_empleado, nil)
		if err != nil {
			tx.Rollback()
			return err.Error()
		}

		query := `UPDATE EmpleadoTB SET 
			TipoDocumento =@TipoDocumento, 
			NumeroDocumento =@NumeroDocumento, 
			Apellidos =@Apellidos, 
			Nombres =@Nombres, 
			Sexo =@Sexo, 
			FechaNacimiento =@FechaNacimiento, 
			Puesto =@Puesto, 
			Rol =@Rol, 
			Estado =@Estado,
			Telefono =@Telefono, 
			Celular =@Celular, 
			Email =@Email, 
			Direccion =@Direccion, 
			Usuario =@Usuario, 
			Clave =@Clave, 
			Sistema =@Sistema, 
			Huella =@Huella)
			WHERE IdEmpleado = @IdEmpleado`

		result, err := tx.ExecContext(
			contx_empleado,
			query,
			sql.Named("TipoDocumento", empleado.TipoDocumento),
			sql.Named("NumeroDocumento", empleado.NumeroDocumento),
			sql.Named("Apellidos", empleado.Apellidos),
			sql.Named("Nombres", empleado.Nombres),
			sql.Named("Sexo", empleado.Sexo),
			sql.Named("FechaNacimiento", empleado.FechaNacimiento),
			sql.Named("Puesto", empleado.Puesto),
			sql.Named("Rol", empleado.Rol),
			sql.Named("Estado", empleado.Estado),
			sql.Named("Telefono", empleado.Telefono),
			sql.Named("Celular", empleado.Celular),
			sql.Named("Email", empleado.Email),
			sql.Named("Direccion", empleado.Direccion),
			sql.Named("Usuario", empleado.Usuario),
			sql.Named("Clave", empleado.Clave),
			sql.Named("Sistema", empleado.Sistema),
			sql.Named("Huella", empleado.Huella),
			sql.Named("IdEmpleado", empleado.IdEmpleado),
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

func DeleteEmpleado(id string) string {

	db, err := database.CreateConnection()
	if err != nil {
		return err.Error()
	}
	defer db.Close()

	tx, err := db.BeginTx(contx_empleado, nil)
	if err != nil {
		tx.Rollback()
		return err.Error()
	}

	query := `DELETE FROM EmpleadoTB WHERE IdEmpleado = @IdEmpleado`
	result, err := tx.ExecContext(contx_empleado, query, sql.Named("IdEmpleado", id))
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

func ValidarSistemaEmpledo(id string) string {

	db, err := database.CreateConnection()
	if err != nil {
		return err.Error()
	}
	defer db.Close()

	var idEmpleado, rpta string
	query := `SELECT TOP(1) IdEmpleado FROM EmpleadoTB WHERE IdEmpleado = @IdEmpleado AND Sistema = 1`
	row := db.QueryRow(query, sql.Named("IdEmpleado", id))
	err = row.Scan(&idEmpleado)
	if err == sql.ErrNoRows {
		return "empty"
	}
	if err != nil {
		return err.Error()
	}

	if idEmpleado == id {
		rpta = "sistema"
	}

	return rpta

}
