package controller

import (
	"net/http"
	"syssoftintegra-api/src/model"
	"syssoftintegra-api/src/service"

	"github.com/gin-gonic/gin"
)

// PingExample   godoc
// @Summary 	 Lista de roles
// @Schemes
// @Description  Obtener el listado de roles para los modulos que necesiten
// @Tags 		 Rol
// @Accept 		 json
// @Produce 	 json
// @Success 	 200  {object}  model.Rol
// @Failure 	 400  {object}  model.Error
// @Failure 	 500  {object}  model.Error
// @Router /listar-roles [get]
func ListarRoles(c *gin.Context) {
	rol, rpta := service.ListarRoles()
	if rpta != "ok" {
		c.IndentedJSON(http.StatusInternalServerError, model.Error{Message: rpta})
		return
	}

	c.IndentedJSON(http.StatusOK, rol)
}
