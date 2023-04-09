package main

import (
	// "fmt"
	"syssoftintegra-api/src/controller"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	docs "syssoftintegra-api/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {

	router := gin.Default()

	// Middleware para CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	router.Use(cors.New(config))

	// Agregar el swagger
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath

	empleado := router.Group(basePath)
	{
		empleado.GET("/login", controller.Login)
		empleado.GET("/empleados", controller.GetAllEmpleado)
		empleado.GET("/empleado", controller.GetEmpleadoById)
		empleado.POST("/empleado", controller.IUEmpledo)
		// empleado.PUT("/empleado", controller.UpdateEmpleado)
		empleado.DELETE("/empleado/:idEmpleado", controller.DeleteEmpleado)

	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.Run("localhost:3000")

	// Rutas del servicio de la API
	// router.GET(v1+"/login", service.Login)

	// GET /api/v1/users/:id
	// Obtiene un usuario por su ID
	//
	// Respuesta exitosa:
	// Código: 200
	// Cuerpo:
	// {
	//     "id": 1,
	//     "name": "Juan",
	//     "email": "juan@example.com"
	// }
	//
	// Respuesta de error:
	// Código: 404
	// Cuerpo:
	// {
	//     "message": "No se encontró el usuario"
	// }

}
