package service

import (
	"context"
	"database/sql"
	"fmt"
	"syssoftintegra-api/src/database"
	"syssoftintegra-api/src/model"
)

var contx_rol = context.Background()

func ListarRoles() (model.Rol, string) {
	rol := model.Rol{}

	db, err := database.CreateConnection()
	if err != nil {
		fmt.Println(err.Error())
		return rol, "Se cerro la conexión."
	}

	defer db.Close()

	query := "SELECT IdRol,Nombre,Sistema FROM RolTB"
	row := db.QueryRowContext(contx_rol, query)

	err = row.Scan(&rol.IdRol, &rol.Nombre, &rol.Sistema)
	if err == sql.ErrNoRows || err != nil {
		fmt.Println(err.Error())
		return rol, "No pudo leer las filas requeridas."
	}

	return rol, "ok"
}
