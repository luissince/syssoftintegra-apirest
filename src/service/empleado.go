package service

import (
	"context"
	"database/sql"

	"syssoftintegra-api/src/database"
	"syssoftintegra-api/src/model"
)

func Login(contx context.Context, usuario string, clave string) (model.Empleado, string) {
	empleado := model.Empleado{}

	db, err := database.CreateConnection()
	if err != nil {
		return empleado, "error"
	}
	defer db.Close()

	queryStore := `exec Sp_Validar_Ingreso @usuario, @clave`
	row := db.QueryRowContext(contx, queryStore, sql.Named("usuario", usuario), sql.Named("clave", clave))

	err = row.Scan(&empleado.IdEmpleado, &empleado.Apellidos, &empleado.Nombres, &empleado.Rol.Nombre, &empleado.Estado, &empleado.Estado)
	if err == sql.ErrNoRows || err != nil {
		return empleado, "empty"
	}

	return empleado, "ok"
}

func GetAllEmpleado(contx context.Context, opcion int, search string, posicionPagina int, filasPorPagina int) ([]model.Empleado, int, string) {

	empleados := []model.Empleado{}

	db, err := database.CreateConnection()
	if err != nil {
		return empleados, 0, "error"
	}
	defer db.Close()

	queryStoreOne := `exec Sp_Listar_Empleados @opcion, @search, @posicionPagina, @filasPorPagina`

	rows, err := db.QueryContext(contx, queryStoreOne, sql.Named("opcion", opcion), sql.Named("search", search), sql.Named("posicionPagina", posicionPagina), sql.Named("filasPorPagina", filasPorPagina))
	if err == sql.ErrNoRows || err != nil {
		return empleados, 0, "empty"
	}
	defer rows.Close()

	for rows.Next() {
		empleado := model.Empleado{}

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
			return empleados, 0, "error"
		}
		empleados = append(empleados, empleado)
	}

	var total int
	queryStoreTwo := `exec Sp_Listar_Empleados_Count @posicionPagina, @filasPorPagina`

	row := db.QueryRowContext(contx, queryStoreTwo, sql.Named("posicionPagina", posicionPagina), sql.Named("filasPorPagina", filasPorPagina))
	err = row.Scan(&total)
	if err == sql.ErrNoRows || err != nil {
		return empleados, 0, "empty"
	}

	return empleados, total, "ok"
}

func GetEmpleadoById(contx context.Context, idEmpleado string) (model.Empleado, string) {

	empleado := model.Empleado{}

	db, err := database.CreateConnection()
	if err != nil {
		return empleado, "error"
	}
	defer db.Close()

	query := `SELECT TOP(1)
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

	// query :=`SELECT TOP(1) * FROM EmpleadoTB WHERE IdEmpleado = @idEmpleado`

	row := db.QueryRowContext(contx, query, sql.Named("IdEmpleado", idEmpleado))

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
	if err == sql.ErrNoRows || err != nil {
		return empleado, "empty"
	}

	return empleado, "ok"
}

func IUEmpledo(contx context.Context, empleado *model.Empleado) string {

	db, err := database.CreateConnection()
	if err != nil {
		return "error"
	}
	defer db.Close()

	if empleado.IdEmpleado == "" {

		var newId string
		queryfunc := `SELECT dbo.Fc_Empleado_Codigo_Alfanumerico()`
		row := db.QueryRowContext(contx, queryfunc)

		err = row.Scan(&newId)
		if err != nil {
			return "error"
		}

		tx, err := db.BeginTx(contx, nil)
		if err != nil {
			tx.Rollback()
			return "error"
		}

		query := `INSERT INTO EmpleadoTB (IdEmpleado, TipoDocumento, NumeroDocumento, Apellidos, Nombres, Sexo, FechaNacimiento, Puesto, Rol, Estado, Telefono, Celular, Email, Direccion, Usuario, Clave, Sistema, Huella)
				VALUES (IdEmpleado @TipoDocumento, @NumeroDocumento, @Apellidos, @Nombres, @Sexo, @FechaNacimiento, @Puesto, @Rol, @Estado, @Telefono, @Celular, @Email, @Direccion, @Usuario, @Clave, @Sistema, @Huella)`

		result, err := tx.ExecContext(
			contx,
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
			return "error"
		}

		value, err := result.RowsAffected()
		if err != nil {
			tx.Rollback()
			return "error"
		}

		if value == 0 {
			tx.Rollback()
			return "empty"
		}

		tx.Commit()

		return "insert"

	} else {

		tx, err := db.BeginTx(contx, nil)
		if err != nil {
			tx.Rollback()
			return "error"
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
			contx,
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
			return "error"
		}

		value, err := result.RowsAffected()
		if err != nil {
			tx.Rollback()
			return "error"
		}

		if value == 0 {
			tx.Rollback()
			return "empty"
		}

		tx.Commit()

		return "update"

	}
}

func DeleteEmpleado(contx context.Context, id string) string {

	db, err := database.CreateConnection()
	if err != nil {
		return "error"
	}
	defer db.Close()

	var idEmpleado string
	queryfunc := `SELECT TOP(1) * FROM EmpleadoTB WHERE IdEmpleado = @IdEmpleado AND Sistema = 1`

	row := db.QueryRowContext(contx, queryfunc, sql.Named("IdEmpleado", id))
	err = row.Scan(&idEmpleado)
	if err == sql.ErrNoRows || err != nil {
		return "empty"
	}

	if idEmpleado == "" {
		return "El empleado no puede ser eliminado porque es parte del sistema."
	}

	tx, err := db.BeginTx(contx, nil)
	if err != nil {
		tx.Rollback()
		return "error"
	}

	query := `DELETE FROM EmpleadoTB WHERE IdEmpleado = @IdEmpleado`

	result, err := tx.ExecContext(contx, query, sql.Named("IdEmpleado", id))
	if err != nil {
		tx.Rollback()
		return "error"
	}

	value, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		return "error"
	}

	if value == 0 {
		tx.Rollback()
		return "empty"
	}

	tx.Commit()

	return "delete"
}
