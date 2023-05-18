package main

import (
	// "fmt"
	"fmt"
	"os"
	"syssoftintegra-api/src/routes"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"

	docs "syssoftintegra-api/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// func corsMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
// 		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
// 		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
// 		c.Writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
// 		c.Next()
// 	}
// }

// Funcion princiapl de la aplicación.

// @title Api de SysSoft Integra
// @version 1.0
// @description Api para consultar las rutas de la aplicación.
func main() {
	// Cargar las variables de entorno
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error en cargar las viriables de entorno: ", err.Error())
		return
	}

	var url_domain string = os.Getenv("URL_DOMAIN")
	var tz_location string = os.Getenv("TZ_LOCATION")

	// Estabecle la zora horario
	time.LoadLocation(tz_location)

	// Inicializa GIN para correr el servidor
	app := gin.Default()

	// Middleware para CORS
	//app.Use(corsMiddleware())
	app.Use(cors.Default())

	// Agregar el swagger
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath

	routes.EmpresaRoutes(app.Group(basePath))
	routes.EmpleadoRoutes(app.Group(basePath))
	routes.MonedaRoutes(app.Group(basePath))
	routes.RolRoutes(app.Group(basePath))
	routes.DetalleRoutes(app.Group(basePath))

	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	app.Run(url_domain)
}
