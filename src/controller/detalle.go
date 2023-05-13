package controller

import (
	"net/http"
	"syssoftintegra-api/src/model"
	"syssoftintegra-api/src/service"

	"github.com/gin-gonic/gin"
)

// PingExample   godoc
// @Summary 	 Obtener detalle
// @Schemes
// @Description  Obtener el listado de detalle por id mantenimiento
// @Tags 		 Detalle
// @Accept 		 json
// @Produce 	 json
// @Param opcion query string true "Opciones de filtro 0-libre 1-para excluir el nombre al inciar la busqueda"
// @Param idMantenimiento query string true "Id del matenimiento"
// @Param nombre query string false "nombre a excluir" default:""
// @Success 	 200  {object}  model.Detalle
// @Failure 	 400  {object}  model.Error
// @Failure 	 500  {object}  model.Error
// @Router /lista-detalle-idmantenimiento [get]
func ListaDetalleIdMantenimiento(c *gin.Context) {
	opcion := c.Query("opcion")
	idMantenimiento := c.Query("idMantenimiento")
	nombre := c.Query("nombre")

	detalle, rpta := service.ListaDetalleIdMantenimiento(opcion, idMantenimiento, nombre)
	if rpta != "ok" {
		c.IndentedJSON(http.StatusInternalServerError, model.Error{Message: rpta})
		return
	}

	c.IndentedJSON(http.StatusOK, detalle)
}
