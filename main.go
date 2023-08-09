package main

import (
	"github.com/gin-gonic/gin"
	"github.com/khip01/gin-restfulapi-go/controllers/product_controller"
	"github.com/khip01/gin-restfulapi-go/models"
)

func main() {
	// Declare Route Gin Framework
	route := gin.Default()

	// Model Connect Database
	models.ConnectDatabase()

	// Set Route
	route.GET("/api/products", product_controller.Index)
	route.GET("/api/product/:id", product_controller.Show)
	route.POST("/api/product", product_controller.Create)
	route.PUT("/api/product/:id", product_controller.Update)
	route.DELETE("/api/product", product_controller.Delete)

	// Eksekusi di port 8081 karena port 8080 di laptop ini digunakan untuk IIS (localhost:8080)
	route.Run(":8081")
}
