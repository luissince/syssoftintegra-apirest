package controller

import (
	"fmt"
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
// @Param opcion body  model.Login true "Estructura para realizar la consulta"
// @Success 	 200  {object}  model.Empleado
// @Failure 	 400  {object}  model.Error
// @Failure 	 500  {object}  model.Error
// @Router /login [post]
func Login(c *gin.Context) {
	login := model.Login{}

	err := c.BindJSON(&login)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, model.Error{Message: "No se pudo parsear el body"})
		return
	}

	empleado, rpta := service.Login(login.Usuario, login.Clave)
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

// PingExample   godoc
// @Summary 	 Lista de empleados o usarios del sistema
// @Schemes
// @Description  Listado de empleados o usuario con los datos principales
// @Tags 		 Empleado
// @Accept 		 json
// @Produce 	 json
// @Param opcion query int true "Opciones de filtro 0-libre 1-para iniciar la busqueda"
// @Param search query string false "Datos para el filtro" default:""
// @Param posicionPagina query int true "Inicio de la paginación"
// @Param filasPorPagina query int true "Filas por paginación"
// @Success 	 200  {object}  []model.Empleado
// @Failure 	 400  {object}  model.Error
// @Failure 	 500  {object}  model.Error
// @Router /empleados [get]
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
	if rpta != "ok" {
		c.IndentedJSON(http.StatusInternalServerError, model.Error{Message: "No se encontraron resultados"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"total": total, "resultado": empleados})
}

// PingExample   godoc
// @Summary 	 Obtener empleado po su Id
// @Schemes
// @Description  Ruta usada para traer datos relevante al momento de realizar una edición
// @Tags 		 Empleado
// @Accept 		 json
// @Produce 	 json
// @Param idEmpleado query string true "Id del empleado"
// @Success 	 200  {object}  model.Empleado
// @Failure 	 400  {object}  model.Error
// @Failure 	 500  {object}  model.Error
// @Router /empleado [get]
func GetEmpleadoById(c *gin.Context) {

	idEmpleado := c.Query("idEmpleado")

	empleado, rpta := service.GetEmpleadoById(idEmpleado)
	if rpta != "ok" {
		c.IndentedJSON(http.StatusInternalServerError, model.Error{Message: "No se encontraron resultados"})
		return
	}

	c.IndentedJSON(http.StatusOK, empleado)
}

// PingExample   godoc
// @Summary 	 Registrar Empleado
// @Schemes
// @Description  Proceso para registrar empleado con la estructura predefinida
// @Tags 		 Empleado
// @Accept 		 json
// @Produce 	 json
// @Param opcion body  model.Empleado true "Estructura para realizar la consulta"
// @Success 	 200  {string}  string
// @Failure 	 400  {object}  model.Error
// @Failure 	 500  {object}  model.Error
// @Router /empleado [post]
func InsertEmpledo(c *gin.Context) {
	var empleado model.Empleado

	rol := &model.Rol{}
	empleado.Rol = rol

	detalle := &model.Detalle{}
	empleado.Detalle = detalle

	if err := c.BindJSON(&empleado); err != nil {
		fmt.Println(err)
		c.IndentedJSON(http.StatusBadRequest, model.Error{Message: "No se pudo parsear el body"})
		return
	}

	rpta := service.InsertEmpledo(&empleado)
	if rpta != "insert" {
		c.IndentedJSON(http.StatusInternalServerError, model.Error{Message: rpta})
		return
	}

	c.IndentedJSON(http.StatusCreated, "Se registro correctamente el empleado.")
}

// PingExample   godoc
// @Summary 	 Actualizar Empleado
// @Schemes
// @Description  Proceso para actualizar empleado con la estructura predefinida
// @Tags 		 Empleado
// @Accept 		 json
// @Produce 	 json
// @Success 	 200  {string}  string
// @Failure 	 400  {object}  model.Error
// @Failure 	 500  {object}  model.Error
// @Router /empleado [put]
func UpdateEmpledo(c *gin.Context) {

	var empleado model.Empleado

	err := c.BindJSON(&empleado)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, model.Error{Message: "No se pudo parsear el body"})
		return
	}

	rpta := service.UpdateEmpledo(&empleado)
	if rpta == "update" {
		c.IndentedJSON(http.StatusOK, gin.H{"message": rpta})
		return
	}

	if rpta == "empty" {
		c.IndentedJSON(http.StatusInternalServerError, model.Error{Message: "No se pudo realizar la operación"})
		return
	}

	c.IndentedJSON(http.StatusInternalServerError, model.Error{Message: rpta})
}

// PingExample   godoc
// @Summary 	 Actualizar Empleado
// @Schemes
// @Description  Proceso para actualizar empleado con la estructura predefinida
// @Tags 		 Empleado
// @Accept 		 json
// @Produce 	 json
// @Success 	 200  {string}  string
// @Failure 	 400  {object}  model.Error
// @Failure 	 500  {object}  model.Error
// @Router /empleado [delete]
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
