package model

type Suministro struct {
	IdSuministro       string  `json:"idSuministro,omitempty"`
	Origen             int     `json:"origen,omitempty"`
	Clave              string  `json:"clave,omitempty"`
	ClaveAlterna       string  `json:"claveAlterna,omitempty"`
	NombreMarca        string  `json:"nombreMarca,omitempty"`
	NombreGenerico     string  `json:"nombreGenerico,omitempty"`
	Categoria          int     `json:"categoria,omitempty"`
	Marca              int     `json:"marca,omitempty"`
	Presentacion       int     `json:"presentacion,omitempty"`
	UnidadCompra       int     `json:"unidadCompra,omitempty"`
	UnidadVenta        int     `json:"unidadVenta,omitempty"`
	Estado             int     `json:"estado,omitempty"`
	StockMinimo        float64 `json:"stockMinimo,omitempty"`
	StockMaximo        float64 `json:"stockMaximo,omitempty"`
	Cantidad           float64 `json:"cantidad,omitempty"`
	Impuesto           int     `json:"impuesto,omitempty"`
	TipoPrecio         bool    `json:"tipoPrecio,omitempty"`
	PrecioCompra       float64 `json:"precioCompra,omitempty"`
	PrecioVentaGeneral float64 `json:"precioVentaGeneral,omitempty"`
	Lote               bool    `json:"lote,omitempty"`
	Inventario         bool    `json:"inventario,omitempty"`
	ValorInventario    int     `json:"valorInventario,omitempty"`
	Imagen             string  `json:"imagen,omitempty"`
	ClaveSat           string  `json:"claveSat,omitempty"`
	NuevaImagen        []byte  `json:"nuevaImagen,omitempty"`
	Descripcion        string  `json:"descripcion,omitempty"`
}
