package controller

import (
	"context"
	"net/http"
	"strconv"
	"syssoftintegra-api/src/model"
	"syssoftintegra-api/src/service"

	"github.com/gin-gonic/gin"
)

var contx = context.Background()

// PingExample   godoc
// @Summary 	 Para el inicio de sesión
// @Schemes
// @Description  Iniciar Sesión del Empleado
// @Tags 		 Empleado
// @Accept 		 json
// @Produce 	 json
// @Param usuario query string true "Usuario para iniciar sesión"
// @Param clave query string true "Clave para iniciar sesión"
// @Success 	 200  {object}  model.Empleado
// @Failure 	 400  {object}  model.Error
// @Failure 	 500  {object}  model.Error
// @Router /login [get]
func Login(c *gin.Context) {

	usuario := c.Query("usuario")
	clave := c.Query("clave")

	empleado, err := service.Login(contx, usuario, clave)
	if err == "error" {
		c.IndentedJSON(http.StatusInternalServerError, model.Error{
			Message: "Se produjo un error interno.",
		})
		return
	}

	if err == "empty" {
		c.IndentedJSON(http.StatusBadRequest, model.Error{
			Message: "Usuario o clave incorrecta.",
		})
		return
	}

	c.IndentedJSON(http.StatusOK, empleado)

}

func GetAllEmpleado(c *gin.Context) {

	//Se lee los parametros del la url

	opcion, err := strconv.Atoi(c.Query("opcion"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, model.Error{
			Message: "No se puede parcear el primer parametro.",
		})
		return
	}

	search := c.Query("search")

	posicionPagina, err := strconv.Atoi(c.Query("posicionPagina"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, model.Error{
			Message: "No se puede parcear el tercer parametro.",
		})
		return
	}

	filasPorPagina, err := strconv.Atoi(c.Query("filasPorPagina"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, model.Error{
			Message: "No se puede parcear el cuarto parametro.",
		})
		return
	}

	// Se hace la petición a la base de datos
	empleados, total, errString := service.GetAllEmpleado(contx, opcion, search, posicionPagina, filasPorPagina)
	if errString == "error" {
		c.IndentedJSON(http.StatusInternalServerError, model.Error{
			Message: "Se produjo un error interno.",
		})
		return
	}

	if errString == "empty" {
		c.IndentedJSON(http.StatusBadRequest, model.Error{
			Message: "No se encontrantron resultado.",
		})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"total": total, "resultado": empleados})

}

func GetEmpleadoById(c *gin.Context) {

	// idEmpleado, err := strconv.Atoi(c.Query("idEmpleado"))
	idEmpleado := c.Query("idEmpleado")
	// idEmpleado := c.Param("idEmpleado")

	empleado, err := service.GetEmpleadoById(contx, idEmpleado)

	if err == "error" {
		c.IndentedJSON(http.StatusInternalServerError, model.Error{
			Message: "Se produjo un error interno.",
		})
		return
	}

	if err == "empty" {
		c.IndentedJSON(http.StatusBadRequest, model.Error{
			Message: "No se encontrantron resultado.",
		})
		return
	}

	c.IndentedJSON(http.StatusOK, empleado)

}

func IUEmpledo(c *gin.Context) {

	var empleado model.Empleado

	err := c.BindJSON(&empleado)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, model.Error{
			Message: "No se pudo parsear el body",
		})
		return
	}

	rpta := service.IUEmpledo(contx, &empleado)

	if rpta == "error" {
		c.IndentedJSON(http.StatusInternalServerError, model.Error{
			Message: "Se produjo un error interno.",
		})
		return
	}
	if rpta == "empty" {
		c.IndentedJSON(http.StatusBadRequest, model.Error{
			Message: "Se produjo un error al realizar la operación.",
		})
		return
	}
	if rpta == "update" {
		c.IndentedJSON(http.StatusOK, gin.H{"message": rpta})
		return
	}
	if rpta == "insert" {
		c.IndentedJSON(http.StatusOK, gin.H{"message": rpta})
		return
	}

}

func DeleteEmpleado(c *gin.Context) {
	idEmpleado := c.Param("idEmpleado")

	rpta := service.DeleteEmpleado(contx, idEmpleado)

	if rpta == "error" {
		c.IndentedJSON(http.StatusInternalServerError, model.Error{
			Message: "Se produjo un error interno.",
		})
		return
	}

	if rpta == "empty" {
		c.IndentedJSON(http.StatusBadRequest, model.Error{
			Message: "Se produjo un error al realizar la operación.",
		})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": rpta})
}
