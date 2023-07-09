package service

import (
	"context"
	"database/sql"
	"syssoftintegra-api/src/database"
	"syssoftintegra-api/src/model"
)

var contx_suministro = context.Background()

func GetAllSuministro(opcion int, clave string, nombreMarca string, categoria int, marca int, posicionPagina int, filasPorPagina int) ([]model.Suministro, int, string) {
	suministros := []model.Suministro{}

	db, err := database.CreateConnection()
	if err != nil {
		return nil, 0, err.Error()
	}
	defer db.Close()

	queryStoreOne := `exec Sp_Listar_Suministros @opcion, @clave, @nombreMarca, @categoria, @marca, @posicionPagina, @filasPorPagina`
	rows, err := db.QueryContext(
		contx_suministro,
		queryStoreOne,
		sql.Named("opcion", opcion),
		sql.Named("clave", clave),
		sql.Named("nombreMarca", nombreMarca),
		sql.Named("categoria", categoria),
		sql.Named("marca", marca),
		sql.Named("posicionPagina", posicionPagina),
		sql.Named("filasPorPagina", filasPorPagina),
	)
	if err == sql.ErrNoRows {
		return nil, 0, "empty"
	}
	if err != nil {
		return nil, 0, err.Error()
	}
	defer rows.Close()

	count := 0

	for rows.Next() {
		suministro := model.Suministro{}

		imp := model.Impuesto{}
		suministro.Impuesto = &imp

		detUnidadCompra := model.Detalle{}
		suministro.DetalleUnidadCompra = &detUnidadCompra

		detMarca := model.Detalle{}
		suministro.DetalleMarca = &detMarca

		detCategoria := model.Detalle{}
		suministro.DetalleCategoria = &detCategoria

		detEstado := model.Detalle{}
		suministro.DetalleEstado = &detEstado

		count++
		suministro.Id = count + posicionPagina

		err := rows.Scan(
			&suministro.IdSuministro,
			&suministro.Clave,
			&suministro.ClaveAlterna,
			&suministro.NombreMarca,
			&suministro.NombreGenerico,
			&suministro.StockMinimo,
			&suministro.StockMaximo,
			&suministro.Cantidad,
			&suministro.DetalleUnidadCompra.Nombre,
			&suministro.DetalleMarca.Nombre,
			&suministro.PrecioCompra,
			&suministro.PrecioVentaGeneral,
			&suministro.DetalleCategoria.Nombre,
			&suministro.DetalleEstado.Nombre,
			&suministro.Inventario,
			&suministro.ValorInventario,
			&suministro.Imagen,
			&suministro.NuevaImagen,
			&suministro.Impuesto.Nombre,
			&suministro.Impuesto.Valor,
		)
		if err != nil {
			return nil, 0, err.Error()
		}

		suministros = append(suministros, suministro)
	}

	var total int
	queryStoreTwo := `exec Sp_Listar_Suministros_Count @opcion, @clave, @nombreMarca, @categoria, @marca`
	row := db.QueryRowContext(
		contx_suministro,
		queryStoreTwo,
		sql.Named("opcion", opcion),
		sql.Named("clave", clave),
		sql.Named("nombreMarca", nombreMarca),
		sql.Named("categoria", categoria),
		sql.Named("marca", marca),
	)

	err = row.Scan(&total)
	if err == sql.ErrNoRows {
		return nil, 0, "empty"
	}
	if err != nil {
		return nil, 0, err.Error()
	}

	return suministros, total, "ok"
}

func GetSuministroById(idSuministro string) (model.Suministro, string) {

	suministro := model.Suministro{}

	db, err := database.CreateConnection()
	if err != nil {
		return suministro, err.Error()
	}
	defer db.Close()

	query := `SELECT
		IdSuministro,
		Origen,
		Clave,
		ClaveAlterna,
		NombreMarca,
		NombreGenerico,
		Categoria,
		Marca,
		Presentacion,
		UnidadCompra,
		UnidadVenta,
		Estado,
		StockMinimo,
		StockMaximo,
		Cantidad,
		Impuesto,
		TipoPrecio,
		PrecioCompra,
		PrecioVentaGeneral,
		Lote,
		Inventario,
		ValorInventario,
		Imagen,
		ClaveSat,
		ISNULL(NuevaImagen, '') as NuevaImagen
		Descripcion
		FROM SuministroTB WHERE IdSuministro = @idSuministro`

	row := db.QueryRowContext(contx_suministro, query, sql.Named("IdSuministro", idSuministro))
	err = row.Scan(
		&suministro.IdSuministro,
		&suministro.Origen,
		&suministro.Clave,
		&suministro.ClaveAlterna,
		&suministro.NombreMarca,
		&suministro.NombreGenerico,
		&suministro.IdDetalleCategoria,
		&suministro.IdDetalleMarca,
		&suministro.Presentacion,
		&suministro.IdDetalleUnidadCompra,
		&suministro.UnidadVenta,
		&suministro.IdDetalleEstado,
		&suministro.StockMinimo,
		&suministro.StockMaximo,
		&suministro.Cantidad,
		&suministro.IdImpuesto,
		&suministro.TipoPrecio,
		&suministro.PrecioCompra,
		&suministro.PrecioVentaGeneral,
		&suministro.Lote,
		&suministro.Inventario,
		&suministro.ValorInventario,
		&suministro.Imagen,
		&suministro.ClaveSat,
		&suministro.NuevaImagen,
		&suministro.Descripcion,
	)
	if err == sql.ErrNoRows {
		return suministro, "empty"
	}
	if err != nil {
		return suministro, err.Error()
	}

	return suministro, "ok"
}

func InsertSuministro(suministro *model.Suministro) string {
	// Obtén la conexión de la base de datos
	db, err := database.CreateConnection()
	if err != nil {
		return "No se puedo establecer conexión, intente nuevamente en un par de minutos."
	}

	// Cerramos la consulta al final de la transacción "defer libera la consulta aunque falla la ejecución"
	defer db.Close()

	// Inicia la transacción
	tx, err := db.BeginTx(contx_suministro, nil)
	if err != nil {
		tx.Rollback()
		return "No se pudo crear el contexto para la transacción."
	}

	// Se obtener el código único
	var idSuministro string
	queryCodAlfa := `SELECT dbo.Fc_Suministro_Codigo_Alfanumerico()`
	row := db.QueryRowContext(contx_suministro, queryCodAlfa)
	err = row.Scan(&idSuministro)

	if err != nil {
		tx.Rollback()
		return "No fue posible obtener el código de suministro."
	}

	// Consulta para registrar
	query := `INSERT INTO SuministroTB (
				IdSuministro,
				Origen,
				Clave,
				ClaveAlterna,
				NombreMarca,
				NombreGenerico,
				Categoria,
				Marca,
				Presentacion,
				UnidadCompra,
				UnidadVenta,
				Estado,
				StockMinimo,
				StockMaximo,
				Cantidad,
				Impuesto,
				TipoPrecio,
				PrecioCompra,
				PrecioVentaGeneral,
				Lote,
				Inventario,
				ValorInventario,
				Imagen,
				ClaveSat,
				NuevaImagen,
				Descripcion)
				VALUES (
				@IdSuministro,
				@Origen,
				@Clave,
				@ClaveAlterna,
				@NombreMarca,
				@NombreGenerico,
				@Categoria,
				@Marca,
				@Presentacion,
				@UnidadCompra,
				@UnidadVenta,
				@Estado,
				@StockMinimo,
				@StockMaximo,
				@Cantidad,
				@Impuesto,
				@TipoPrecio,
				@PrecioCompra,
				@PrecioVentaGeneral,
				@Lote,
				@Inventario,
				@ValorInventario,
				@Imagen,
				@ClaveSat,
				@NuevaImagen,
				@Descripcion)`

	// Ejecuta la consulta dentro de la transacción
	_, err = tx.ExecContext(
		contx_suministro,
		query,
		sql.Named("IdSuministro", idSuministro),
		sql.Named("Origen", suministro.Origen),
		sql.Named("Clave", suministro.Clave),
		sql.Named("ClaveAlterna", suministro.ClaveAlterna),
		sql.Named("NombreMarca", suministro.NombreMarca),
		sql.Named("NombreGenerico", suministro.NombreGenerico),
		sql.Named("Categoria", suministro.IdDetalleCategoria),
		sql.Named("Marca", suministro.IdDetalleMarca),
		sql.Named("Presentacion", suministro.Presentacion),
		sql.Named("UnidadCompra", suministro.IdDetalleUnidadCompra),
		sql.Named("UnidadVenta", suministro.UnidadVenta),
		sql.Named("Estado", suministro.IdDetalleEstado),
		sql.Named("StockMinimo", suministro.StockMinimo),
		sql.Named("StockMaximo", suministro.StockMaximo),
		sql.Named("Cantidad", suministro.Cantidad),
		sql.Named("Impuesto", suministro.IdImpuesto),
		sql.Named("TipoPrecio", suministro.TipoPrecio),
		sql.Named("PrecioCompra", suministro.PrecioCompra),
		sql.Named("PrecioVentaGeneral", suministro.PrecioVentaGeneral),
		sql.Named("Lote", suministro.Lote),
		sql.Named("Inventario", suministro.Inventario),
		sql.Named("ValorInventario", suministro.ValorInventario),
		sql.Named("Imagen", suministro.Imagen),
		sql.Named("ClaveSat", suministro.ClaveSat),
		sql.Named("NuevaImagen", suministro.NuevaImagen),
		sql.Named("Descripcion", suministro.Descripcion),
	)

	// Si ocurre un error, haz un rollback de la transacción.
	if err != nil {
		tx.Rollback()
		return "No se pudo registrar los datos."
	}

	// Si toda ha sido bien, haz commit de la transacción
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return "El genero un problema al guardar la transacción."
	}

	return "insert"
}

func UpdateSuministro(suministro *model.Suministro) string {
	db, err := database.CreateConnection()
	if err != nil {
		return err.Error()
	}

	defer db.Close()

	tx, err := db.BeginTx(contx_suministro, nil)
	if err != nil {
		tx.Rollback()
		return err.Error()
	}

	query := `UPDATE SuministroTB SET
			Origen = @Origen,
			Clave = @Clave,
			ClaveAlterna = @ClaveAlterna,
			NombreMarca = @NombreMarca,
			NombreGenerico = @NombreGenerico,
			Categoria = @Categoria,
			Marca = @Marca,
			Presentacion = @Presentacion,
			UnidadCompra = @UnidadCompra,
			UnidadVenta = @UnidadVenta,
			Estado = @Estado,
			StockMinimo = @StockMinimo,
			StockMaximo = @StockMaximo,
			Cantidad = @Cantidad,
			Impuesto = @Impuesto,
			TipoPrecio = @TipoPrecio,
			PrecioCompra = @PrecioCompra,
			PrecioVentaGeneral = @PrecioVentaGeneral,
			Lote = @Lote,
			Inventario = @Inventario,
			ValorInventario = @ValorInventario,
			Imagen = @Imagen,
			ClaveSat = @ClaveSat,
			NuevaImagen = @NuevaImagen,
			Descripcion = @Descripcion)
			WHERE IdSuministro = @IdSuministro`

	result, err := tx.ExecContext(
		contx_suministro,
		query,
		sql.Named("Origen", suministro.Origen),
		sql.Named("Clave", suministro.Clave),
		sql.Named("ClaveAlterna", suministro.ClaveAlterna),
		sql.Named("NombreMarca", suministro.NombreMarca),
		sql.Named("NombreGenerico", suministro.NombreGenerico),
		sql.Named("Categoria", suministro.IdDetalleCategoria),
		sql.Named("Marca", suministro.IdDetalleMarca),
		sql.Named("Presentacion", suministro.Presentacion),
		sql.Named("UnidadCompra", suministro.IdDetalleUnidadCompra),
		sql.Named("UnidadVenta", suministro.UnidadVenta),
		sql.Named("Estado", suministro.IdDetalleEstado),
		sql.Named("StockMinimo", suministro.StockMinimo),
		sql.Named("StockMaximo", suministro.StockMaximo),
		sql.Named("Cantidad", suministro.Cantidad),
		sql.Named("Impuesto", suministro.IdImpuesto),
		sql.Named("TipoPrecio", suministro.TipoPrecio),
		sql.Named("PrecioCompra", suministro.PrecioCompra),
		sql.Named("PrecioVentaGeneral", suministro.PrecioVentaGeneral),
		sql.Named("Lote", suministro.Lote),
		sql.Named("Inventario", suministro.Inventario),
		sql.Named("ValorInventario", suministro.ValorInventario),
		sql.Named("Imagen", suministro.Imagen),
		sql.Named("ClaveSat", suministro.ClaveSat),
		sql.Named("NuevaImagen", suministro.NuevaImagen),
		sql.Named("Descripcion", suministro.Descripcion),
		sql.Named("IdSuministro", suministro.IdSuministro),
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

	// Si toda ha sido bien, haz commit de la transacción
	tx.Commit()

	return "update"
}

func DeleteSuministro(idSuministro string) string {

	db, err := database.CreateConnection()
	if err != nil {
		return err.Error()
	}
	defer db.Close()

	tx, err := db.BeginTx(contx_suministro, nil)
	if err != nil {
		tx.Rollback()
		return err.Error()
	}

	query := `DELETE FROM SuministroTB WHERE IdSuministro = @idSuministro`
	result, err := tx.ExecContext(contx_suministro, query, sql.Named("IdSuministro", idSuministro))
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
