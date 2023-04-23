package controller

import (
	"net/http"
	"syssoftintegra-api/src/model"
	"syssoftintegra-api/src/service"

	"github.com/gin-gonic/gin"
)

// PingExample   godoc
// @Summary 	 Para obtener la empresa actual
// @Schemes
// @Description  Obtener la información de la empresa
// @Tags 		 Empresa
// @Accept 		 json
// @Produce 	 json
// @Success 	 200  {object}  model.Empresa
// @Failure 	 400  {object}  model.Error
// @Failure 	 500  {object}  model.Error
// @Router /empresa [get]
func ObtenerEmpresa(c *gin.Context) {
	empresa, rpta := service.ObtenerEmpresa()
	if rpta == "empty" {
		c.IndentedJSON(http.StatusBadRequest, model.Error{Message: "No existe la empresa."})
		return
	}

	if rpta != "ok" {
		c.IndentedJSON(http.StatusInternalServerError, model.Error{Message: rpta})
		return
	}

	c.IndentedJSON(http.StatusOK, empresa)
}
