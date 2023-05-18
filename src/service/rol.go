package service

import (
	"context"
	"fmt"
	"syssoftintegra-api/src/database"
	"syssoftintegra-api/src/model"
)

var contx_rol = context.Background()

func ListarRoles() ([]model.Rol, string) {
	roles := []model.Rol{}

	db, err := database.CreateConnection()
	if err != nil {
		fmt.Println(err.Error())
		return roles, "Se cerro la conexión."
	}

	defer db.Close()

	query := "SELECT IdRol,Nombre,Sistema FROM RolTB"
	rows, _ := db.QueryContext(contx_rol, query)

	defer rows.Close()

	// err = row.Scan(&rol.IdRol, &rol.Nombre, &rol.Sistema)
	// if err == sql.ErrNoRows || err != nil {
	// 	fmt.Println(err.Error())
	// 	return roles, "No pudo leer las filas requeridas."
	// }

	for rows.Next() {
		rol := model.Rol{}

		rows.Scan(
			&rol.IdRol,
			&rol.Nombre,
			&rol.Sistema,
		)

		roles = append(roles, rol)
	}

	return roles, "ok"
}
