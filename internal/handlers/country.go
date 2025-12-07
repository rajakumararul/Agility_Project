package handlers

import (
	"net/http"

	"github.com/rajakumararul/Agility_Project/internal/services"

	"github.com/gin-gonic/gin"
)

func RegisterCountryRoutes(r *gin.Engine) {
	group := r.Group("/api/country")
	{
		group.GET("/name/:name", GetByName)
		group.GET("/currency/:currency", GetByCurrency)
		group.GET("/capital/:capital", GetByCapital)
	}
}

func GetByName(c *gin.Context) {
	name := c.Param("name")
	data, err := services.FetchCountryByName(name)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}

func GetByCurrency(c *gin.Context) {
	currency := c.Param("currency")
	data, err := services.FetchCountryByCurrency(currency)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}

func GetByCapital(c *gin.Context) {
	capital := c.Param("capital")
	data, err := services.FetchCountryByCapital(capital)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}
