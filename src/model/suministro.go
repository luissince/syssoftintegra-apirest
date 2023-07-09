package model

type Suministro struct {
	Id                    int     `json:"id,omitempty"`
	IdSuministro          string  `json:"idSuministro,omitempty"`
	Origen                int     `json:"origen,omitempty"`
	Clave                 string  `json:"clave,omitempty"`
	ClaveAlterna          string  `json:"claveAlterna,omitempty"`
	NombreMarca           string  `json:"nombreMarca,omitempty"`
	NombreGenerico        string  `json:"nombreGenerico,omitempty"`
	IdDetalleCategoria    int     `json:"idDetalleCategoria,omitempty"`
	IdDetalleMarca        int     `json:"idDetalleMarca,omitempty"`
	Presentacion          int     `json:"presentacion,omitempty"`
	IdDetalleUnidadCompra int     `json:"idDetalleUnidadCompra,omitempty"`
	UnidadVenta           int     `json:"unidadVenta,omitempty"`
	IdDetalleEstado       int     `json:"idDetalleEstado,omitempty"`
	StockMinimo           float64 `json:"stockMinimo,omitempty"`
	StockMaximo           float64 `json:"stockMaximo,omitempty"`
	Cantidad              float64 `json:"cantidad,omitempty"`
	IdImpuesto            int     `json:"idImpuesto,omitempty"`
	TipoPrecio            bool    `json:"tipoPrecio,omitempty"`
	PrecioCompra          float64 `json:"precioCompra,omitempty"`
	PrecioVentaGeneral    float64 `json:"precioVentaGeneral,omitempty"`
	Lote                  bool    `json:"lote,omitempty"`
	Inventario            bool    `json:"inventario,omitempty"`
	ValorInventario       int     `json:"valorInventario,omitempty"`
	Imagen                string  `json:"imagen,omitempty"`
	ClaveSat              string  `json:"claveSat,omitempty"`
	NuevaImagen           []byte  `json:"nuevaImagen,omitempty"`
	Descripcion           string  `json:"descripcion,omitempty"`

	Impuesto *Impuesto `json:"impuesto,omitempty"`

	DetalleUnidadCompra *Detalle `json:"detalleUnidadCompra"`
	DetalleMarca        *Detalle `json:"detalleMarca"`
	DetalleCategoria    *Detalle `json:"detalleCategoria"`
	DetalleEstado       *Detalle `json:"detalleEstado"`
}
