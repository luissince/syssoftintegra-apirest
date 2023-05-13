package routes

import (
	"syssoftintegra-api/src/controller"

	"github.com/gin-gonic/gin"
)

func SuministroRouters(router *gin.RouterGroup) {

}

func EmpresaRoutes(router *gin.RouterGroup) {
	router.GET("/empresa", controller.ObtenerEmpresa)
}

func EmpleadoRoutes(router *gin.RouterGroup) {
	/*
		usersRouter := router.Group("/users")

			{
				usersRouter.GET("/", getUsersHandler)
				usersRouter.POST("/", createUserHandler)
				usersRouter.PUT("/:id", updateUserHandler)
				usersRouter.DELETE("/:id", deleteUserHandler)
			}
	*/

	router.GET("/login", controller.Login)
	router.GET("/empleados", controller.GetAllEmpleado)
	router.GET("/empleado", controller.GetEmpleadoById)
	router.POST("/empleado", controller.InsertUpdateEmpledo)
	router.DELETE("/empleado/:idEmpleado", controller.DeleteEmpleado)

}

func MonedaRoutes(router *gin.RouterGroup) {
	router.GET("/moneda-combobox", controller.GetMonedaComboBox)
	router.GET("/monedas", controller.GetAllMoneda)
	router.GET("/moneda", controller.GetMonedaById)
	router.POST("/moneda", controller.InsertUpdateMoneda)
	router.DELETE("/moneda/:idMoneda", controller.DeleteMoneda)
}

func BancoRoutes(router *gin.RouterGroup) {

}

func RolRoutes(router *gin.RouterGroup) {
	router.GET("/listar-roles", controller.ListarRoles)
}

func DetalleRoutes(router *gin.RouterGroup) {
	router.GET("/lista-detalle-idmantenimiento", controller.ListaDetalleIdMantenimiento)
}
