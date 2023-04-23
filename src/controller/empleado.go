package controller

import (
	"net/http"
	"strconv"
	"syssoftintegra-api/src/model"
	"syssoftintegra-api/src/service"

	"github.com/gin-gonic/gin"
)

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

	empleado, rpta := service.Login(usuario, clave)
	if rpta == "empty" {
		c.IndentedJSON(http.StatusBadRequest, model.Error{Message: "El usuario o la clave es incorrecto, o no existe el usuario"})
		return
	}
	if rpta == "ok" {
		c.IndentedJSON(http.StatusOK, empleado)
		return
	}

	c.IndentedJSON(http.StatusInternalServerError, model.Error{Message: rpta})

}

func GetAllEmpleado(c *gin.Context) {

	//Se lee los parametros del la url
	opcion, err := strconv.Atoi(c.Query("opcion"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, model.Error{Message: "No se puede parcear el primer parametro"})
		return
	}

	search := c.Query("search")

	posicionPagina, err := strconv.Atoi(c.Query("posicionPagina"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, model.Error{Message: "No se puede parcear el tercer parametro"})
		return
	}
	filasPorPagina, err := strconv.Atoi(c.Query("filasPorPagina"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, model.Error{Message: "No se puede parcear el cuarto parametro"})
		return
	}

	// Se hace la petición a la base de datos
	empleados, total, rpta := service.GetAllEmpleado(opcion, search, posicionPagina, filasPorPagina)
	if rpta == "empty" {
		c.IndentedJSON(http.StatusInternalServerError, model.Error{Message: "No se encontraron resultados"})
		return
	}
	if rpta == "ok" {
		c.IndentedJSON(http.StatusOK, gin.H{"total": total, "resultado": empleados})
	}

	c.IndentedJSON(http.StatusInternalServerError, model.Error{Message: rpta})

}

func GetEmpleadoById(c *gin.Context) {

	idEmpleado := c.Query("idEmpleado")

	empleado, rpta := service.GetEmpleadoById(idEmpleado)
	if rpta == "empty" {
		c.IndentedJSON(http.StatusInternalServerError, model.Error{Message: "No se encontraron resultados"})
		return
	}
	if rpta == "ok" {
		c.IndentedJSON(http.StatusOK, empleado)
		return
	}

	c.IndentedJSON(http.StatusInternalServerError, model.Error{Message: rpta})

}

func InsertUpdateEmpledo(c *gin.Context) {

	var empleado model.Empleado

	err := c.BindJSON(&empleado)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, model.Error{Message: "No se pudo parsear el body"})
		return
	}

	rpta := service.InsertUpdateEmpledo(&empleado)
	if rpta == "empty" {
		c.IndentedJSON(http.StatusInternalServerError, model.Error{Message: "No se pudo realizar la operación"})
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

	c.IndentedJSON(http.StatusInternalServerError, model.Error{Message: rpta})

}

func DeleteEmpleado(c *gin.Context) {
	idEmpleado := c.Param("idEmpleado")

	validar := service.ValidarSistemaEmpledo(idEmpleado)
	if validar == "sistema" {
		c.IndentedJSON(http.StatusBadRequest, model.Error{Message: "No se puede eliminar el empledo por que es parte del sistema"})
		return
	}
	if validar == "empty" {
		operacion := service.DeleteEmpleado(idEmpleado)
		if operacion == "empty" {
			c.IndentedJSON(http.StatusInternalServerError, model.Error{Message: "No se pudo realizar la operación"})
			return
		} else if operacion == "delete" {
			c.IndentedJSON(http.StatusOK, gin.H{"Message": "delete"})
			return
		} else {
			c.IndentedJSON(http.StatusInternalServerError, model.Error{Message: operacion})
			return
		}

	}

	c.IndentedJSON(http.StatusInternalServerError, model.Error{Message: validar})

}
