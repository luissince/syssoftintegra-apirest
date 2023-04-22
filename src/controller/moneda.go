package controller

import (
	"context"
	"net/http"
	"syssoftintegra-api/src/model"
	"syssoftintegra-api/src/service"

	"strconv"

	"github.com/gin-gonic/gin"
)

var ctx = context.Background()

func GetMonedaComboBox(c *gin.Context) {

	monedas, rpta := service.GetMonedaComboBox(ctx)
	if rpta == "empty" {
		c.IndentedJSON(http.StatusBadRequest, model.Error{Message: "No existe ninguna moneda"})
		return
	}
	if rpta == "ok" {
		c.IndentedJSON(http.StatusOK, monedas)
		return
	}

	c.IndentedJSON(http.StatusInternalServerError, model.Error{Message: rpta})

}

func GetAllMoneda(c *gin.Context) {
	opcion, err := strconv.Atoi(c.Query("opcion"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, model.Error{Message: "No se puede parcear el primer parametro"})
		return
	}
	search := c.Query("search")
	posicionPagina, err := strconv.Atoi(c.Query("posicionPagina"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, model.Error{Message: "No se puede parcear el segundo parametro"})
		return
	}
	filasPorPagina, err := strconv.Atoi(c.Query("filasPorPagina"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, model.Error{Message: "No se puede parcear el tercer parametro"})
		return
	}

	monedas, total, rpta := service.GetAllMoneda(ctx, opcion, search, posicionPagina, filasPorPagina)
	if rpta == "empty" {
		c.IndentedJSON(http.StatusBadRequest, model.Error{Message: "No se encontraron registros"})
		return
	}
	if rpta == "ok" {
		c.IndentedJSON(http.StatusOK, gin.H{"total": total, "resultado": monedas})
		return
	}

	c.IndentedJSON(http.StatusInternalServerError, model.Error{Message: rpta})
}

func GetMonedaById(c *gin.Context) {

	idMoneda, err := strconv.Atoi(c.Query("idMoneda"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, model.Error{Message: "El tipo de dato no es numérico"})
		return
	}

	moneda, rpta := service.GetMonedaById(ctx, idMoneda)
	if rpta == "empty" {
		c.IndentedJSON(http.StatusBadRequest, model.Error{Message: "La moneda no existe"})
		return
	}
	if rpta == "ok" {
		c.IndentedJSON(http.StatusOK, moneda)
		return
	}

	c.IndentedJSON(http.StatusInternalServerError, model.Error{Message: rpta})
}

func InsertUpdateMoneda(c *gin.Context) {

	var moneda model.Moneda

	err := c.BindJSON(&moneda)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, model.Error{Message: "No se pudo parsear el body"})
		return
	}

	rpta := service.ValidarNombreMoneda(moneda.IdMoneda, moneda.Nombre)
	if rpta == "exists" {
		c.IndentedJSON(http.StatusBadRequest, model.Error{Message: "Existe una moneda con el mismo nombre"})
		return
	}
	if rpta == "empty" {

		operacion := service.InsertUpdateMoneda(ctx, &moneda)
		if operacion == "empty" {
			c.IndentedJSON(http.StatusInternalServerError, model.Error{Message: "No se pudo realizar la operación"})
			return
		} else if operacion == "insert" {
			c.IndentedJSON(http.StatusOK, gin.H{"Message": "insert"})
			return
		} else if operacion == "update" {
			c.IndentedJSON(http.StatusOK, gin.H{"Message": "update"})
			return
		} else {
			c.IndentedJSON(http.StatusInternalServerError, model.Error{Message: operacion})
			return
		}

	}

	c.IndentedJSON(http.StatusInternalServerError, model.Error{Message: rpta})

}

func DeleteMoneda(c *gin.Context) {

	idMoneda, err := strconv.Atoi(c.Param("idMoneda"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, model.Error{Message: "El tipo de dato no es numérico"})
		return
	}

	rpta := service.ValidarSistemaMoneda(idMoneda)
	if rpta == "sistema" {
		c.IndentedJSON(http.StatusBadRequest, model.Error{Message: "La moneda es parte del sistema"})
		return
	}
	if rpta == "empty" {
		operacion := service.DeleteMoneda(ctx, idMoneda)
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

	c.IndentedJSON(http.StatusInternalServerError, model.Error{Message: rpta})
}
