package service

import (
	"context"
	"database/sql"
	"fmt"
	"syssoftintegra-api/src/database"
	"syssoftintegra-api/src/model"
)

var contx_detalle = context.Background()

func ListaDetalleIdMantenimiento(opcion string, idMantenimiento string, nombre string) ([]model.Detalle, string) {
	detalles := []model.Detalle{}

	db, err := database.CreateConnection()
	if err != nil {
		fmt.Println(err.Error())
		return detalles, "Se cerro la conexión."
	}

	defer db.Close()

	query := "Sp_Listar_Detalle_IdMantenimiento @Opcion,@IdMantenimiento,@Nombre"
	rows, err := db.QueryContext(
		contx_detalle,
		query,
		sql.Named("Opcion", opcion),
		sql.Named("IdMantenimiento", idMantenimiento),
		sql.Named("Nombre", nombre),
	)

	if err == sql.ErrNoRows || err != nil {
		fmt.Println(err.Error())
		return detalles, "No se pudo ejecutar el procedimiento."
	}

	defer rows.Close()

	for rows.Next() {
		detalle := model.Detalle{}

		rows.Scan(
			&detalle.IdDetalle,
			&detalle.Nombre,
			&detalle.IdAuxiliar,
		)

		detalles = append(detalles, detalle)
	}

	return detalles, "ok"
}
