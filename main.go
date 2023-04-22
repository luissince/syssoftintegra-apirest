package main

import (
	// "fmt"
	"syssoftintegra-api/src/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	docs "syssoftintegra-api/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {

	app := gin.Default()

	// Middleware para CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	app.Use(cors.New(config))

	// Agregar el swagger
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath

	routes.EmpleadoRoutes(app.Group(basePath))
	routes.MonedaRoutes(app.Group(basePath))

	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	app.Run("localhost:3000")

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
