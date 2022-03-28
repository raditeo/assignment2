package main

import (
	"assignment2/controllers"
	"assignment2/database"

	"github.com/gin-gonic/gin"
)

func main() {
	database.StartDB()
	router := gin.Default()

	router.POST("/orders", controllers.AddOrder)
	router.GET("/orders", controllers.GetOrders)
	router.DELETE("/orders/:orderId", controllers.DeleteOrder)
	router.PUT("/orders/:orderId", controllers.EditOrder)
	router.Run(":8080")
}
